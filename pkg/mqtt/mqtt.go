package mqtt

import (
	"go-boilerplate/pkg/config"
)

// Init Init
func Init() (err error) {
	client, err = NewClient(config.App.Name)

	return
}
