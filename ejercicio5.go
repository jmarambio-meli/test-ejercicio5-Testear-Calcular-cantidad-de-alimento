package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalDog, msg := animal(dog)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalCat, msg := animal(cat)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalHamster, msg := animal(hamster)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	animalTarantula, msg := animal(tarantula)
	if msg != nil {
		fmt.Println(msg.Error())
	}

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	amount += animalHamster(5)
	amount += animalTarantula(8)

	fmt.Println("La cantidad de comida que debe comprar es de:", amount, "Kg")

}

func animal(animal string) (func(cantidad float64) float64, error) {
	switch animal {
	case dog:
		return AnimalDog, nil
	case cat:
		return AnimalCat, nil
	case hamster:
		return AnimalHamster, nil
	case tarantula:
		return AnimalTarantula, nil
	default:
		return nil, errors.New("animal no existe")
	}
}

func AnimalDog(cantidad float64) float64 {
	return cantidad * 10
}

func AnimalCat(cantidad float64) float64 {
	return cantidad * 5
}

func AnimalHamster(cantidad float64) float64 {
	return (cantidad * 250) / 1000 //conversion a KG
}

func AnimalTarantula(cantidad float64) float64 {
	return (cantidad * 150) / 1000 // conversion a KG
}
