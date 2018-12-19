package test

import (
	"fmt"
	"testing"
)

type Developer interface {
	DrinkCoffee()
}

type GolangDeveloper struct {
	versionsIKnow []int
	golangCert []string
}

type JSDeveloper struct {
	versionsIKnow []int
	golangCert []string
}

func (g GolangDeveloper) DrinkCoffee() {
	fmt.Println("I drink 5 cups of coffee a day")
}

func (g JSDeveloper) DrinkCoffee() {
	fmt.Println("I drink 0 cups of coffee a day")
}

func TestInterface(t *testing.T) {
	d1 := GolangDeveloper{}
	d2 := JSDeveloper{}
	PrintDeveloperCoffeeCount(d1)
	PrintDeveloperCoffeeCount(d2)
}

func PrintDeveloperCoffeeCount(dev Developer) {
	dev.DrinkCoffee()
}