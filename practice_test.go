package test

import (
	"fmt"
	"testing"
)

type Payment struct {
	amount int
}

func Test1(t *testing.T) {
	p1 := Payment{
		amount: 20,
	}
	p2 := p1

	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	p1.amount = 30
	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	p3 := &p1
	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	fmt.Printf("payemnt3 is %v\n", p3)
	p3.amount = 40
	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	fmt.Printf("payemnt3 is %v\n", p3)
	p1.amount = 50
	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	fmt.Printf("payemnt3 is %v\n", p3)

	SetPayment(&p2)
	fmt.Printf("payemnt1 is %v\n", p1)
	fmt.Printf("payemnt2 is %v\n", p2)
	fmt.Printf("payemnt3 is %v\n", p3)
}

func SetPayment(p *Payment) {
	p.amount = 100
	fmt.Printf("payemn2 setPayment is %v\n", p)
}

func TestVar(t *testing.T) {
	num := 10
	SetVariable(&num)
	fmt.Printf("testvar is %d\n", num)
}

func SetVariable(n *int) {

	fmt.Printf("setvar is %d\n", n)
}

