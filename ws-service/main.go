package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all origins for WebSocket connections
			return true
		},
	}
	websocketClients = sync.Map{} // Map to store connected clients

)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	websocketClients.Store(conn.RemoteAddr(), conn)
	defer websocketClients.Delete(conn.RemoteAddr())

	for {
		t, p, err := conn.ReadMessage()
		if err == nil {
			log.Println("read:", string(p), t, err)
			startSendingMessages(string(p))
			break
		}

	}

}

func broadcastMessage(message []byte) {
	websocketClients.Range(func(key, value interface{}) bool {
		client := value.(*websocket.Conn)
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("write:", err)
			return false
		}
		return true
	})
}
func main() {
	// Start the WebSocket server in a goroutine
	go func() {
		// Register the WebSocket handler
		http.HandleFunc("/ws", websocketHandler)

		// Start the WebSocket server
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Block the main goroutine to prevent the program from exiting
	select {}
}

func startSendingMessages(msg string) {
	// Send a message to all connected clients
	message := []byte(msg)
	broadcastMessage(message)
}
