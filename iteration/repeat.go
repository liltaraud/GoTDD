package iteration

// Repeat takes a charactecter and returns
// a string with the char repeated 5 times
func Repeat(char string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}
