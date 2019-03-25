package main

import (
	"flag"
	"fmt"
	"os"
)

//工厂方法只生产一种产品，抽象工厂生产多种类型的


//product type one: hamburger
type Hamburger interface {
	Deliver()
}

type KfcHamburger struct{}
func (h KfcHamburger) Deliver() {
	fmt.Println("This is a hamburger from KFC.")
}

type McdonaldHamburger struct{}
func (h McdonaldHamburger) Deliver() {
	fmt.Println("This is a hamburger from McDonalds.")
}


// product type two: chicken wing
type ChickenWing interface {
	Deliver()
}

type KfcChickenWing struct{}

func (wing KfcChickenWing) Deliver() {
	fmt.Println("This is a chicken wing from KFC.")
}

type McdonaldChickenWing struct{}

func (wing McdonaldChickenWing) Deliver() {
	fmt.Println("This is a chicken wing from McDonalds.")
}


// abstract factory
type FastFoodFactory interface {
	CreateHamburger() Hamburger
	CreateChickenWing() ChickenWing
}

// concrete factory one: KFC
type Kfc struct{}

func (factory Kfc) CreateHamburger() Hamburger {
	return new(KfcHamburger)
}
func (factory Kfc) CreateChickenWing() ChickenWing {
	return new(KfcChickenWing)
}

// concrete factory two: McDonald
type Mcdonald struct{}

func (factory Mcdonald) CreateHamburger() Hamburger {
	return new(McdonaldHamburger)
}
func (factory Mcdonald) CreateChickenWing() ChickenWing {
	return new(McdonaldChickenWing)
}



func main() {
	prefer := GetPreferredFastFood()

	var factory FastFoodFactory
	switch prefer {
	case "KFC":
		factory = new(Kfc)
	case "McDonalds":
		factory = new(Mcdonald)
	default:
		fmt.Printf("%s not supported yet.\n", prefer)
		os.Exit(1)
	}

	hamburger := factory.CreateHamburger()
	chickenWing := factory.CreateChickenWing()

	hamburger.Deliver()
	chickenWing.Deliver()
}

func GetPreferredFastFood() string {

	var prefer string
	flag.StringVar(&prefer, "prefer", "KFC", "preferred fast food type")
	flag.Parse()

	if len(os.Args) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	return prefer
}

