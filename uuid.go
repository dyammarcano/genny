package main

import (
	"fmt"
	"github.com/google/uuid"
)

//Usage: uuid - generate and print a UUID

func main() {
	fmt.Printf("UUID: %s\n", uuid.New().String())
}
