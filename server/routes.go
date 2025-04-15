package server

import "github.com/gin-gonic/gin"

func (s *Server) AddRoutes() {
	s.engine.GET("/json_data", func(c *gin.Context) {
		handleWebSocket(c.Writer, c.Request)
	})

	s.engine.GET("/", func(c *gin.Context) {
		c.String(200, "WebSocket Broadcast Server is running!")
	})
}
