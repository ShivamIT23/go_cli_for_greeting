package greet

import "fmt"

func Hello(name string) string {
	if name == ""{
		return "Hello Stranger!"
	}
	return fmt.Sprintf("Hello %v",name)
}

	
