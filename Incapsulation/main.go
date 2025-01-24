package main

import (
	"awesomeProject/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2006)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
}
