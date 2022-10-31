package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"user-service/internal/app/domain"
)

type fakeRepo struct{}

func (fr fakeRepo) GetByID(ctx context.Context, id int) (domain.User, error) {
	if id == 1 {
		return domain.User{
			ID:   1,
			Name: "satu",
		}, nil
	}
	return domain.User{}, errors.New("test error")
}

func (fr fakeRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func (fr fakeRepo) Create(ctx context.Context, user *domain.User) error {
	return nil
}

func TestUserService_GetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{context.Background(), 1},
			want: domain.User{
				ID:   1,
				Name: "satu",
			},
			wantErr: false,
		},
		{
			name:    "error when id < 0",
			args:    args{context.Background(), -1},
			want:    domain.User{},
			wantErr: true,
		},
		{
			name:    "error when getting data from repository",
			args:    args{context.Background(), 2},
			want:    domain.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fkRepo := new(fakeRepo)
			service := NewUserService(fkRepo)
			got, err := service.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestUserService_GetByIDWithMock(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		id  int
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    domain.User
// 		wantErr bool
// 		mcFunc  func(*mocks.UserRepository)
// 	}{
// 		{
// 			name: "success",
// 			args: args{context.Background(), 1},
// 			want: domain.User{
// 				ID:   1,
// 				Name: "satu",
// 			},
// 			wantErr: false,
// 			mcFunc: func(ur *mocks.UserRepository) {
// 				ur.On("GetByID", mock.Anything, 1).Return(domain.User{
// 					ID:   1,
// 					Name: "satu",
// 				}, nil)
// 			},
// 		},
// 		{
// 			name:    "error when id < 0",
// 			args:    args{context.Background(), -1},
// 			want:    domain.User{},
// 			wantErr: true,
// 			mcFunc:  func(ur *mocks.UserRepository) {},
// 		},
// 		{
// 			name:    "error when getting data from repository",
// 			args:    args{context.Background(), 2},
// 			want:    domain.User{},
// 			wantErr: true,
// 			mcFunc: func(ur *mocks.UserRepository) {
// 				ur.On("GetByID", mock.Anything, 2).Return(domain.User{}, errors.New("error"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mcRepo := new(mocks.UserRepository)
// 			tt.mcFunc(mcRepo)
// 			service := NewUserService(mcRepo)
// 			got, err := service.GetByID(tt.args.ctx, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("UserService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UserService.GetByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
