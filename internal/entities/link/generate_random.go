package link

import (
	"math/rand"
	"strings"
	"time"
)

var _alphabet = []rune(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-`)

func GenerateRandomPath(length int64) string {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	var b strings.Builder
	for i := int64(0); i < length; i++ {
		b.WriteRune(_alphabet[randomizer.Intn(len(_alphabet))])
	}

	return b.String()
}
