package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

// Bot contains configuration, portfolio, exchange & ticker data and is the
// overarching type across this code base.
type FeedManager struct {
	shutdown chan bool
}

// TODO: Why are we using globals like a chump
var feedManager FeedManager

func main() {
	HandleInterrupt()
	AdjustGoMaxProcs()

	fmt.Println("The Feed")
	fmt.Println("========")

	//go WebsocketHandler()
	//go TickerUpdaterRoutine()
	//go OrderbookUpdaterRoutine()

	Shutdown()
}

// AdjustGoMaxProcs adjusts the maximum processes that the CPU can handle.
func AdjustGoMaxProcs() {
	log.Println("Adjusting bot runtime performance..")
	maxProcsEnv := os.Getenv("GOMAXPROCS")
	maxProcs := runtime.NumCPU()
	log.Println("Number of CPU's detected:", maxProcs)

	if maxProcsEnv != "" {
		log.Println("GOMAXPROCS env =", maxProcsEnv)
		env, err := strconv.Atoi(maxProcsEnv)
		if err != nil {
			log.Println("Unable to convert GOMAXPROCS to int, using", maxProcs)
		} else {
			maxProcs = env
		}
	}
	if i := runtime.GOMAXPROCS(maxProcs); i != maxProcs {
		log.Fatal("Go Max Procs were not set correctly.")
	}
	log.Println("Set GOMAXPROCS to:", maxProcs)
}

// HandleInterrupt monitors and captures the SIGTERM in a new goroutine then
// shuts down bot
func HandleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Captured %v.", sig)
		Shutdown()
	}()
}

// Shutdown correctly shuts down bot saving configuration files
func Shutdown() {
	log.Println("Feed shutting down..")
	os.Exit(1)
}
