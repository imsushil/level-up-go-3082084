package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "friends.json"

// visited friends
type VisitedFriends struct {
	visited map[string]bool 
}

func (vf *VisitedFriends) isVisited(id string) bool {
	return vf.visited[id]
}

func (vf *VisitedFriends) setVisited(id string) {
	vf.visited[id] = true
}

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections
type Friends struct {
	fmap map[string]Friend
}

// getFriend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.fmap[id]
}

// getRandomFriend returns an random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.fmap)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

// spreadGossip ensures that all the friends in the map have heard the news
func spreadGossip(root Friend, friends Friends, visitedFriends VisitedFriends) {
	visitedFriends.setVisited(root.ID)
	for _, id := range root.Friends {
		if visitedFriends.isVisited(id) || root.ID == id {
			continue
		}
		friend := friends.getFriend(id)
		friend.hearGossip()
		spreadGossip(friend, friends, visitedFriends)
	}
}

func main() {
	friends, visitedFriends := importData()
	root := friends.getRandomFriend()
	root.hearGossip()
	spreadGossip(root, friends, visitedFriends)
}

// importData reads the input data from file and
// creates the friends map.
func importData() (Friends, VisitedFriends) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []Friend
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	fm := make(map[string]Friend, len(data))
	vf := make(map[string]bool, len(data))
	for _, d := range data {
		fm[d.ID] = d
		vf[d.ID] = false
	}

	return Friends{
		fmap: fm,
	}, VisitedFriends {
		visited: vf,
	}
}
