package main

import (
   "os"
   "bufio"
   "strings"
   "fmt"
   "go_cli_for_greeting/greet"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your name?")

	input , err := reader.ReadString('\n')

	if err !=nil {
		fmt.Println("An error occured, maybe your name is dangerous", err)
		return
	}

	name := strings.TrimSpace(input)

	message := greet.Hello(name)
	fmt.Println(message) 
}
