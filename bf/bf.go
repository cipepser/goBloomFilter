package bf

import (
	"strconv"

	"../util"
)

var (
	// k is the number of hash functions.
	k = 3
	// k int = int(math.Log(2) * float64(size) / float64(n))
)

// BloomFilter is a slice of boolen.
type BloomFilter []bool

// NewBloomFilter constructs BloomFilter with table sise.
func NewBloomFilter(size int) BloomFilter {
	return make([]bool, size)
}

// Add adds the elements to BloomFilter.
// Calcurates the hash value of element, and divede it to two,
// and calcurates k hash values by Double-Hashing method.
// We seem the k hash values as indeces,
// make BloomFilter to be true.
func (bf BloomFilter) Add(element string) {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		bf[util.DoubleHashing(i64HashA, i64HashB, i, len(bf))] = true
	}
}

// Exists chech wheter BloomFilter have the element or not.
// Calcurates the hash value of element, and divede it to two,
// and calcurates k hash values by Double-Hashing method.
// If any element of BloomFilter indicated by the k indeces is false,
// BloomFilter don't have the element.
// If all elements are true, BloomFilter may have the element
// with low probability of false positive.
func (bf BloomFilter) Exists(element string) bool {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		if !bf[util.DoubleHashing(i64HashA, i64HashB, i, len(bf))] {
			return false
		}
	}
	return true
}
