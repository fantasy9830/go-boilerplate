package mqtt

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/websocket"
)

// Init Init
func Init() (err error) {
	client, err = NewClient(config.App.Name)
	if err != nil {
		return
	}

	client.Subscribe("#", 0, client.Message)

	// Broadcast Test
	go func() {
		manager := websocket.GetManager()

		for {
			select {
			case msg := <-client.Message:
				manager.Broadcast(msg.Payload())
			}
		}
	}()

	return
}
