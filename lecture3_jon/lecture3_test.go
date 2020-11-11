package lecture3_jon

import "testing"

func Test_Sum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test simple",
			args: args{
				n: 3,
			},
			want: 6,
		},
		{
			name: "test 0",
			args: args{
				n: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nSum(tt.args.n); got != tt.want {
				t.Errorf("nSum() =  #{got}, want = #{tt.want} ")
			}
			if got := nSumDp(tt.args.n); got != tt.want {
				t.Errorf("nSum() =  #{got}, want = #{tt.want} ")
			}
		})
	}
}
