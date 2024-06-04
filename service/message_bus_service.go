package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/CorrectRoadH/CasaOS-Debugger/codegen"
	"github.com/CorrectRoadH/CasaOS-Debugger/codegen/message_bus"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/googollee/go-socket.io/parser"
	"github.com/samber/lo"

	mapset "github.com/deckarep/golang-set/v2"
)

type MessageBusService struct {
	record         bool
	messageHistory []message_bus.Event
	eventTypeList  []message_bus.EventType
}

func NewMessageBusService() *MessageBusService {
	s := &MessageBusService{
		record:         false,
		messageHistory: []message_bus.Event{},
	}
	return s
}

func (s *MessageBusService) EventTypeList(ctx context.Context) []message_bus.EventType {
	if len(s.eventTypeList) == 0 {
		url := fmt.Sprintf("http://%s/%s", rootURL, BasePathMessageBus)

		client, err := message_bus.NewClientWithResponses(url)
		if err != nil {
			log.Fatalln(err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
		defer cancel()

		response, err := client.GetEventTypesWithResponse(ctx)
		if err != nil {
			log.Fatalln(err.Error())
		}
		for _, eventType := range *response.JSON200 {
			s.eventTypeList = append(s.eventTypeList, eventType)
		}
	}
	return s.eventTypeList
}

func (s *MessageBusService) StartRecord() {
	dialer := engineio.Dialer{
		Transports: []transport.Transport{
			websocket.Default,
			polling.Default,
		},
	}

	sioURL := fmt.Sprintf("http://%s/%s/socket.io", strings.TrimRight(rootURL, "/"), BasePathMessageBus)
	conn, err := dialer.Dial(sioURL, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	log.Printf("subscribed to %s via socketio", sioURL)

	decoder := parser.NewDecoder(conn)

	s.record = true
	for {
		header := parser.Header{}
		name := ""
		if err := decoder.DecodeHeader(&header, &name); err != nil {
			if err == io.EOF {
				time.Sleep(time.Millisecond * 100)
				continue
			}

			log.Fatalln(err.Error())
		}

		values, err := decoder.DecodeArgs([]reflect.Type{
			reflect.TypeOf(map[string]interface{}{}),
		})
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Millisecond * 100)
				continue
			}

			log.Println(err.Error())
		}
		decoder.Close()

		for _, value := range values {
			var event message_bus.Event

			rawEvent := value.Interface().(map[string]interface{})

			event.Name = name

			if _, ok := rawEvent["SourceID"]; ok {
				event.SourceID = rawEvent["SourceID"].(string)
			}

			if _, ok := rawEvent["Properties"]; ok {
				rawPropertise := rawEvent["Properties"].(map[string]interface{})
				// string json to map
				var properties map[string]string = make(map[string]string)
				for key, value := range rawPropertise {
					properties[key] = value.(string)
				}
				event.Properties = properties
			}

			if _, ok := rawEvent["uuid"]; ok {
				event.Uuid = lo.ToPtr(rawEvent["uuid"].(string))
			}

			if _, ok := rawEvent["Timestamp"]; ok {
				rawTimestamp := rawEvent["Timestamp"].(float64)
				timestamp, err := time.Parse(time.RFC3339, fmt.Sprintf("%f", rawTimestamp))
				if err == nil {
					event.Timestamp = lo.ToPtr(timestamp)
				}
			}

			// fmt.Println(string(output))
			s.messageHistory = append(s.messageHistory, event)
		}
	}
}

func (s *MessageBusService) MessageHistory(sourceID *string, name *string, offset int, length int) ([]message_bus.Event, error) {
	var result []message_bus.Event
	if sourceID != nil {
		result = lo.Filter(s.messageHistory, func(event message_bus.Event, _ int) bool {
			return event.SourceID == *sourceID
		})
	}
	if name != nil {
		result = lo.Filter(result, func(event message_bus.Event, _ int) bool {
			return event.Name == *name
		})
	}
	if offset < 0 {
		offset = 0
	}
	if length < 0 {
		length = 0
	}
	if offset+length > len(result) {
		length = len(result) - offset
	}
	return result[offset : offset+length], nil
}

func (s *MessageBusService) Sources() ([]codegen.SourceID, error) {
	sourceSet := mapset.NewSet[string]()
	for _, event := range s.messageHistory {
		sourceSet.Add(event.SourceID)
	}
	return sourceSet.ToSlice(), nil
}
