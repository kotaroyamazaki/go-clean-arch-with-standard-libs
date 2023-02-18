package user

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	expectedUser, _ := New(nil, "kotaro", time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC), []int{1, 2, 3})

	type args struct {
		id            *int
		nickName      string
		birthDate     time.Time
		favoriteItems []int
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:            nil,
				nickName:      "kotaro",
				birthDate:     time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				favoriteItems: []int{1, 2, 3},
			},
			want:    expectedUser,
			wantErr: false,
		},
		{
			name: "名前が空の場合はエラーが返る",
			args: args{
				id:        nil,
				nickName:  "",
				birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "生年月日が未来の日付の場合はエラーが返る",
			args: args{
				id:            nil,
				nickName:      "",
				birthDate:     time.Now().Add(1 * time.Microsecond),
				favoriteItems: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.id, tt.args.nickName, tt.args.birthDate, tt.args.favoriteItems)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
