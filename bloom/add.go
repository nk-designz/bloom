package bloom

func (filter *Filter) Add(key ...string) {
	for _, k := range key {
		hashedKey := [][]byte{}
		for _, fn := range filter.Functions {
			hashedKey = append(hashedKey, fn([]byte(k)))
		}
		for _, hashBytes := range hashedKey {
			halfArr := (len(hashBytes)/2)-1;
			part := 0;
			for j := 0; j <= halfArr; j++ {
				part += int(hashBytes[j])
			}
			filter.Hash[part] = true;
			part = 0;
			for j := halfArr; j <= (len(hashBytes)-1); j++ {
				part += int(hashBytes[j])
			}
			filter.Hash[part] = true;
		}
	}
}