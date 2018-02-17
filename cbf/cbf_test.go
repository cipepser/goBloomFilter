package cbf

import (
	"reflect"
	"testing"
)

func TestNewCountingBloomFilter(t *testing.T) {
	size := 1024
	expect := make([]uint8, size)
	actual := NewCountingBloomFilter(size)

	if reflect.DeepEqual(actual, expect) {
		t.Errorf("\nactual: %v\nexpect: %v\n", actual, expect)
	}

	if len(expect) != len(actual) {
		t.Errorf("\nlength is not same:\nactual: %v\nexpect: %v\n", len(actual), len(expect))
	}
}

func TestAddandExists(t *testing.T) {
	cbf := NewCountingBloomFilter(10)
	cbf.Add("1")

	if cbf.Exists("a") {
		t.Error("expect to get false, but get true")
	}

	if !cbf.Exists("1") {
		t.Error("expect to get true, but get false")
	}
}

func TestRemove(t *testing.T) {
	cbf := NewCountingBloomFilter(10)
	cbf.Add("1")
	cbf.Remove("1")

	if cbf.Exists("1") {
		t.Error("expect to get false, but get true")
	}
}
