package php

import (
	"reflect"
	"sort"
)

// ArrayKeys array_keys returns the keys, numeric and string, from the array.
func ArrayKeys(array interface{}) interface{} {
	t, v, l := getCommon(array)
	res := make([]interface{}, l)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			res[i] = i
		}
	case reflect.Map:
		for i, k := range v.MapKeys() {
			res[i] = k.Interface()
		}
	default:
		panic("expects parameter 1 to be array")
	}
	return res
}

// ArrayValues array_values returns all the values from the array and indexes the array numerically.
func ArrayValues(array interface{}) interface{} {
	t, v, l := getCommon(array)
	res := make([]interface{}, l)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			res[i] = v.Index(i).Interface()
		}
	case reflect.Map:
		for i, k := range v.MapKeys() {
			res[i] = v.MapIndex(k)
		}
	default:
		panic("expects parameter 1 to be array")
	}
	return res
}

// ArrayKeyExists array_key_exists — Checks if the given key or index exists in the array
func ArrayKeyExists(key, array interface{}) bool {
	t, v, l := getCommon(array)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			if reflect.DeepEqual(key, i) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if reflect.DeepEqual(key, k.Interface()) {
				return true
			}
		}
	default:
		panic("expects parameter 2 to be array")
	}
	return false
}

// InArray in_array — Checks if a value exists in an array
func InArray(needle, haystack interface{}) bool {
	t, v, l := getCommon(haystack)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			if reflect.DeepEqual(needle, v.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if reflect.DeepEqual(needle, v.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("expects parameter 2 to be array")
	}
	return false
}

// ArrayFilp array_flip — Exchanges all keys with their associated values in an array
func ArrayFilp(array interface{}) map[interface{}]interface{} {
	t, v, l := getCommon(array)
	res := make(map[interface{}]interface{}, l)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			res[v.Index(i).Interface()] = i
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			res[v.MapIndex(k).Interface()] = k.Interface()
		}
	default:
		panic("expects parameter 1 to be array")
	}
	return res
}

// ArrayUnique array_unique — Removes duplicate values from an array
func ArrayUnique(array interface{}) interface{} {
	t, v, l := getCommon(array)
	res := make(map[interface{}]int)
	switch t.Kind() {
	case reflect.Slice:
		for i := 0; i < l; i++ {
			res[v.Index(i).Interface()] = 1
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			res[v.MapIndex(k).Interface()] = 1
		}
	default:
		panic("expects parameter 1 to be array")
	}
	return ArrayKeys(res)
}

// Sort can only sort []int, []string, []float64
func Sort(array interface{}) {
	t, v, _ := getCommon(array)
	// res := make([]interface{}, l)
	if t.Kind() == reflect.Slice {
		switch v.Index(0).Kind() {
		case reflect.Int:
			array := array.([]int)
			sort.Ints(array)
		case reflect.String:
			array := array.([]string)
			sort.Strings(array)
		case reflect.Float64:
			array := array.([]float64)
			sort.Float64s(array)
		default:
			panic("the param can only be int/string/float64 array")
		}
	} else {
		panic("expects parameter 1 to be array")
	}
}

func getCommon(array interface{}) (reflect.Type, reflect.Value, int) {
	t := reflect.TypeOf(array)
	v := reflect.ValueOf(array)
	l := v.Len()
	return t, v, l
}



// ArrayColum array_column — Return the values from a single column in the input array
func ArrayColum(s [] map[string] interface{}, string key) []interface{} {
	n := len(s)
	ret = make([]interface{}, n)
	for k,v := range s {
		ret[k] = s[k][key]
	}
	return ret
}

// ArrayChunk array_chunk — Split an array into chunks
func ArrayChunk(s []interface{}, size int) [][]interface{} {
	if (size < 1) {
		panic("Size parameter expected to be greater than 0")
	}
	n_ := len(s)
	n = n_ / size
	if !(n*size == n_) {
		n += 1
	}
	ret = make([][]interface{}, n)
	for i:=0; i<n-1; i++ {
		ret[i] = s[i*size:size]
	}
	ret[i] = s[i*size:]
	return ret
}
	
// ArrayDiff array_diff — Computes the difference of arrays
function ArrayDiff(a []string, b...[]string) {
	m := make(map[string]int, len(a))
	for k,v := range a {
		m[v] = 0
	}
	
	for _,_b := range b {
		for k,v := range _b {
			_,ok := m[v]
			if ok {
				m[v]++
			}
		}
	}
	
	ret := make([]string)
	for k,v := range m {
		if v == 0 {
			ret = append(ret, v)
		}
	}
}

// array_filter — Filters elements of an array using a callback function
func ArrayFilterSlice(a[]interface{}, f func(interface{}) bool) ret []interface{} {
	for _,v := range a {
		if f(v) {
			ret = append(ret, v)
		}
	}
}
func ArrayFilterMap(a map[string]interface{}, f func(interface{}) bool) ret [string]interface{} {
	for k,v := range a {
		if f(v) {
			ret[k] = v
		}
	}
}

// array_map — Applies the callback to the elements of the given arrays
func ArrayMapSlice(a[]interface{}, f func(interface{}) bool) ret []interface{} {
	ret = make([]interface{}, len(a))
	for k,v := range a {
		ret[k] = f(v)
	}
}
func ArrayMap(a map[string]interface{}, f func(interface{}) bool) ret [string]interface{} {
	ret = make(map[string]interface{})
	for k,v := range a {
		ret[k] = f(v)
	}
}

// ArrayCountValues array_count_values — Counts all the values of an array
func ArrayCountValues(a[]string) m map[string]int {
	m := make(map[string]int)
	for _,v := range a {
		m[v]++
	}
}

// ArrayReduce array_reduce — Iteratively reduce the array to a single value using a callback function
// ArrayReduce::func(a[]T, f func(T,T)T, init T) T
func ArrayReduce(a[]interface{}, f func(interface{}, interface{}) interface{}, init interface{}) s interface{} {
	s := init
	for _,v := range a {
		s = f(v, s)
	}
}
