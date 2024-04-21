package gettel

import (
	"regexp"
)

func GetTel(text string) []string {
	var tels []string
	reg := regexp.MustCompile(`[(]?\d{3}[^\d]?[\s]?\d{3}[^\d]?\d{4}`)

	nums := reg.FindAllString(text, -1)

	for _, num := range nums {
		tels = append(tels, num)
	}

	return tels
}
