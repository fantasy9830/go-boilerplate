package user

import (
	"go-boilerplate/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile Profile
type Profile struct {
	*models.User
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// Controller Controller
type Controller struct {
	serv *Service
}

// NewController New Controller
func NewController(service *Service) *Controller {
	return &Controller{
		serv: service,
	}
}

// Profile Profile
func (c *Controller) Profile(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	profile := Profile{
		User:        user,
		Roles:       make([]string, 0),
		Permissions: make([]string, 0),
	}

	keys := make(map[string]bool)
	for _, role := range user.Roles {
		profile.Roles = append(profile.Roles, role.Name)
		for _, permission := range role.Permissions {
			if _, isExist := keys[permission.Name]; !isExist {
				keys[permission.Name] = true
				profile.Permissions = append(profile.Permissions, permission.Name)
			}
		}
	}

	for _, permission := range user.Permissions {
		if _, isExist := keys[permission.Name]; !isExist {
			keys[permission.Name] = true
			profile.Permissions = append(profile.Permissions, permission.Name)
		}
	}

	ctx.JSON(http.StatusOK, profile)
	return
}
