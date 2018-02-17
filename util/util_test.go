package util

import "testing"

func TestCalcMD5Hash(t *testing.T) {
	str := "test"

	expect := "098f6bcd4621d373cade4e832627b4f6"
	actual := CalcMD5Hash(str)

	if actual != expect {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestDoubleHashing(t *testing.T) {
	var hashA, hashB int64 = 13, 10
	n := 3
	length := 7

	var expect int64 = 1
	actual := DoubleHashing(hashA, hashB, n, length)

	if actual != expect {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}
