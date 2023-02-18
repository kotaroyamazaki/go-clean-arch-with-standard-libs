package user

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		nickName  string
		birthDate time.Time
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
				nickName:  "kotaro",
				birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
			},
			want: &User{
				id:       0,
				NickName: "kotaro",
				BirthDate: BitrhDate{
					birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				},
			},
			wantErr: false,
		},
		{
			name: "名前が空の場合はエラーが返る",
			args: args{
				nickName:  "",
				birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "生年月日が未来の日付の場合はエラーが返る",
			args: args{
				nickName:  "",
				birthDate: time.Now().Add(1 * time.Microsecond),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.nickName, tt.args.birthDate)
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
