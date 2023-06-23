package do

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		fn   func() any
		want any
	}{
		{
			"works with strings",
			func() any { return "result" },
			"result",
		},
		{
			"works with integers",
			func() any { return 1 },
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
