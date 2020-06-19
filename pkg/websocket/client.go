package websocket

import (
	"errors"
	"go-boilerplate/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// Error constants
var (
	ErrUserNotExist = errors.New("User Not Exist")
)

// Client Client
type Client struct {
	ID      string
	UserID  uint
	ws      *websocket.Conn
	Message chan []byte
}

// Message Message
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
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

	if user, ok := ctx.MustGet("user").(*models.User); ok {
		client = &Client{
			ID:      uuid.New().String(),
			UserID:  user.ID,
			ws:      ws,
			Message: make(chan []byte),
		}

		manager.Register(client)

		go client.ReadServe()
		go client.WriteServe()
	} else {
		return nil, ErrUserNotExist
	}

	return
}

// ReadServe ReadServe
func (c *Client) ReadServe() {
	defer func() {
		manager.Unregister(c)
		c.ws.Close()
	}()

	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			break
		}

		// jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		// manager.Broadcast <- jsonMessage
	}
}

// WriteServe WriteServe
func (c *Client) WriteServe() {
	defer func() {
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.ws.WriteMessage(websocket.TextMessage, message)
		}
	}
}
