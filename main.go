package learnGoWithTests

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "Paul" {
		name = "World NOT LOL"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Paul"))
}
