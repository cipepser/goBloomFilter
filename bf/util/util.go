package util

import (
	"crypto/md5"
	"encoding/hex"
)

// CalcMD5Hash calcurates hash value of the string input by MD5.
func CalcMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

// DoubleHashing calcurates hash value by double-hashing.
//
// func DoubleHashing(hashA, hashB int64, n int) (hash int64) {
// 	// h = hashA + n * hashBの計算
// 	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
// 	h = new(big.Int).Add(big.NewInt(hashA), h)
// 	h = new(big.Int).Rem(h, big.NewInt(int64(size)))
//
// 	// 余りが負の数になったときは正の余りにする
// 	hash = h.Int64()
// 	if hash < 0 {
// 		hash += int64(size)
// 	}
// 	return
// }
