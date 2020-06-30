package websocket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Client Client
type Client struct {
	ID           string
	UserID       uint
	ws           *websocket.Conn
	WriteMessage chan []byte
	ReadMessage  chan []byte
}

// NewClient NewClient
func NewClient(ctx *gin.Context) (client *Client, err error) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}

	client = &Client{
		ID:           uuid.New().String(),
		ws:           ws,
		WriteMessage: make(chan []byte),
		ReadMessage:  make(chan []byte),
	}

	manager.Register(client)

	go client.ReadServe()
	go client.WriteServe()

	return
}

// SetUserID SetUserID
func (c *Client) SetUserID(id uint) *Client {
	c.UserID = id

	return c
}

// ReadServe ReadServe
func (c *Client) ReadServe() {
	defer func() {
		manager.Unregister(c)
		c.ws.Close()
	}()

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		c.ReadMessage <- message
	}
}

// WriteServe WriteServe
func (c *Client) WriteServe() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.WriteMessage:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}
