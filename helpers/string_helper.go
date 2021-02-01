package helpers

import "strings"

//RemoveURLFromText remove URL from text
func RemoveURLFromText(text string) string {
	textList := strings.Split(text, " ")

	for i, w := range textList {
		if strings.Contains(w, "https") {
			textList[i] = ""
		}
	}

	return strings.Join(textList[:len(textList)-1], " ")
}
