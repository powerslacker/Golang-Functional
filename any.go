package main

import (
	"errors"
	"reflect"
)

func any(list interface{}, pred func(interface{}) (bool, error)) (bool, error) {
	reflection := reflect.ValueOf(list)

	if reflection.Kind() != reflect.Slice {
		return false, errors.New("any: given argument 'list' is not a slice")
	}

	iterable := make([]interface{}, reflection.Len())

	for i, _ := range iterable {
		iterable[i] = reflection.Index(i).Interface()
	}

	for _, v := range iterable {
		found, err := pred(v)

		if err != nil {
			return false, err
		}

		if found {
			return true, nil
		}
	}

	return false, nil
}
