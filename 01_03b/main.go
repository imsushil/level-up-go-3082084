package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
	"os"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	// TODO: Fill in definition
	Id string `json="id"`
	Name string `json="name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	reader, _ := os.Open("entries.json")
	decoder := json.NewDecoder(reader)
	_, err := decoder.Token() // check if format is valid
	if err != nil {
		panic("Invalid data")
	}
	rEntrySlice := make([]raffleEntry, 0)
	for decoder.More() {
		var rEntry raffleEntry
		err = decoder.Decode(&rEntry)
		if err != nil {
			panic("Error occurred while decoding data" + err.Error())
		}
		rEntrySlice = append(rEntrySlice, rEntry)
	}
	return rEntrySlice
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}