package iteration

const repeatCount = 5

func Repeat(character string) string {
	repeated := ""
	for i := 0; i < repeatCount; i++ {
		// += called "the Add AND assignment operator", adds the right operand to the left operand and assigns the result to left operand. It works with other types like integers.
		repeated += character
	}
	return repeated
}
