package handlers

import (
	"errors"
	"testing"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"
)

func TestGetStatusCode(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "validation error returns 400",
			args: args{
				err: entities.ErrValidattion,
			},
			want: 400,
		},
		{
			name: "not found error returns 500",
			args: args{
				err: entities.ErrNotFound,
			},
			want: 500,
		},
		{
			name: "duplicated primary key error returns 500",
			args: args{
				err: entities.ErrDuplicatedPrimaryKey,
			},
			want: 500,
		},
		{
			name: "unknown error returns 500",
			args: args{
				err: errors.New("unknown error"),
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStatusCode(tt.args.err); got != tt.want {
				t.Errorf("GetStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
