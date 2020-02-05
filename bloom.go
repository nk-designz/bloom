package bloom

import(
)

type Filter struct {
	Hash		[]bool	`json:"hash"`
	Functions	[]HashFunction
}

func Init(size uint32, functions ...HashFunction) Filter {
	return Filter {
		Hash:	make([]bool, size),
		Functions: functions}
}