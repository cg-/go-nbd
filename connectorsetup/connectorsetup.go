package main

import (
	"fmt"
	"os"

	"github.com/cg-/go-nbd"
)

func main() {
	asd, err := nbd.CreateNbdConnector("source", "dest")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(asd.Dump())
}

// End of file.
