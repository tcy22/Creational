package main

import (
	"flag"
	"fmt"
	"os"
)

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

type HamburgerFactory struct{}

func (f HamburgerFactory) CreateHamburger(prefer string) Hamburger {
	switch prefer {
	case "KFC":
		return new(KfcHamburger)
	case "McDonald":
		return new(McdonaldHamburger)
	default:
		return nil
	}
}

func main(){
	prefer := getPreferHamburger()

	factory := new(HamburgerFactory)

	hamburger := factory.CreateHamburger(prefer)

	if hamburger == nil {
		fmt.Printf("%s not supported",prefer)
	}

	hamburger.Deliver()

}

func getPreferHamburger() string {

	var prefer string
	flag.StringVar(&prefer, "prefer", "KFC", "Preferenced Hamburger")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	return prefer

}