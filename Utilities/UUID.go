package Utilities

import (
	"math/big"

	"github.com/google/uuid"
)

func GenerateNumericUUID() *big.Int {
	newUuid := uuid.New()

	// Convert a bytes the UUID to a big number (big.Int)
	num := new(big.Int)
	num.SetBytes(newUuid[:])

	return num
}
