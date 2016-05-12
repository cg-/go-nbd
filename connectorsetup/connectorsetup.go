package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cg-/go-nbd"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	asd, err := nbd.CreateNbdConnector("/tmp/test", "dest")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}


	fmt.Println("Mounted")
	asd.Mount()
	time.Sleep(5 * time.Second)
	fmt.Println("Unmounted")
	asd.Unmount()
	wg.Wait()
}

// End of file.
