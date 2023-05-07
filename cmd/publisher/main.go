package main

import (
	"Wb-L0/internal/publisher"
)

func main() {
	if err := publisher.Start(); err != nil {
		panic(err)
	}
}
