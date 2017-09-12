func InterfaceToSlice(input interface{}) ([]interface{}, error) {
	reflection := reflect.ValueOf(input)

	if reflection.Kind() != reflect.Slice {
		return make([]interface{}, 0), errors.New("any: given argument 'list' is not a slice")
	}

	list := make([]interface{}, reflection.Len())

	for i, _ := range list {
		list[i] = reflection.Index(i).Interface()
	}

	return list, nil
}
