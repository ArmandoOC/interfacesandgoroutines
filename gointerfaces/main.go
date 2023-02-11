package main

import (
	"fmt"

	"github.com/ArmandoOC/interfacesandgoroutines/gointerfaces/users"
)

func main() {
	//armando := users.NewEmployee("Armando")
	//total := armando.CalcTotal(90.4, 100, 100)
	//fmt.Println(total)

	var lopez users.Cashier = users.NewEmployee("Lopez")
	var ramirez users.Admin = users.NewEmployee("Ramirez")

	//Al declarar que armando es de tipo Cashier, sólo tiene
	//acceso al método CalcTotal y no a todas las propiedades.
	total := lopez.CalcTotal(90.4, 100, 100)
	fmt.Println(total)

	ramirez.DeactivateEmployee(lopez)
	total2 := lopez.CalcTotal(90.4, 100, 100)
	fmt.Println(total2)
}

//Cada uno de los usuarios tiene acceso sólo a lo que nosotros le indicamos

//https://www.youtube.com/watch?v=gVzGJ5SlobI
