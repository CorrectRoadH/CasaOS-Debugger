package service

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type LogService struct{}

func NewLogService() *LogService {
	return &LogService{}
}

var serviceMap = map[string]string{
	"casaos-installer":      "/var/log/casaos/installer.log",
	"casaos-app-management": "/var/log/casaos/app-management.log",
	"zimaos-mod-management": "/var/log/casaos/ZimaOS-ModManagement.log",
	"zimaos":                "/var/log/casaos/log.log",
	"zimaos-local-storage":  "/var/log/casaos/local-storage.log",
	"casaos-gateway":        "/var/log/casaos/gateway.log",
	"casaos-user-service":   "/var/log/casaos/user-service.log",
}

var (
	ErrServiceNameNotFound = errors.New("service name not found")
	ErrLogNotFound         = errors.New("log not found")
)

func (s *LogService) QueryLog(_ context.Context, serviceName string, offset int, length int) ([]string, error) {
	logPath, ok := serviceMap[serviceName]
	if !ok {
		return []string{}, ErrServiceNameNotFound
	}

	// read file
	// if file not found, return ErrLogNotFound
	// if error occurred, return error
	// if success, return log content
	openedFile, err := os.Open(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, ErrLogNotFound
		}
		return []string{}, err
	}
	defer openedFile.Close()

	// Read the file content
	content, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return []string{}, err
	}

	// spill the content into lines
	lines := strings.Split(string(content), "\n")

	// check offset and length
	if offset < 0 {
		offset = 0
	}
	if length < 0 {
		length = 0
	}
	if offset >= len(lines) {
		return []string{}, nil
	}
	if offset+length >= len(lines) {
		return lines[offset:], nil
	}
	return lines[offset : offset+length], nil
}
