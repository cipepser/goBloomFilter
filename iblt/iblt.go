package iblt

import (
	"strconv"

	"github.com/cipepser/goBloomFilter/util"
)

type cell struct {
	count, keySum, valueSum int
}

// InvertibleBloomLookupTables is a slice of cell.
type InvertibleBloomLookupTables []cell

var (
	k = 3
)

// NewInvertibleBloomLookupTables constructs InvertibleBloomLookupTables with table size.
func NewInvertibleBloomLookupTables(size int) InvertibleBloomLookupTables {
	return make([]cell, size)
}

// Insert inserts a key-value pair into InvertibleBloomLookupTables.
// This operation always succeeds.
func (iblt InvertibleBloomLookupTables) Insert(key, value int) {
	hash := util.CalcMD5Hash(strconv.Itoa(key))
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		iblt[util.DoubleHashing(i64HashA, i64HashB, i, len(iblt))].count++
		iblt[util.DoubleHashing(i64HashA, i64HashB, i, len(iblt))].keySum ^= key
		iblt[util.DoubleHashing(i64HashA, i64HashB, i, len(iblt))].valueSum ^= value
	}
}

// func (iblt InvertibleBloomLookupTables) LookUp() {
//
// }
//
// func (iblt InvertibleBloomLookupTables) Delete() {
//
// }
//
// func (iblt InvertibleBloomLookupTables) ListUp() {
//
// }
