package main

import(
	"log"
	"strings"
	"math/rand"
	//"time"
	"../bloom"
)

const(
	lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

func test(runs int) {
	set := strings.Split(lorem, " ")
	result := map[bool]int{}
	test := 0
	filter := bloom.Init(4096, bloom.Murmur, bloom.Spooky, bloom.Blake3, bloom.Xxhash)
	for i := 0; i < runs; i++ {
		ran := rand.Int31n(int32(len(set)-1))
		probe := set[ran]
		if(ran % 2 == 0) {
			filter.Add(probe)
		} else {
			test++
		}
		result[filter.Probe(probe)]++
	}
	log.Printf("Hits: %d\tMissed: %d\tTraps: %d\tTotal: %d", result[true], -(result[false]-test), test, result[true]+test)
}

func main() {
	for {
		test(int(rand.Int31n(100000)))
		//time.Sleep(500 * time.Millisecond)
	}
}