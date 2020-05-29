package tools

import "math/rand"

func GetName() string {
	name := ""
	name_length := rand.Intn(3) + 2
	for i := 0; i < name_length; i++ {
		name += GetRandStr(rand.Intn(7) + 3)
		if i < (name_length - 1) {
			name += " "
		}
	}
	return name
}

func GetRandStr(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		rune := rune(97 + rand.Intn(26))
		result += string(rune)
	}
	result = string(result[0]-32) + result[1:]
	return result
}
