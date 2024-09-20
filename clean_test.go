package main

import "testing"

func Test_clean(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "index out of bounds -1",
			args: args{
				input: `\`,
			},
			want: ``,
		},
		{
			name: "index out of bounds -2",
			args: args{
				input: `\"`,
			},
			want: `"`,
		},
		{
			name: "clean",
			args: args{
				input: `{\n  \"field1\" : \"value1\",\n  \"field2\" : 123.5\n}`,
			},
			want: `{"field1":"value1","field2":123.5}`,
		},
		{
			name: "escaped string inside quotes",
			args: args{
				input: `{\n  \"field1\" : \"value1 \\"with escaped quotes\\"\",\n  \"field2\" : 123.5\n}`,
			},
			want: `{"field1":"value1 \"with escaped quotes\"","field2":123.5}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clean(tt.args.input); got != tt.want {
				t.Errorf("clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
