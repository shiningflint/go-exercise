package main

import (
	"fmt"
	"strconv"
	"time"
)

type customer struct {
	name string
}

type milk struct {
	qty int
}

type espresso struct {
	qty int
}

type latte struct {
	name  string
	ratio int
}

func newLatte(mi milk, es espresso) latte {
	return latte{name: "a_new_latte", ratio: mi.qty + es.qty}
}

func main() {
	queue := generateCustomers(10)
	fmt.Println("All customers queued!", len(queue))

	fmt.Println("Placing one coffee order:")
	l := orderLattte()
	fmt.Println("got latte", l)
}

func generateCustomers(num int) []customer {
	var sc []customer
	for i := 0; i < num; i++ {
		newc := customer{name: "Mister customer-" + strconv.Itoa(i)}
		sc = append(sc, newc)
	}
	return sc
}

func queueCustomers(sc []customer, mc *[]customer) {
	for _, c := range sc {
		fmt.Printf("Queuing %s\n", c.name)
		time.Sleep(100 * time.Millisecond)

		*mc = append(*mc, c)
		fmt.Println("queued")
	}
}

func orderLattte() latte {
	mi := milk{qty: 70}
	es := espresso{qty: 30}
	cof := newLatte(mi, es)
	return cof
}
