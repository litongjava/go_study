package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	guid := uuid.New()
	guidStr := guid.String()
	fmt.Printf("guid: %s\n", guid)
	fmt.Printf("guidStr: %s\n", guidStr)
}
