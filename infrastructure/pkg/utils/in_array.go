package utils

func InArray[T string | int | int32](array []T, data T) bool {
	for _, item := range array {
		if data == item {
			return true
		}
	}
	return false
}
