	package main

	import (
		"fmt"
		"math/rand"
		"time"
	)

	func main() {
		rand.Seed(time.Now().UnixNano())

		secretNumber := rand.Intn(100) + 1
		attempts := 0
		maxAttempts := 10

		fmt.Println("Welcome to the Smart Guessing Game!")
		fmt.Println("I have picked a number between 1 and 100. Can you guess it?")

		for attempts < maxAttempts {
			attempts++

			var guess int
			fmt.Printf("Attempt %d/%d — Enter your guess: ", attempts, maxAttempts)
			_,err :=fmt.Scan(&guess)

			if err != nil{
				fmt.Println("Invalid input. Please enter a number.")
				attempts-- 
			}

			if guess < secretNumber {
				fmt.Println("Too low! Try again.")
			} else if guess > secretNumber {
				fmt.Println("Too high! Try again.")
			} else {
				fmt.Printf("Congratulations! You've guessed the number %d in %d attempts!\n", secretNumber, attempts)
				return
			}
		}
		fmt.Println("Game Over! You've used all your attempts. The secret number was:", secretNumber)
	}
