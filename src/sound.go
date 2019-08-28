package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 21
	pin = rpio.Pin(21)
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	pin.Input()
	pin.PullUp()
	pin.Detect(rpio.AnyEdge) // enable any edge event detection

	fmt.Println("Searching for sound...")

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pin.Detect(rpio.NoEdge) // disable edge event detection
		os.Exit(0)
	}()

	for {
		if pin.EdgeDetected() { // check if event occured
			fmt.Println("Sound detected....\n")
		}
		time.Sleep(time.Second / 2)
	}
	pin.Detect(rpio.NoEdge) // disable edge event detection
}