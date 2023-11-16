package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-boilerplate/internal/auth/config"
	"go-boilerplate/internal/auth/database"
	"go-boilerplate/internal/auth/entity"
	"go-boilerplate/internal/auth/repository/postgres"
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/middleware"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

var (
	once    sync.Once
	jwt     auth.JWTer
	userSvc middleware.UserService[entity.User]
)

type userService struct {
	userRep entity.UserRepository
	rdb     redis.UniversalClient
}

func InitialUserService() middleware.UserService[entity.User] {
	rdb, err := database.GetRedis()
	if err != nil {
		slog.Error("failed to initialize Redis", "err", err)
		os.Exit(1)
	}

	return NewUserService(postgres.InitialUserRepository(), rdb)
}

func NewUserService(userRep entity.UserRepository, rdb redis.UniversalClient) *userService {
	return &userService{
		userRep: userRep,
		rdb:     rdb,
	}
}

func (s *userService) Find(ctx context.Context, id string) (user entity.User, err error) {
	key := fmt.Sprintf("auth:user:%s", id)
	if err := s.rdb.Get(ctx, key).Scan(&user); err == nil {
		return user, nil
	}

	var g singleflight.Group
	value, err, _ := g.Do(id, func() (interface{}, error) {
		go func() {
			time.Sleep(1 * time.Second)
			g.Forget(id)
		}()

		u, err := s.userRep.FindFirst(ctx, "id", id)
		if err != nil {
			return nil, err
		}

		if err := s.rdb.Set(ctx, key, *u, 2*time.Hour).Err(); err != nil {
			return nil, err
		}

		return *u, nil
	})
	if err != nil {
		return user, err
	}

	if u, ok := value.(entity.User); ok {
		return u, nil
	}

	return user, errors.New("record not found")
}

func AuthRequired() gin.HandlerFunc {
	once.Do(func() {
		jwt = auth.NewJWT(config.App.Key)
		userSvc = InitialUserService()
	})

	return middleware.AuthRequired[entity.User](jwt, userSvc)
}
