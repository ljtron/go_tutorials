package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)
// Hello returns a message and includes the name with in it
func Hello(name string) (string, error){
	// returns an error if no name was filled
	if(name == ""){
		return "", errors.New("empty string")
	}

	// saves the response within the variable message
	message := fmt.Sprintf(randomFunc(), name)

	//another form of initialization
	// var message string
	// message = fmt.Sprintf(..., value)
	return message, nil
}

// Hellos returns a map of all the names and their messages
func Hellos(names []string) (map[string]string, error){

	// a map of all the names with messages
	// map: is the equivalent of a python dictionary or hashmap
	messages := make(map[string]string)

	// loop through the names and calls hello
	for i, name := range names{
		// just printing the index
		fmt.Println("loop through name index: %v", i)

		// storing the value of message
		message, err = Hello(name)
		if err != nil{
			return nil, err
		}

		// maps the name with the value
		messages[name] = message
	}
	return messages, nil
}

// randomFunc returns a random message out of a list of messages
func randomFunc() string{

	// list of potential responses
	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}

	// returns a random message from the list
	return formats[rand.Intn(len(formats))]
}