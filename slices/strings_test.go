package slices

import "testing"

func TestStringInSlice(t *testing.T) {
	const happyPathValue = "hello"
	const badPathValue = "BONK"
	inputFixture := []string{
		happyPathValue,
		"bonkers",
		"world",
	}
	type args struct {
		text  string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true if the value exist in the slice",
			args: args{
				text:  happyPathValue,
				slice: inputFixture,
			},
			want: true,
		},
		{
			name: "should return false if the value exist in the slice",
			args: args{
				text:  badPathValue,
				slice: inputFixture,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.text, tt.args.slice); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
