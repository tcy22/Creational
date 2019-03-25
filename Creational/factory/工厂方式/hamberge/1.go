package hamberge

import (
	"flag"
	"fmt"
	"os"
)

//产品
type Hamburger interface {
	Deliver()
}

type KfcHamburger struct{}

func (h KfcHamburger) Deliver() {
	fmt.Println("A hamburger from KFC")
}

type McDonaldHamburger struct{}

func (h McDonaldHamburger) Deliver() {
	fmt.Println("A hamburger from McDonald")
}


//工厂
type HamburgerFactory interface {
	Create() Hamburger
}

type Kfc struct{}

func (f Kfc) Create() Hamburger {
	return new(KfcHamburger)
}

type McDonald struct{}

func (f McDonald) Create() Hamburger {
	return new(McDonaldHamburger)
}


//客户端
func main(){
	prefer := getPreferHamburger()

	var factory HamburgerFactory

	switch prefer {
	case "KFC":
		factory = new(Kfc)
	case "McDonald":
		factory = new(McDonald)
	default:
		fmt.Printf("%s not supported",prefer)
	}

	hamburger := factory.Create()
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
