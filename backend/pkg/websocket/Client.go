package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			pool.broadcast(Message{Type: 1, Body: "New User Joined..."})

		case client := <-pool.Unregister:
			if _, ok := pool.Clients[client]; ok {
				delete(pool.Clients, client)
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				pool.broadcast(Message{Type: 1, Body: "User Disconnected..."})
			}

		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			pool.broadcast(message)
		}
	}
}

// broadcast sends msg to every client. A client whose connection fails is
// removed and closed, and delivery continues to everyone else.

func (pool *Pool) broadcast(msg Message) {
	for c := range pool.Clients {
		if err := c.Conn.WriteJSON(msg); err != nil {
			log.Println("write failed, removing client:", err)
			delete(pool.Clients, c) // deleting during range is safe in Go
			c.Conn.Close()
		}
	}
}
