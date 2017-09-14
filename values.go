// Values takes a struct and returns all the field values in the struct
func Values(input interface{}) []interface{} {
	reflection := reflect.ValueOf(input)

	result := make([]interface{}, 0)

	for i := 0; i < reflection.NumField(); i++ {
		result = append(result, reflection.Field(i).Interface())
	}

	return result
}
