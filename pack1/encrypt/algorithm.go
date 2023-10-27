package encrypt

func Nimbus(str string) string {

	encryptStr := ""

	for _, ch := range str {
		asciCode := int(ch)
		character := string(rune(asciCode + 3))
		encryptStr += character
	}

	return encryptStr
}
