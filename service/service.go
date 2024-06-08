package service

import (
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/external"
)

var (
	MyService          *Services
	DefaultTimeout     = 10 * time.Second
	rootURL            = "localhost:80"
	BasePathMessageBus = "v2/message_bus"
)

type Services struct {
	gateway        external.ManagementService
	record         *RecordService
	runtimePath    string
	db             *DBService
	messageService *MessageBusService
	logService     *LogService
}

func Initialize(runtimePath string) {
	MyService = &Services{
		runtimePath: runtimePath,
	}
}

func (s *Services) DBService() *DBService {
	if s.db == nil {
		s.db = NewDBService()
	}
	return s.db
}

func (s *Services) RecordService() *RecordService {
	if s.record == nil {
		s.record = NewRecordService()
	}
	return s.record
}

func (s *Services) MessageBus() *MessageBusService {
	if s.messageService == nil {
		s.messageService = NewMessageBusService()
	}
	return s.messageService
}

func (s *Services) LogService() *LogService {
	if s.logService == nil {
		s.logService = NewLogService()
	}
	return s.logService
}

func (s *Services) Gateway() external.ManagementService {
	if s.gateway == nil {
		gateway, err := external.NewManagementService(s.runtimePath)
		if err != nil && len(s.runtimePath) > 0 {
			panic(err)
		}

		s.gateway = gateway
	}

	return s.gateway
}
