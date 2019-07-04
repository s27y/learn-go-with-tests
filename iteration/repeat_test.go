package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	testCases := []struct {
		inChar  string
		inCount int
		want    string
	}{
		{
			"c",
			5,
			"ccccc",
		},
		{
			"a",
			3,
			"aaa",
		},
	}
	for _, tc := range testCases {
		actual := Repeat(tc.inChar, tc.inCount)
		if actual != tc.want {
			t.Errorf("got '%s' but expect '%s'", actual, tc.want)
		}
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
