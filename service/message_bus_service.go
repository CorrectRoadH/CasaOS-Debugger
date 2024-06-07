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
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/googollee/go-socket.io/parser"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type MessageBusService struct {
	record        bool
	eventTypeList []message_bus.EventType
}

func NewMessageBusService() *MessageBusService {
	s := &MessageBusService{
		record: false,
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

		response, err := client.GetEventTypesWithResponse(ctx)
		if err != nil {
			log.Fatalln(err.Error())
		}
		s.eventTypeList = append(s.eventTypeList, *response.JSON200...)
	}
	return s.eventTypeList
}

func (s *MessageBusService) StartRecord(_ context.Context) {
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
				timestamp, cerr := time.Parse(time.RFC3339, fmt.Sprintf("%f", rawTimestamp))
				if cerr == nil {
					event.Timestamp = lo.ToPtr(timestamp)
				}
			}

			err = MyService.DBService().InsertEvent(*event.Uuid, event.Properties, event.SourceID, event.Name, event.Timestamp)
			if err != nil {
				logger.Error("InsertEvent error: ", zap.Error(err))
			}
		}
	}
}

func (s *MessageBusService) MessageHistory(sourceID *string, name *string, offset int, length int) ([]message_bus.Event, error) {
	var result []message_bus.Event
	if offset < 0 {
		offset = 0
	}
	if length < 0 {
		length = 0
	}

	result, err := MyService.DBService().QueryEvent(name, sourceID, offset, length)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *MessageBusService) Sources() ([]codegen.SourceID, error) {
	sourceList, err := MyService.DBService().SourceList()
	if err != nil {
		return nil, err
	}
	return lo.Map(sourceList, func(sourceID string, _ int) codegen.SourceID {
		return codegen.SourceID(sourceID)
	}), nil
}
