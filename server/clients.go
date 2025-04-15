package server

import (
	"context"
	"log"
	"sync"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/zuadi/tecamino-dbm.git/models"
)

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
)

type Client struct {
	conn *websocket.Conn
	ctx  context.Context
}

func registerClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[c] = true
	log.Printf("Client connected (%d total)", len(clients))
}

func unregisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, c)
	log.Printf("Client disconnected (%d total)", len(clients))
}

func broadcast(data models.JsonData) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for c := range clients {
		go func(client *Client) {
			err := wsjson.Write(client.ctx, client.conn, data)
			if err != nil {
				log.Printf("Broadcast error: %v", err)
			}
		}(c)
	}
}
