package utils

import (
	"math/rand"
	"time"
)

func GenerateSixDigitCode() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := rng.Intn(900000) + 100000 // Generate a random number between 100000 and 999999

	return otp
}
