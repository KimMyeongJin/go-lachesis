package crypto

import (
	"github.com/Fantom-foundation/go-lachesis/src/common"
)

// AddressOf calcs hash of the PublicKey.
func AddressOf(pk common.PublicKey) common.Address {
	return common.Address(Keccak256Hash(pk.Bytes()))
}
