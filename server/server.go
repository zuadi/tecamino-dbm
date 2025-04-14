package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

type Client struct {
	conn *websocket.Conn
	ctx  context.Context
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
)

func NewServer() *Server {
	s := Server{}

	s.engine = gin.Default()

	s.engine.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c.Writer, c.Request)
	})

	s.engine.GET("/", func(c *gin.Context) {
		c.String(200, "WebSocket Broadcast Server is running!")
	})
	return &s
}

func (s *Server) Serve(port uint) error {
	if err := s.engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("failed to run server: %v", err)
	}
	return nil
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		log.Println("WebSocket accept error:", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "Internal error")

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Minute)
	defer cancel()

	client := &Client{conn: conn, ctx: ctx}

	// Register client
	registerClient(client)
	defer unregisterClient(client)

	// Read loop
	for {
		var msg string
		err := wsjson.Read(ctx, conn, &msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s", msg)

		// Broadcast to all
		broadcast("Broadcast: " + msg)
	}

	conn.Close(websocket.StatusNormalClosure, "Normal closure")
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

func broadcast(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for c := range clients {
		go func(client *Client) {
			err := wsjson.Write(client.ctx, client.conn, message)
			if err != nil {
				log.Printf("Broadcast error: %v", err)
			}
		}(c)
	}
}
