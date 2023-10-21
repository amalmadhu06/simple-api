package usecase

import (
	"context"
	"testing"

	"github.com/amalmadhu06/simple-api/pkg/repository"
)

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name    string
		want    string
		wantErr error
	}{
		{
			name:    "success",
			want:    "User1: Amal, User2: Madhu",
			wantErr: nil,
		},
	}

	mockRepo := repository.NewUserRepository(nil)
	mockUsecase := NewUserUseCase(mockRepo)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := mockUsecase.FindAll(context.TODO())
			if err != tc.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("FindAll() got = %v, want %v", got, tc.want)
			}
		})
	}
}
