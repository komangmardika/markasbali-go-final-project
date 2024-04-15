package services_test

import (
	"github.com/stretchr/testify/assert"
	"markasbali_go_final_project/cli-service/services"
	"testing"
)

func TestSendErrorToWebSocketServer(t *testing.T) {
	err := services.SendErrorToWebSocketServer("error")
	assert.Nil(t, err)
}
