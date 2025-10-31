package Container

import (
	"reflect"
)

type Service interface{}

type Container struct {
	pService map[string]*Service
}

func NewContainer() *Container {
	return &Container{
		pService: make(map[string]*Service),
	}
}

func (container *Container) Register(service Service) *Container {
	name := getServiceName(service)

	if ser := container.pService[name]; ser == nil {
		container.pService[name] = &service
	}

	return container
}

func getServiceName(service Service) string {
	t := reflect.TypeOf(service)

	return t.Elem().Name()
}

func (container *Container) Get(service Service) *Service {
	name := getServiceName(service)

	if ser := container.pService[name]; ser == nil {
		container.pService[name] = &service
	}

	return container.pService[name]
}
