package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

var (
	upgrader websocket.Upgrader
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func SendErrorToWebSocketServer(errorMessage string) error {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create an ErrorMessage struct
	errMsg := ErrorMessage{
		Message: errorMessage,
	}

	// Marshal the ErrorMessage to JSON
	jsonBytes, err := json.Marshal(errMsg)
	if err != nil {
		return err
	}

	// Write the JSON message to the WebSocket server
	if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
		return err
	}

	return nil
}
