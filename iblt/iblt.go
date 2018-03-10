package iblt

import (
	"strconv"

	"github.com/cipepser/goBloomFilter/util"
)

type cell struct {
	count, keySum, valueSum int
}

// IBLT is a slice of cell.
type IBLT []cell

var (
	k = 3
)

// NewIBLT constructs InvertibleBloomLookupTables with table size.
func NewIBLT(size int) IBLT {
	return make([]cell, size)
}

// Insert inserts a key-value pair into InvertibleBloomLookupTables.
// This operation always succeeds.
func (iblt IBLT) Insert(key, value int) {
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

// Get checks wether the key exists in IBLT or not.
func (iblt IBLT) Get(key int) (bool, int) {
	hash := util.CalcMD5Hash(strconv.Itoa(key))
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		idx := util.DoubleHashing(i64HashA, i64HashB, i, len(iblt))
		if iblt[idx].count == 0 {
			return false, 0
		}
		if iblt[idx].count == 1 {
			if iblt[idx].keySum == key {
				return true, iblt[idx].valueSum
			}
			return false, 0
		}
	}
	return false, 0
}
