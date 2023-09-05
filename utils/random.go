package utils

import (
	"math/rand"
	"strings"
	"time"
)

var randStr = "qwertyuioplkjhgfdsazxcvbnm"

func init() {
	rand.Seed(time.Now().Unix())

}

func generateRandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)

}

func generateString() string {
	var sb strings.Builder

	var k = len(randStr)

	for i := 0; i < 5; i++ {
		sb.WriteByte(randStr[rand.Intn(k)])

	}
	return sb.String()

}

func GetOwnerName() string {
	return generateString()
}

func GenerateBalance() int64 {
	return generateRandomNumber(0, 10000)
}

func GenerateCurrency() string {
	cs := []string{"EU", "USD", "BDT"}

	c := len(cs)
	return cs[rand.Intn(c)]

}
