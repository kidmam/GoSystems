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
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handleSignals(sig)
			case syscall.SIGTERM:
				handleSignals(sig)
				fmt.Println("Got:", sig)
				os.Exit(-1)
			case syscall.SIGUSR1:
				handleSignals(sig)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep( 10 * time.Second)
	}
}
