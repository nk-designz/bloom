package bloom

import(
	"fmt"
	"math/rand"
	"testing"
)

var(
	set = []string{"Lorem","ipsum","dolor","sit","amet","consectetur","adipiscing","elit","sed","do","eiusmod","tempor","incididunt","ut","labore","et","dolore","magna","aliqua","Ut","enim","ad","minim","veniam,","quis","nostrud","exercitation","ullamco","laboris","nisi","ut","aliquip","ex","ea","commodo","consequat","Duis","aute","irure","dolor","in","reprehenderit","in","voluptate","velit","esse","cillum","dolore","eu","fugiat","nulla","pariatur","Excepteur","sint","occaecat","cupidatat","non","proident","sunt","in","culpa","qui","officia","deserunt","mollit","anim","id","est","laborum"}
)

func TestAdd(t *testing.T) {
	filter := Init(4096, Murmur, Spooky, Blake3, Xxhash)
	ran := rand.Int31n(int32(len(set)-1))
	probe := set[ran]
	filter.Add(probe)
	if(!filter.Probe(probe)) {
		t.Errorf("Failed to probe for %s", probe)
	}
}

func BenchmarkLoremrand(b *testing.B) {
	result := map[bool]int{}
	test := 0
	filter := Init(4096, Murmur, Spooky, Blake3, Xxhash)
	for i := 0; i < b.N; i++ {
		ran := rand.Int31n(int32(len(set)-1))
		probe := set[ran]
		if(ran % 2 == 0) {
			filter.Add(probe)
		} else {
			test++
		}
		result[filter.Probe(probe)]++
	}
	fmt.Printf("\nHits: %d\tMissed: %d\tTraps: %d\tTotal: %d\n", result[true], -(result[false]-test), test, result[true]+test)
}

func BenchmarkAddProbe(b *testing.B) {
	filter := Init(4096, Murmur, Spooky, Blake3, Xxhash)
	misses := 0
	for i := 0; i < b.N; i++ {
		ran := rand.Int31n(int32(len(set)-1))
		probe := set[ran]
		filter.Add(probe)
		if(!filter.Probe(probe)) {
			misses++
		}
	}
	fmt.Printf("\nMissed: %d\n", misses)
}

func BenchmarkAdd(b *testing.B) {
	filter := Init(4096, Murmur, Spooky, Blake3, Xxhash)
	for i := 0; i < b.N; i++ {
		ran := rand.Int31n(int32(len(set)-1))
		probe := set[ran]
		filter.Add(probe)
	}
}