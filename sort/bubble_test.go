package sort

import "testing"

// go test -v -run=TestBubbleSort/test1
func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{nums: []int{1, 3, 11, 23, 112, 2, 66}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.nums)
			BubbleSort(tt.args.nums)
			t.Log(tt.args.nums)
		})
	}
}
