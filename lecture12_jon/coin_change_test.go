package lecture12_jon

import "testing"

func Test_change(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "simple test #1",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "simple test #2",
			args: args{n: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := change(tt.args.n); got != tt.want {
				t.Errorf("Want: #{tt.want}, but got: #{got}")
			}
		})
	}
}
