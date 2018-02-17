package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
)

// CalcMD5Hash calcurates hash value of the string input by MD5.
func CalcMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

// DoubleHashing calcurates hash value by double-hashing.
// h = hashA + n * hashB % length
func DoubleHashing(hashA, hashB int64, n, length int) (hash int64) {
	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
	h = new(big.Int).Add(big.NewInt(hashA), h)
	h = new(big.Int).Rem(h, big.NewInt(int64(length)))

	// if the rem is negative, make it positive.
	hash = h.Int64()
	if hash < 0 {
		hash += int64(length)
	}
	return
}
