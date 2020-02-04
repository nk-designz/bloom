package main

import(
	"fmt"
	"strings"
	"math/rand"
	"../bloom"
)

const(
	lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

func main() {
	set := strings.Split(lorem, " ")
	result := map[bool]int{}
	filter := bloom.Init(4096, bloom.Murmur, bloom.Spooky, bloom.Blake3, bloom.Xxhash)
	for i := 0; i <= 100; i++ {
		ran := rand.Int31n(int32(len(set)-1))
		probe := set[ran]
		if(ran % 3 == 1) {
			filter.Add(probe)
		} else {
			result[false]--
		}
		result[filter.Probe(probe)]++
	}
	fmt.Printf("Hits: %d\tMissed: %d\n", result[true], result[false])
}