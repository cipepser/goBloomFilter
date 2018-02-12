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

}
