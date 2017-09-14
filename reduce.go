// https://play.golang.org/p/-JWp7ERi3H

func Reduce(input interface{}, pred func(interface{}, interface{}) (interface{}, error)) (interface{}, error) {
	list, err := InterfaceToSlice(input)

	result := input

	if err != nil {
		return result, err
	}

	accumulator := list[0]

	for i := 1; i < len(list); i++ {
		accumulator, err = pred(accumulator, list[i])

		if err != nil {
			return result, err
		}

		result = accumulator
	}

	return result, nil
}
