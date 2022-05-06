package utils

import (
	"math/rand"
	"strings"
	"time"
)

// TODO: Añadir compatibilidad con caracteres latinos
const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	a := sb.String()

	return a
}

func RandomUser() string {
	return RandomString(6)
}

func RandomUnits() string {
	units := []string{"days", "months", "years"}
	n := len(units)
	return units[rand.Intn(n)]
}

func RandomPeriod() int64 {
	return RandomInt(1, 12)
}

func RandomPrice() int64 {
	return RandomInt(500000, 25000000)
}
