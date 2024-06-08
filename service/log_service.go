package service

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

var serviceMap = map[string]string{
	"casaos-installer":          "/var/log/casaos/installer.log",
	"casaos-app-management.log": "/var/log/casaos/app-management.log",
	"zimaos-mod-management":     "/var/log/casaos/ZimaOS-ModManagement.log",
}

var (
	ErrServiceNameNotFound = errors.New("service name not found")
	ErrLogNotFound         = errors.New("log not found")
)

func (s *LogService) QueryLog(_ context.Context, serviceName string) (string, error) {
	logPath, ok := serviceMap[serviceName]
	if !ok {
		return "", ErrServiceNameNotFound
	}

	// read file
	// if file not found, return ErrLogNotFound
	// if error occurred, return error
	// if success, return log content
	openedFile, err := os.Open(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", ErrLogNotFound
		}
		return "", err
	}
	defer openedFile.Close()

	// Read the file content
	content, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
