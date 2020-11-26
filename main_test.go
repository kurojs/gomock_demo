package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kurojs/gomock_demo/mock"
)

func Test_calculate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		getMock func() Stacker
		arr     []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with Mock",
			args: args{
				getMock: func() Stacker {
					mock := mock.NewMockStacker(ctrl)
					// Assert mock are called 3 times
					for i := 1; i <= 3; i++ {
						mock.EXPECT().Append(gomock.Eq(i))
					}
					// And return 3
					mock.EXPECT().Pop().Return(3)

					return mock
				},
				arr: []interface{}{1, 2, 3},
			},
			want: 3,
		},
		{
			name: "test with Stub",
			args: args{
				getMock: func() Stacker {
					mock := mock.NewMockStacker(ctrl)
					// No assertion, just return the expected data
					mock.EXPECT().Append(gomock.Any()).AnyTimes()
					mock.EXPECT().Pop().DoAndReturn(func() interface{} {
						return 3
					})

					return mock
				},
				arr: []interface{}{1, 2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.getMock(), tt.args.arr...); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
