// Example: https://play.golang.org/p/9Gfa8N1Tfe

func Map(input interface{}, pred func(interface{}) (interface{}, error)) ([]interface{}, error) {
	list, err := InterfaceToSlice(input)

	result := make([]interface{}, 0)

	if err != nil {
		return result, err
	}

	for _, v := range list {
		transform, err := pred(v)

		if err != nil {
			return result, err
		}

		result = append(result, transform)
	}

	return result, nil
}
