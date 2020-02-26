package main

import (
	"reflect"
	"testing"
)

func Test_getAllIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{ip: "102.18.12.2-4"},
			want: []string{"102.18.12.2", "102.18.12.3", "102.18.12.4"},
		},
		{
			name: "test1",
			args: args{ip: "121.21.3.10-13,192.149.12.12"},
			want: []string{"121.21.3.10","121.21.3.11", "121.21.3.12", "121.21.3.13", "192.149.12.12"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllIP(tt.args.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
