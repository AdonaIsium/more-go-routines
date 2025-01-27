package main

import (
	"fmt"
	"time"
)

func main() {
	id := getUserByName("chulsu")
	fmt.Println(id)

	chats := getUserChats(id)
	friends := getUserFriends(id)

	fmt.Println(chats)
	fmt.Println(friends)
}

func getUserFriends(id string) []string {
	time.Sleep(time.Second * 1)

	return []string{
		"chulsu",
		"sungmin",
		"jeonging",
		"dahyun",
		"jeongyeon",
		"mina",
	}
}

func getUserChats(id string) []string {
	time.Sleep(time.Second * 2)

	return []string{
		"chulsu",
		"sungmin",
		"jeonging",
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}
