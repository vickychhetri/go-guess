package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("******************************** GUESS GAME *******************************")
	fmt.Print("Press Y to start Game: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error to start the game")
	}
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "y" {
		fmt.Println("GAME STARTED")
		target := int64(rand.Intn(100))
		matched := false
		for i := 0; i < 10; i++ {
			fmt.Print("Guess::  ")
			num, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Please choose Integer Number Only.")
				continue
			}
			num = strings.TrimSpace(num)
			fmt.Println(num)

			guess, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if guess == target {
				fmt.Printf("You won the game , choice :: %d vs Target :: %d Matched\n", guess, target)
				matched = true
				break
			} else if guess > target {
				fmt.Println("Oops, Your Guess is HIGH")
				continue
			} else if guess < target {
				fmt.Println("Oops, Your Guess is LOW")
				continue
			}
		}
		if !matched {
			fmt.Println("Sorry, You did not guess my number.")
		}

	} else {
		fmt.Println("Thanks ---- exiting game ")
	}

	fmt.Println("******************************** FINISH GUESS GAME *******************************")

}
