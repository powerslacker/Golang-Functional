// Example: https://play.golang.org/p/jnwiatxBRK

func Reject(input interface{}, pred func(interface{}) (bool, error)) ([]interface{}, error) {
	list, err := InterfaceToSlice(input)

	result := make([]interface{}, 0)

	if err != nil {
		return result, err
	}

	for _, v := range list {
		found, err := pred(v)

		if err != nil {
			return result, err
		}

		if !found {
			result = append(result, v)
		}
	}

	return result, nil
}
