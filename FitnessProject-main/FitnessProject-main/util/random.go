package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomWeight() string {
	return fmt.Sprint(RandomInt(0, 300))
}

func RandomHeight() string {
	return fmt.Sprint(RandomInt(0, 80))
}

func RandomAge() string {
	return fmt.Sprint(RandomInt(0, 95))
}

func RandomWeightLifted() string {
	return fmt.Sprint(RandomInt(0, 500))
}

func RandomReps() string {
	return fmt.Sprint(RandomInt(0, 20))
}