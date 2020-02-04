package bloom

import(
	"lukechampine.com/blake3"
	"github.com/zond/god/murmur"
	"github.com/inspirent/go-spooky"
	"github.com/OneOfOne/xxhash"
)

type Function func([]byte) []byte

func Murmur(bytes []byte) []byte {
	return murmur.HashBytes(bytes)
}

func Blake3(bytes []byte) []byte {
	hashVal32 := blake3.Sum256(bytes)
	return hashVal32[:]
}

func Spooky(bytes []byte) []byte {
	seed := spooky.Hash64(bytes)
	spook := spooky.New(seed, seed)
	return spook.Sum(bytes)
}

func Xxhash(bytes []byte) []byte {
	return xxhash.New32().Sum(bytes)
}