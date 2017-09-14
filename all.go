// Example: https://play.golang.org/p/tXEfe9ojPA

func All(input interface{}, pred func(interface{}) (bool, error)) (bool, error) {
	list, err := InterfaceToSlice(input)

	if err != nil {
		return false, err
	}

	for _, v := range list {
		found, err := pred(v)

		if err != nil {
			return false, err
		}

		if !found {
			return false, nil
		}
	}

	return true, nil
}
