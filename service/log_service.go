package service

import (
	"context"
	"errors"

	"github.com/coreos/go-systemd/sdjournal"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

var serviceMap = map[string]string{
	"casaos-installer":      "casaos-installer.service",
	"casaos-app-management": "/var/log/casaos/app-management.log",
	"zimaos-mod-management": "/var/log/casaos/ZimaOS-ModManagement.log",
	"zimaos":                "zimaos.service",
	"zimaos-local-storage":  "/var/log/casaos/local-storage.log",
	"casaos-gateway":        "casaos-gateway.service",
	"casaos-user-service":   "/var/log/casaos/user-service.log",
}

var (
	ErrServiceNameNotFound = errors.New("service name not found")
	ErrLogNotFound         = errors.New("log not found")
)

func (s *LogService) QueryLog(_ context.Context, serviceName string, offset int, length int) ([]string, error) {
	systemdServiceName, ok := serviceMap[serviceName]
	logs := []string{}
	if !ok {
		return logs, ErrServiceNameNotFound
	}

	j, err := sdjournal.NewJournal()
	if err != nil {
		return logs, err

	}
	defer j.Close()

	// 添加匹配条件，只查看 smb.service 的日志
	err = j.AddMatch("_SYSTEMD_UNIT=" + systemdServiceName)
	if err != nil {
		return logs, err

	}

	// 从尾部开始读取最近的 10 条日志
	if err := j.SeekTail(); err != nil {
		return logs, err

	}
	if _, err := j.PreviousSkip(uint64(length)); err != nil {
		return logs, err

	}

	// 读取并打印日志
	for {
		n, err := j.Next()
		if err != nil {
			return logs, err
		}
		if n == 0 {
			break
		}

		entry, err := j.GetEntry()
		if err != nil {
			return logs, err
		}

		logs = append(logs, entry.Fields["MESSAGE"])
	}
	return logs, nil
}
