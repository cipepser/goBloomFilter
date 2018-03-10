package iblt

import (
	"reflect"
	"testing"
)

func TestNewIBLT(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want IBLT
	}{
		{"Construct Test", args{10}, make([]cell, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIBLT(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvertibleBloomLookupTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIBLT_InsertndGet(t *testing.T) {
	type args struct {
		key    int
		value  int
		getkey int
	}
	tests := []struct {
		name string
		iblt IBLT
		args args
		want bool
	}{
		{
			"Insert {1, 3} and Get(1)",
			NewIBLT(10),
			args{1, 3, 1},
			true,
		},
		{
			"Insert {1, 3} and Get(2)",
			NewIBLT(10),
			args{1, 3, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.iblt.Insert(tt.args.key, tt.args.value)
			got, _ := tt.iblt.Get(tt.args.getkey)
			if got != tt.want {
				t.Errorf("Get(1) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIBLT_Delete(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Delete {2, 5}",
			args{2, 5},
			false,
		},
	}
	iblt := NewIBLT(10)
	// prepare for tests: Insert {2, 5} and {3, 8}
	iblt.Insert(2, 5)
	iblt.Insert(3, 8)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iblt.Delete(tt.args.key, tt.args.value)
			got, _ := iblt.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get(1) = %v, want %v", got, tt.want)
			}
		})
	}
}
