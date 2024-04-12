package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader         = websocket.Upgrader{}
	websocketClients = sync.Map{} // Map to store connected clients
)

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()

	websocketClients.Store(conn.RemoteAddr(), conn)
	defer websocketClients.Delete(conn.RemoteAddr())

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
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
	http.HandleFunc("/ws", websocketHandler)
	go startSendingMessages() // Start a goroutine to send messages periodically
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startSendingMessages() {

	for {
		// Send a message to all connected clients
		message := []byte("Hello from Go!")
		broadcastMessage(message)
		time.Sleep(5 * time.Second)
	}
}
