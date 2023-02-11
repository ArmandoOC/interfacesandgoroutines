package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

//Una interfaz es la definición de un comportamiento de un struct. Es decir, un set de métodos que va a contener este struct.

type Cashier interface {
	CalcTotal(item ...float32) float32 //Uso externo
	deactivate()                       //Uso interno
}

//El cashier sólo quiero que calcule totales

type Admin interface {
	DeactivateEmployee(c Cashier)
}

//El admin sólo quiero que desactive a los empleados

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Cannot get total. Reason: Cashier deactivated")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}

	return sum * 1.15
}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}
