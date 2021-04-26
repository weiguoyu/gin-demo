package sliceutil

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func IsSlice(s interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(s)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

func TransferInterfaceToSlice(slice interface{}) ([]interface{}, bool) {
	val, ok := IsSlice(slice)
	if !ok {
		return nil, false
	}
	sliceLen := val.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = val.Index(i).Interface()
	}
	return out, true
}

func StringSliceContains(slice []string, val string) bool {
	for _, el := range slice {
		if strings.EqualFold(strings.TrimSpace(el), strings.TrimSpace(val)) {
			return true
		}
	}
	return false
}

func SortStringNumberSlice(slice []string) []string {
	if len(slice) < 2 {
		return slice
	}
	sort.SliceStable(slice, func(i, j int) bool {
		numA := fmt.Sprintf("%012v", slice[i])
		numB := fmt.Sprintf("%012v", slice[j])
		if numA > numB {
			return false
		} else {
			return true
		}
	})
	return slice
}
