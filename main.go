package main

import (
	"fmt"
	"gochip/internal/emulator"
	"gochip/internal/frontend"
	"log"
	"os"
	"strconv"
	"time"
)

func StartEmulator(em *emulator.Emulator, delay time.Duration) {
	// Wait a bit to allow frontend to start
	time.Sleep(3 * time.Second)

	lastCycleTime := time.Now()
	for {
		currentTime := time.Now()
		dt := currentTime.Sub(lastCycleTime)

		time.Sleep(delay - dt)

		lastCycleTime = time.Now()
		err := em.Cycle()

		if err != nil {
			panic(err)
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: %s <delay> <rom>\n", args[0])
		os.Exit(1)
	}

	delay, err := strconv.Atoi(args[1])

	if err != nil {
		log.Fatalf("Expected delay as number, got: %s\n", args[1])
	}

	filepath := args[2]

	em := emulator.NewEmulator()
	err = em.LoadROM(filepath)

	if err != nil {
		log.Fatalf("Failed to load %v: %v", filepath, err)
	}

	go StartEmulator(&em, time.Duration(delay)*time.Millisecond)

	frontend.Start(em.Screen[:], emulator.SCREEN_WIDTH, emulator.SCREEN_HEIGHT, 10, em.Keypad[:])
}
