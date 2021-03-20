package formats

var formats [3]string = [3]string{"png", "jpg", "jpeg"}

func IsSupportedFormat(str string) bool {
	for _, v := range formats {
		if v == str {
			return true
		}
	}
	return false
}
