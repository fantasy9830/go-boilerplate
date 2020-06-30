package app

import (
	"go-boilerplate/pkg/mqtt"
	"go-boilerplate/pkg/websocket"
)

// Init Init
func Init() error {
	// init mqtt
	if err := mqtt.Init(); err != nil {
		return err
	}

	// init websocket
	if err := websocket.Init(); err != nil {
		return err
	}

	client := mqtt.GetClient()
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

	return nil
}
