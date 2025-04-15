package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/zuadi/tecamino-dbm.git/models"
)

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
		var data models.JsonData
		err := wsjson.Read(ctx, conn, &data)
		if err != nil {
			log.Println("Read error:", err)
			var response struct {
				Code    int    `json:"errorCode"`
				Message string `json:"message"`
				Error   bool   `json:"error"`
			}
			response.Code = 404
			response.Error = true
			response.Message = err.Error()

			err = wsjson.Write(ctx, conn, response)
			if err != nil {
				log.Println("Read error:", err)
			}
			break
		}

		log.Printf("Received: %v", data)

		// Broadcast to all
		broadcast(data)
	}

	conn.Close(websocket.StatusNormalClosure, "Normal closure")
}
