package helper

import (
	"reflect"
)

func Has (haystack interface{}, needle interface{}) bool {
	arr := reflect.ValueOf(haystack)
	length := arr.Len()

	if length == 0 {
		return false
	}

	if arr.Kind() == reflect.Slice {
		for i, j := 0, length-1 ; i <= j ; i, j = i+1, j-1 {
			hay1 := arr.Index(i).Interface()
			hay2 := arr.Index(j).Interface()

			if hay1 == needle || hay2 == needle {
				return true
			}
		}
	}

	return false
}
