package services

import "math/rand"

func GenerateRandom(size int) string {
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