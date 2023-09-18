package utils

import (
	"math/rand"
	"regexp"
	"strings"
)

const URL_REGEX = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

var urlRegexp, _ = regexp.Compile(URL_REGEX)

func GenerateRandomHash(size int) string {
	randomId := make([]byte, size)
	for i := 0; i < size; i++ {
		randomAscii := byte(rand.Intn(26))

		if num, isNumber := generateNumber(); isNumber {
			randomId[i] = 48 + num
			continue
		}
		randomId[i] = generateCharacter(randomAscii)
	}
	return string(randomId)
}

func IsValidUrl(value string) bool {
	valid := urlRegexp.MatchString(value)
	return valid
}

func AddHttpUrlPrefix(value string) string {
	if strings.HasPrefix(value, "http") {
		return value
	}
	sb := &strings.Builder{}
	sb.WriteString("http://")
	sb.WriteString(value)
	return sb.String()
}

func generateNumber() (byte, bool) {
	if 70 > rand.Intn(100) {
		return 0, false
	}
	return byte(rand.Intn(10)), true
}

func generateCharacter(randomAscii byte) byte {
	isSmall := (rand.Intn(2) == 1)
	if isSmall {
		return 97 + randomAscii
	}
	return 65 + randomAscii
}
