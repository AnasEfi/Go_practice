package hello

import (
	"fmt"

	"greetings"
)

func main() {
	message := greetings.Hello("Kostya")
	fmt.Println(message)
}
