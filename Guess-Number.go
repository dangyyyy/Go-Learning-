package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбираю число от 1 до 100")
	fmt.Println("Число выбрано")

	reader := bufio.NewReader(os.Stdin)
	pobeda := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У вас осталось", 10-guesses, "попыток")
		fmt.Print("Напишите ваше число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal()
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal()
		}
		if guess > target {
			fmt.Println("Ваше число больше загадонного")
		} else if guess < target {
			fmt.Println("Ваше число меньше загадонного")
		} else {
			fmt.Println("Поздравляю! Вы победили!")
			pobeda = true
			break
		}

	}
	if !pobeda {
		fmt.Println("Вы проиграли :( Это число было:", target)
	}
}
