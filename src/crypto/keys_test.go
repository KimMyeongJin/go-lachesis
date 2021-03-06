package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Fantom-foundation/go-lachesis/src/common"
)

func TestFakeKeyGeneration(t *testing.T) {
	assert := assert.New(t)

	prev := make([]*common.PrivateKey, 10)
	for i := 0; i < len(prev); i++ {
		prev[i] = GenerateFakeKey(uint64(i))
	}

	for i := 0; i < len(prev); i++ {
		again := GenerateFakeKey(uint64(i))
		if !assert.Equal(prev[i], again) {
			return
		}
	}
}
