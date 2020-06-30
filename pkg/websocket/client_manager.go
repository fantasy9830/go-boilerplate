package websocket

import (
	"encoding/json"
	"sync"
)

var manager *ClientManager

// ClientManager ClientManager
type ClientManager struct {
	sync.Mutex
	Clients map[*Client]bool
}

// GetManager GetManager
func GetManager() *ClientManager {
	return manager
}

// Register Register
func (m *ClientManager) Register(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.Clients[client] = true
}

// Unregister Unregister
func (m *ClientManager) Unregister(client *Client) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Clients[client]; ok {
		close(client.ReadMessage)
		close(client.WriteMessage)
		delete(m.Clients, client)
	}
}

// Broadcast Broadcast
func (m *ClientManager) Broadcast(message []byte) {
	m.Lock()
	defer m.Unlock()
	for client := range m.Clients {
		select {
		case client.WriteMessage <- message:
		default:
			close(client.ReadMessage)
			close(client.WriteMessage)
			delete(m.Clients, client)
		}
	}
}

// SendMessage send ws message to ws client
func (m *ClientManager) SendMessage(message interface{}, id uint) {
	jsonMessage, _ := json.Marshal(message)
	for client := range m.Clients {
		if client.UserID == id {
			client.WriteMessage <- jsonMessage
		}
	}
}
