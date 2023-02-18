package usecases

import (
	"context"
	"testing"
	"time"

	entities "github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities/repository"
	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/infra/repository"
)

func Test_userUsecase_Register(t *testing.T) {
	userRepo := repository.NewUserRepository()
	favoriteBookRepo := repository.NewUserFavoriteBooksRepository()
	ctx := context.Background()

	type fields struct {
		userRepo         entities.UserRepository
		favoriteBookRepo entities.UserFavoriteBooksRepository
	}
	type args struct {
		ctx             context.Context
		birthDate       time.Time
		nickname        string
		favoriteBookIDs []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				userRepo:         userRepo,
				favoriteBookRepo: favoriteBookRepo,
			},
			args: args{
				ctx:             ctx,
				birthDate:       time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				nickname:        "kotap",
				favoriteBookIDs: []int{1, 2, 3},
			},
			wantErr: false,
		},
		{
			name: "興味のある本が空の場合でも成功する",
			fields: fields{
				userRepo:         userRepo,
				favoriteBookRepo: favoriteBookRepo,
			},
			args: args{
				ctx:       ctx,
				birthDate: time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				nickname:  "kotap",
			},
			wantErr: false,
		},
		{
			name: "ニックネームが空の場合は失敗する",
			fields: fields{
				userRepo:         userRepo,
				favoriteBookRepo: favoriteBookRepo,
			},
			args: args{
				ctx:             ctx,
				birthDate:       time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				favoriteBookIDs: []int{1, 2, 3},
			},
			wantErr: true,
		},
		{
			name: "生年月日が空の場合は失敗する",
			fields: fields{
				userRepo:         userRepo,
				favoriteBookRepo: favoriteBookRepo,
			},
			args: args{
				ctx:             ctx,
				birthDate:       time.Date(1997, 3, 18, 15, 0, 0, 0, time.UTC),
				favoriteBookIDs: []int{1, 2, 3},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userUsecase{
				userRepo:         tt.fields.userRepo,
				favoriteBookRepo: tt.fields.favoriteBookRepo,
			}
			if err := uc.Register(tt.args.ctx, tt.args.birthDate, tt.args.nickname, tt.args.favoriteBookIDs); (err != nil) != tt.wantErr {
				t.Errorf("userUsecase.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
