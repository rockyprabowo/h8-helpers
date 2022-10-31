package options

import (
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	type args[T string] struct {
		value        T
		defaultValue T
	}
	tests := []struct {
		name string
		args args[string]
		want string
	}{
		{
			name: "should return value",
			args: args[string]{
				value:        "TestData",
				defaultValue: "TestDefaultData",
			},
			want: "TestData",
		},
		{
			name: "should return default value if value is zero value",
			args: args[string]{
				value:        "",
				defaultValue: "TestDefaultData",
			},
			want: "TestDefaultData",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(tt.args.value, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
