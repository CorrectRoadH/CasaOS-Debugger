package service_test

import (
	"context"
	"testing"

	"github.com/CorrectRoadH/CasaOS-Debugger/service"
	"github.com/stretchr/testify/assert"
)

func TestGetServiceLog(t *testing.T) {
	loggerSerivce := service.NewLogService()
	logs, err := loggerSerivce.QueryLog(context.Background(), "casaos-installer", 0, 10)

	assert.Nil(t, err)
	assert.Equal(t, len(logs), 5)
}
