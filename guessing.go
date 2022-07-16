package main

import (
	"fmt"
	"math/rand"
	"strconv"
)



func main() {
	m := make(map[string]int)
	var response string
	randomNumber := strconv.Itoa(rand.Intn(9999 - 1000) + 1000)
	fmt.Println("danrdom: ", randomNumber)

	for i,c := range randomNumber{
		m[string(c)] = i
	}
	var guess string
	var count int
	for {
		fmt.Scanln(&guess)
		for i,c := range guess{
			if value, found := m[string(c)]; found{
				if value != i{
					response += "-1"
				}else{
					response +="+2"
					count++
				}
			}
			
		}
		if count == 4{
			fmt.Printf("response: +4")
			break
		}
		fmt.Printf("response: %s\n", response)	
		count = 0
		response = ""
	}
}
