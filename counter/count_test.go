package counter_test

import (
	"reflect"
	"testing"

	"github.com/samdelacruz/wordcount/counter"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want map[string]int
	}{
		{"One word", "hello", map[string]int{"hello": 1}},
		{"Multi word", "hello hello", map[string]int{"hello": 2}},
		{"Case insensitive", "Hello hello", map[string]int{"hello": 2}},
		{"Many words", "hello world hello", map[string]int{"hello": 2, "world": 1}},
		{"Empty", "", map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := counter.FromString(&tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
