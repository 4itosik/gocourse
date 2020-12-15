package main

import (
	"fmt"
	"math/rand"
	"time"
)

type game struct {
	winner string
	ch     chan string
}

func main() {
	game := game{
		winner: "",
		ch:     make(chan string),
	}
	done := make(chan string)

	go play("Player 1", &game, done)
	go play("Player 2", &game, done)
	game.ch <- "begin"

	<-done
	fmt.Println("Game finish!")
	fmt.Printf("Won: %s\n", game.winner)
}

func play(name string, game *game, done chan<- string) {
	for {
		response := "ping"
		request, ok := <-game.ch

		if !ok {
			break
		}

		if request == "stop" {
			close(game.ch)
			close(done)
			break
		}

		if request == "ping" {
			response = "pong"
		}

		if rundWin() {
			response = "stop"
			game.winner = name
		}

		fmt.Printf("Name: %s - Step: %s\n", name, response)
		time.Sleep(1 * time.Second)
		game.ch <- response
	}
}

func rundWin() bool {
	v := rand.Intn(100)
	if v <= 20 {
		return true
	}

	return false
}
