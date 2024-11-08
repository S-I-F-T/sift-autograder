package validation

type Validator func(expected, actual string, max_points int) (int, error)

func AllOrNothing(expected, actual string, max_points int) (points int, err error) {
	if expected == actual {
		points = max_points
	} else {
		points = 0
	}
	return
}
