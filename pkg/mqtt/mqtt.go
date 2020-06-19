package mqtt

import "go-boilerplate/pkg/config"

// Init Init
func Init() (err error) {
	client, err = NewClient(config.App.Name)
	if err != nil {
		return
	}

	client.Subscribe("#", 0, client.Message)

	return
}
