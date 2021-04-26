package secret

import (
	"math/bits"
)



func Handshake(input uint) []string {
	var secret []string

	if input&(1<<0) > 0 {
		secret = append(secret, "wink")
	}
	if input&(1<<1) > 0 {
		secret = append(secret, "double blink")
	}
	if input&(1<<2) > 0 {
		secret = append(secret, "close your eyes")
	}
	if input&(1<<3) > 0 {
		secret = append(secret, "jump")
	}
	if bits.Len(input) >= 5 {
		var secretReverse []string
		for len(secret) > 0 {
			secretReverse = append(secretReverse, secret[len(secret) - 1])
			secret = secret[:len(secret) -1 ]

		}
		secret = secretReverse
	}
	return secret

}
