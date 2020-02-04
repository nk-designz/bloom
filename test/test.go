package main

import(
	"fmt"
)

func main() {
	filter := Init(4096, Murmur, Spooky, Blake3, Xxhash)
	fmt.Println(filter)
	filter.Add("peter","test")
	fmt.Println(filter)
	fmt.Println(filter.Probe("peter"))
	fmt.Println(filter.Probe("test"))
	fmt.Println(filter.Probe("cools"))
}