package controllers

import "di/services"

type PropertyInjectedController struct {
	// public
	GS services.GreetingService
}

func (m *PropertyInjectedController) getGreeting() string {
	return m.GS.Greet()
}

type SetterInjectedController struct {
	// private
	gs services.GreetingService
}

func (s *SetterInjectedController) setGreetingService(service services.GreetingService) {
	s.gs = service
}

func (s *SetterInjectedController) getGreeting() string {
	return s.gs.Greet()
}

// ConstructorInjectedController :
// fuck! there is no such thing as final or constructor in golang
// but let us mimic what it is like in "Jawa"
type ConstructorInjectedController struct {
	gs services.GreetingService
}

func NewConstructorInjectedController(s services.GreetingService) *ConstructorInjectedController {
	return &ConstructorInjectedController{
		gs: s,
	}
}

func (s *ConstructorInjectedController) getGreeting() string {
	return s.gs.Greet()
}
