package controllers

import (
	"di/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	propertyInjectedController    *PropertyInjectedController
	setterInjectedController      *SetterInjectedController
	constructorInjectedController *ConstructorInjectedController
)

func init() {
	propertyInjectedController = &PropertyInjectedController{}
	propertyInjectedController.GS = &services.GreetingServiceImpl{}

	setterInjectedController = &SetterInjectedController{}
	setterInjectedController.setGreetingService(&services.GreetingServiceImpl{})

	constructorInjectedController = NewConstructorInjectedController(&services.GreetingServiceImpl{})
}

func TestPropertyInjectedController(t *testing.T) {
	greetStr := propertyInjectedController.getGreeting()
	assert.Equal(t, "What a beautiful greeting", greetStr)
}

func TestSetterInjectedController(t *testing.T) {
	greetStr := setterInjectedController.getGreeting()
	assert.Equal(t, "What a beautiful greeting", greetStr)
}

func TestConstructorInjectedController(t *testing.T) {
	greetStr := constructorInjectedController.getGreeting()
	assert.Equal(t, "What a beautiful greeting", greetStr)
}
