package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxscore = 11

type game struct {
	score map[string]int
	ch    chan string
}

func (g *game) completed() bool {
	for _, score := range g.score {
		if score >= maxscore {
			return true
		}
	}

	return false
}

func main() {
	game := game{
		score: make(map[string]int),
		ch:    make(chan string),
	}
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go play("Player 1", &game, wg)
	go play("Player 2", &game, wg)
	game.ch <- "begin"

	wg.Wait()
	fmt.Println("Game finish!")
	for player, score := range game.score {
		fmt.Printf("%s: %d\n", player, score)
	}
}

func play(name string, game *game, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		response := "ping"
		request, ok := <-game.ch

		if !ok {
			break
		}

		if game.completed() {
			close(game.ch)
			break
		}

		if request == "ping" {
			response = "pong"
		}

		if rundWin() {
			response = "stop"
			game.score[name]++
		}

		fmt.Printf("Name: %s - Step: %s\n", name, response)
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
