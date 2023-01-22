package services

import "fmt"

type GreetingService interface {
	Greet() string
}

type GreetingServiceImpl struct {
}

func (g *GreetingServiceImpl) Greet() string {
	fmt.Println("Hello World")
	return "What a beautiful greeting"
}
