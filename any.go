func Any(input interface{}, pred func(interface{}) (bool, error)) (bool, error) {
	list, err := InterfaceToSlice(input)
	
	if err != nil {
		return false, err
	}
	
	for _, v := range list {
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
