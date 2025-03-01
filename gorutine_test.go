package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGorutine(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
