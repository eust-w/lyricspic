package main

func SortByValue(in map[string]int) [][]interface{} {
	inLen := len(in)
	out := make([][]interface{}, inLen, inLen)
	//fmt.Println(out[3][1].(int))
	index := 0
	for k, v := range in {
		for i := 0; i <= index; i++ {
			if out[i] == nil || len(out[i]) != 2 {
				out[i] = []interface{}{v, k}
				break
			}
			if out[i][0].(int) <= v {
				copyOut := OutCopy(out)
				out = append(out[:i], []interface{}{v, k})
				out = append(out[:i+1], copyOut[i:]...)
				out = out[:inLen]
				break
			}
		}
		index++
	}
	return out
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMaps := DeepCopy(v)
			newMap[k] = newMaps
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlices := DeepCopy(v)
			newSlice[k] = newSlices
		}

		return newSlice
	}

	//fmt.Println(value.([][]interface{}))
	return value
}

func OutCopy(in [][]interface{}) [][]interface{} {
	newSlice := make([][]interface{}, len(in))
	for k0, v0 := range in {
		if v0 != nil {
			ele := make([]interface{}, 0, 2)
			ele = append(ele, v0[0])
			ele = append(ele, v0[1])
			newSlice[k0] = ele
		}
	}
	return newSlice
}
