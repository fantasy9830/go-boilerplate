package websocket

// Init Init
func Init() error {
	manager = &ClientManager{
		Clients: make(map[*Client]bool),
	}

	return nil
}
