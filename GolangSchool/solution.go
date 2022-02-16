package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage(word string) string {
	return emoji.Sprintf(":%v:", word)
}

func GetBeer() string {
	return emoji.Sprint("Hello :beer:!")
}

func main() {
	s := []string{"beer", "world_map", "dog", "pizza", "sushi", "car",
		"eyes", "elevator", "female_vampire", "film_strip", "knot", "ram",
		"girl", "bank", "book", "heart"}
	for _, v := range s {
		fmt.Printf("%v\t", GetMessage(v))
	}

	beer := GetBeer()
	fmt.Println(beer)

}
