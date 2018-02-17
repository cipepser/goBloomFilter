package cbf

import (
	"strconv"

	"github.com/cipepser/goBloomFilter/util"
)

var (
	// k is the number of hash functions.
	k = 3
	// k int = int(math.Log(2) * float64(size) / float64(n))
)

// CountingBloomFilter is a slice of uint8.
// We can count up until 255.
type CountingBloomFilter []uint8

// NewBloomFilter constructs BloomFilter with table sise.
func NewCountingBloomFilter(size int) CountingBloomFilter {
	return make([]uint8, size)
}

// Add adds a element to CountingBloomFilter.
// Calcurates the hash value of element, and divede it to two,
// and calcurates k hash values by Double-Hashing method.
// We seem the k hash values as indeces,
// inclement the count of  CountingBloomFilter.
func (cbf CountingBloomFilter) Add(element string) {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		cbf[util.DoubleHashing(i64HashA, i64HashB, i, len(cbf))]++
	}
}

// Exists chech wheter CountingBloomFilter have the element or not.
// Calcurates the hash value of element, and divede it to two,
// and calcurates k hash values by Double-Hashing method.
// If any element of CountingBloomFilter indicated by the k indeces is 0,
// CountingBloomFilter don't have the element.
// If all elements are positive number, BloomFilter may have the element
// with low probability of false positive.
func (cbf CountingBloomFilter) Exists(element string) bool {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		if cbf[util.DoubleHashing(i64HashA, i64HashB, i, len(cbf))] == 0 {
			return false
		}
	}
	return true
}

// Remove removes a element from CountingBloomFilter.
// Calcurates the hash value of element, and divede it to two,
// and calcurates k hash values by Double-Hashing method.
// We seem the k hash values as indeces,
// declement the count of  CountingBloomFilter.
func (cbf CountingBloomFilter) Remove(element string) {
	hash := util.CalcMD5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]

	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		cbf[util.DoubleHashing(i64HashA, i64HashB, i, len(cbf))]--
	}
}
