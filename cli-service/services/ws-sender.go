package services

import (
	"encoding/json"
	"github.com/gorilla/websocket"
)

var (
	upgrader websocket.Upgrader
)

type ErrorMessage struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	TableName string `json:"tableName"`
}

func SendErrorToWebSocketServer(errorMessage string) error {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		return err
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Create an ErrorMessage struct
	errMsg := ErrorMessage{
		Type:      "error",
		Message:   errorMessage,
		TableName: "",
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

func SendSeedingProcessToWebSocketServer(dbName string, tableName string) error {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		return err
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// Create an ErrorMessage struct
	errMsg := ErrorMessage{
		Type:      "seeding",
		Message:   dbName,
		TableName: tableName,
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
