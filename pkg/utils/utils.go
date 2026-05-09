package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func GenerateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("otp length must be positive")
	}

	var builder strings.Builder
	builder.Grow(length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		builder.WriteByte(byte('0' + n.Int64()))
	}

	return builder.String(), nil
}
