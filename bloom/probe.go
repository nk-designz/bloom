package bloom

func (filter *Filter) Probe(key string) bool {
	hashedKey := [][]byte{}
	value := false;
	for _, fn := range filter.Functions {
		hashedKey = append(hashedKey, fn([]byte(key)))
	}
	for _, hashBytes := range hashedKey {
		halfArr := (len(hashBytes)/2)-1;
		part := 0;
		for j := 0; j <= halfArr; j++ {
			part += int(hashBytes[j])
		}
		value = filter.Hash[part] || value;
		part = 0;
		for j := halfArr; j <= (len(hashBytes)-1); j++ {
			part += int(hashBytes[j])
		}
		value = filter.Hash[part] || value;
	}
	return value
}
