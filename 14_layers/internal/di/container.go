package di

import (
	"14_layers/internal/db"
	"14_layers/internal/mail"
	"14_layers/internal/mapper"
	"14_layers/internal/repositories"
	"14_layers/internal/services"
	"14_layers/internal/services/observers"
	"gorm.io/gorm"
)

type ServiceContainer struct {
	dbConnection *gorm.DB

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
	s.dbConnection = db.GetConnection()

	emailSender := mail.NewEmailSender()
	orderObservers := []services.OrderObserver{
		observers.NewEmailObserver(emailSender),
	}

	s.orderRepository = repositories.NewOrderRepository(s.dbConnection)

	s.orderManager = services.NewOrderManager(s.orderRepository, orderObservers)
	s.orderMapper = mapper.NewOrderMapper()

	return s
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{}
}
