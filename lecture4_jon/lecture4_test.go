package lecture4_jon

import "testing"

func TestFoo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test simple2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test simple2",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climb(tt.args.n); got != tt.want {
				t.Errorf("Getting incorrect result, want: #{tt.want} but got #{got}")
			}
		})
	}
}
