package cache

import (
	"Wb-L0/internal/model"
	"errors"
	"fmt"
	"sync"
)

type Cache map[string]string

func New() *Cache {
	return &Cache{}
}

var mutex = &sync.Mutex{}

func (c *Cache) Add(order *model.Order) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := (*c)[order.UUID]; ok {
		return errors.New(fmt.Sprintf("UUID: %s already exist", order.UUID))
	}
	(*c)[order.UUID] = order.OrderJson
	return nil
}

func (c *Cache) GetByUUID(uuid string) string {
	return (*c)[uuid]
}
