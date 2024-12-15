package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)


func main(){
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	// stores the result from hello in the variables
	message, err := greetings.Hello("Lincoln Mcloud")

	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(message)
}