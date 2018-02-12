package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
)

const (
	size int = 1024 // BloomFilterの大きさ[bit]
)

var (
	// k int = 10 // ハッシュ関数の数

	n = 70 // 挿入する要素の数
	k int = int(math.Log(2) * float64(size) / float64(n)) // ハッシュ関数の数
)

type BloomFilter struct {
	BloomFilter [size]bool
}

// MD5でハッシュ値を求める
func GetMD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	
	return hex.EncodeToString(hasher.Sum(nil))
}

// DoubleHashing法にてハッシュ値を求める
func DoubleHashing(hashA, hashB int64, n int) (hash int64) {	
	// h = hashA + n * hashBの計算
	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
	h = new(big.Int).Add(big.NewInt(hashA), h)
	h = new(big.Int).Rem(h, big.NewInt(int64(size)))
	
	// 余りが負の数になったときは正の余りにする
	hash = h.Int64()
	if hash < 0 {
		hash += int64(size)
	}
	return
}

// BloomFilterに要素を追加する
// ハッシュ値を2つに分けて、DoubleHashing法でk個のハッシュ値を求める。
// k個のハッシュ値のindexをtrueにする
func Add(bf *BloomFilter, element string)  {
	hash := GetMD5Hash(element)
	hashA := hash[:int(len(hash)/2)] // 前半
	hashB := hash[int(len(hash)/2):] // 後半
	
	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)
	
	// TODO: goroutineにする
	for i := 0; i < k; i++ {
		bf.BloomFilter[DoubleHashing(i64_hashA, i64_hashB, i)] = true
	}
}

// BloomFilterに要素が含まれているか
// ハッシュ値を2つに分けて、DoubleHashing法でk個のハッシュ値を求める。
// k個のハッシュ値のindexで値がひとつでもfalseであれば要素は含まれていない
// 全部trueのときは偽陽性の可能性あり
func Exists(bf *BloomFilter, element string) (exists bool) {
	hash := GetMD5Hash(element)
	hashA := hash[:int(len(hash)/2)] // 前半
	hashB := hash[int(len(hash)/2):] // 後半
	
	i64_hashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64_hashB, _ := strconv.ParseInt(hashB, 16, 64)
	
	// TODO: goroutineにする
	exists = true
	for i := 0; i < k; i++ {
		if bf.BloomFilter[DoubleHashing(i64_hashA, i64_hashB, i)] == false {
			exists = false
			break
		}
	}
	
	return
}

func main() {
	var bf BloomFilter

	// 要素の追加
	Add(&bf, "1")
	Add(&bf, "2")
	Add(&bf, "3")
	Add(&bf, "4")
	Add(&bf, "5")
	
	// 要素が含まれているか検証
	fmt.Println(Exists(&bf, "2"))
}
