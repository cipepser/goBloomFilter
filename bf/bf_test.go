package bf

import (
	"reflect"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	size := 1024
	expect := make([]bool, size)
	actual := NewBloomFilter(size)

	if reflect.DeepEqual(actual, expect) {
		t.Errorf("\nactual: %v\nexpect: %v\n", actual, expect)
	}

	if len(expect) != len(actual) {
		t.Errorf("\nlength is not same:\nactual: %v\nexpect: %v\n", len(actual), len(expect))
	}
}

func TestAddandExists(t *testing.T) {
	bf := NewBloomFilter(10)
	bf.Add("1")

	if bf.Exists("a") {
		t.Error("expect to get false, but get true")
	}

	if !bf.Exists("1") {
		t.Error("expect to get true, but get false")
	}
}
