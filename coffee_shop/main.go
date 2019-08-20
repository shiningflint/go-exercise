package main

import (
	"fmt"
	"strconv"
	"time"
)

type customer struct {
	name string
}

func main() {
	var queue []customer

	sc := generateCustomers(10)
	queueCustomers(sc, &queue)

	fmt.Println("All customers queued!", queue)
}

func generateCustomers(num int) []customer {
	var sc []customer
	for i := 0; i < num; i++ {
		newc := customer{name: "Mister-" + strconv.Itoa(i)}
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
