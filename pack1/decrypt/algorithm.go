package decrypt

func Nimbus(str string) string {

	decryptStr := ""

	for _, ch := range str {
		asciiCode := int(ch)
		character := string(rune(asciiCode - 3))
		decryptStr += character

	}
	return decryptStr
}
