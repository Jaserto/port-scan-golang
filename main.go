package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("Port Scanner")

	open := port.ScanPort("tcp", "localhost", 1314)
	fmt.Printf("Port Open: %t\n", open)

}
