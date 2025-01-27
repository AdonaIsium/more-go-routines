package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("chulsu")
	fmt.Println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))
}

func getUserFriends(id string, ch chan <- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)

	ch <- &Message {
			friends: []string{
			"chulsu",
			"sungmin",
			"jeonging",
			"dahyun",
			"jeongyeon",
			"mina",
		},
	}
	wg.Done()
}

func getUserChats(id string, ch chan <- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)

	ch <- &Message { 
			chats: []string{
			"chulsu",
			"sungmin",
			"jeonging",
		},
	}
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}
