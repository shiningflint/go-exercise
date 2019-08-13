package main

import "fmt"

type tripPreparer interface {
	prepareTrip()
}

type mechanic struct {
	shopName string
}

type driver struct {
	drivingExp int
}

func (m mechanic) prepareTrip() {
	fmt.Println("The mechanic of", m.shopName, "is preparing a trip.")
}

func (d driver) prepareTrip() {
	fmt.Println("The driver, who has", d.drivingExp, "driving experience is preparing a trip")
}

func newTrip(tp ...tripPreparer) {
	fmt.Println("We are going to start a new trip!")
	for _, preparer := range tp {
		preparer.prepareTrip()
	}
}

func main() {
	m := mechanic{shopName: "Bambang Cycles"}
	d := driver{drivingExp: 90}
	newTrip(m, d)
}
