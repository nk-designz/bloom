package bloom

import(
)

type Filter struct {
	Hash		[]bool	`json:"hash"`
	Functions	[]Function
}

func Init(size uint32, functions ...Function) Filter {
	return Filter {
		Hash:	make([]bool, size),
		Functions: functions}
}