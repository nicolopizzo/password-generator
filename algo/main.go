package algo

import (
	"math/rand"
	"time"
)

/*
parameters -> { digits, special, maiusc }
*/

func sum(m map[string]int) int {
	s := 0
	for _, v := range m {
		s += v
	}

	return s
}

func getLengthMap(parameters map[string]bool, length int) map[string]int {
	lengthByKind := make(map[string]int)
	lengthByKind["digits"] = 0
	lengthByKind["special"] = 0
	lengthByKind["maiusc"] = 0

	if parameters["digits"] {
		lengthByKind["digits"] = length / 8
	}
	if parameters["special"] {
		lengthByKind["special"] = length / 5
	}
	if parameters["maiusc"] {
		lengthByKind["maiusc"] = length / 5
	}

	s := sum(lengthByKind)
	lengthByKind["chars"] = length - s
	return lengthByKind
}

var charset = map[string]string{
	"digits":  "0123456789",
	"maiusc":  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"special": "_@!?/-",
	"chars":   "abcdefghijklmnopqrstuvwxyz",
}

func randomRange(a, b uint8) []uint8 {
	r := make([]uint8, b-a)
	for i := range r {
		r[i] = a + uint8(i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })

	return r
}

func NewPassword(parameters map[string]bool, length int) string {
	lengths := getLengthMap(parameters, length)

	var password []uint8
	for k, v := range lengths {
		for i := range randomRange(0, uint8(v)) {
			password = append(password, charset[k][i])
		}
	}

	rand.Shuffle(length, func(i, j int) { password[i], password[j] = password[j], password[i] })
	return string(password)
}
