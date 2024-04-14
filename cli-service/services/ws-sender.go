package services

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func sendMessageToWebSocketService(whatToSend string) error {
	// Prepare the message to be sent
	message := []byte(whatToSend)

	// Send the message to the WebSocket service
	resp, err := http.Post(os.Getenv("WS_SERVICE_URL"), "application/json", bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}
