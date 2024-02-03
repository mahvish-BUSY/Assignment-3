package main

import (
	"fmt"
	"reflect"
)

// this func simply appends the arr to the result
func appendToResult(result []interface{}, arr interface{}) []interface{} {
	arrVal := reflect.ValueOf(arr)

	if arrVal.IsValid() {
		switch arrVal.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < arrVal.Len(); i++ {
				result = append(result, arrVal.Index(i).Interface())
			}
		default:
			result = append(result, arr)
		}
	}
	return result
}

// converts non slice values to []interface{}
func convertToSlice(arr interface{}) []interface{} {
	resSlice := make([]interface{}, 0)
	if arr == nil {
		return resSlice
	}
	resSlice = append(resSlice, arr)
	return resSlice
}

// merges slices and stores it in result
func mergeSlice(result []interface{}, arr, nn interface{}) []interface{} {
	var i, j = 0, 0
	arrVal := reflect.ValueOf(arr)
	nnVal := reflect.ValueOf(nn)

	for i < arrVal.Len() && j < nnVal.Len() {
		result = append(result, arrVal.Index(i).Interface(), nnVal.Index(j).Interface())
		i++
		j++
	}
	for i < arrVal.Len() {
		result = append(result, arrVal.Index(i).Interface())
		i++
	}
	for j < nnVal.Len() {
		result = append(result, nnVal.Index(j).Interface())
		j++
	}
	return result
}
func merge(arr interface{}, nn interface{}) (interface{}, error) {
	if arr == nil && nn == nil {
		return nil, fmt.Errorf("both the arguments are nil, not possible to merge!")
	}
	// to store result
	result := make([]interface{}, 0)
	arrVal := reflect.ValueOf(arr)
	nnVal := reflect.ValueOf(nn)

	if arrVal.Kind() != reflect.Slice {
		arrVal = reflect.ValueOf(convertToSlice(arr))
	}
	if nnVal.Kind() != reflect.Slice {
		nnVal = reflect.ValueOf(convertToSlice(nn))
	}

	//call the mergeSlice()
	result = mergeSlice(result, arrVal.Interface(), nnVal.Interface())
	return result, nil
}

func main() {

	n1 := "7777666609"
	n2 := []string{
		"6666000034",
		"8888000044",
	}
	if res, err := merge(n2, n1); err == nil {
		fmt.Println("Test case 1 ", res)
	}
	if res, err := merge("a", "b"); err == nil {
		fmt.Println("Test case 2", res)
	}
	if res, err := merge(1, 2); err == nil {
		fmt.Println("Test case 3", res)
	}
	arr1 := []int{4, 5, 6}
	arr2 := []int{1, 2, 3}
	if res, err := merge(arr1, arr2); err == nil {
		fmt.Println("Test case 4", res)
	}
}
