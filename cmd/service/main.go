package main

import (
	"Wb-L0/internal/service"
)

func main() {
	config := service.NewConfig()
	if err := service.Start(config); err != nil {
		panic(err)
	}
}
