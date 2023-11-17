package models

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	expectedUser, err := NewUser("kotaro", time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC), []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
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
				nickName:  "",
				birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "生年月日が未来の日付の場合はエラーが返る",
			args: args{
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
			got, err := NewUser(tt.args.nickName, tt.args.birthDate, tt.args.favoriteItems)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assertUser(t, got, tt.want, "NewUser()")
		})
	}
}

func assertUser(t *testing.T, got, want *User, msg string) {
	if got == nil && want == nil {
		return
	}
	if got == nil || want == nil {
		t.Fatalf("%s: got = %v, want = %v", msg, got, want)
	}

	if got.nickName != want.nickName {
		t.Errorf("%s: nickName got = %v, want = %v", msg, got.nickName, want.nickName)
	}
	if got.birthDate != want.birthDate {
		t.Errorf("%s: birthDate got = %v, want = %v", msg, got.birthDate, want.birthDate)
	}
	if !reflect.DeepEqual(got.favoriteItems, want.favoriteItems) {
		t.Errorf("%s: favoriteItems got = %v, want = %v", msg, got.favoriteItems, want.favoriteItems)
	}
}
