package di

import (
	"14_layers/internal/cache"
	"14_layers/internal/client"
	"14_layers/internal/config"
	"14_layers/internal/db"
	"14_layers/internal/mail"
	"14_layers/internal/mapper"
	"14_layers/internal/repositories"
	"14_layers/internal/services"
	"14_layers/internal/services/observers"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	dbConnection *gorm.DB

	cache cache.Cache

	orderRepository *repositories.OrderRepository

	orderManager *services.OrderManager
	orderMapper  *mapper.OrderMapper
}

func (s *ServiceContainer) DbConnection() *gorm.DB {
	return s.dbConnection
}

func (s *ServiceContainer) OrderRepository() *repositories.OrderRepository {
	return s.orderRepository
}

func (s *ServiceContainer) OrderManager() *services.OrderManager {
	return s.orderManager
}

func (s *ServiceContainer) OrderMapper() *mapper.OrderMapper {
	return s.orderMapper
}

func (s *ServiceContainer) Load() *ServiceContainer {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env("DB_HOST"),
		config.Env("DB_USER"),
		config.Env("DB_PASSWORD"),
		config.Env("DB_NAME"),
	)
	s.dbConnection = db.NewConnection(dsn)

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Env("REDIS_HOST"), config.Env("REDIS_PORT")),
	})
	s.cache = cache.NewRedisCache(redisClient)

	userClient := client.NewUserClient(s.cache)
	userService := services.NewUserService(userClient)
	emailSender := mail.NewEmailSender()
	orderObservers := []services.OrderObserver{
		observers.NewEmailObserver(emailSender, userService),
	}

	s.orderRepository = repositories.NewOrderRepository(s.dbConnection)

	s.orderManager = services.NewOrderManager(s.orderRepository, orderObservers)
	s.orderMapper = mapper.NewOrderMapper()

	return s
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{}
}
