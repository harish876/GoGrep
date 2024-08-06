package lib

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func IsAlpha(char byte) bool {
	return (char >= 'a' && char < 'z') || (char >= 'A' && char <= 'Z')
}

func reverse[T any](slice []T) {
	start := 0
	end := len(slice) - 1
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
}
