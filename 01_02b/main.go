package main

import (
	"log"
	"time"
	"strings"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	// result := make([]string, 0)
	// result = append(result, string(msg[0]))
	slice := strings.Split(msg, " ")
	for _, w := range slice {
		pw := make([]string, 0)
		for i, wb := range w {
			rep := strings.Repeat(string(wb), i + 1)
			pw = append(pw, rep)
		}
		print(strings.Join(pw, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}