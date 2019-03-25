package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignals(signal os.Signal) {
	fmt.Println("Got", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			sig := <-sigs
			fmt.Println(sig)
			handleSignals(sig)
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep( 10 * time.Second)
	}
}
