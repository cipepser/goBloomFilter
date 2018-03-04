package iblt

import (
	"reflect"
	"testing"
)

func TestNewInvertibleBloomLookupTables(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want InvertibleBloomLookupTables
	}{
		{"Construct Test", args{10}, make([]cell, 10)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInvertibleBloomLookupTables(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInvertibleBloomLookupTables() = %v, want %v", got, tt.want)
			}
		})
	}
}
