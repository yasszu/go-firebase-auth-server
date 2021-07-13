package registry

import (
	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/application/usecase/mock"

	"github.com/golang/mock/gomock"
)

type Usecase struct {
	Index usecase.IndexUsecase
	User  usecase.UserUsecase
}

func NewUsecase(indexUsecase usecase.IndexUsecase, userUsecase usecase.UserUsecase) *Usecase {
	return &Usecase{
		Index: indexUsecase,
		User:  userUsecase,
	}
}

type MockUsecase struct {
	Index *mock.MockIndexUsecase
	User  *mock.MockUserUsecase
}

func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		Index: mock.NewMockIndexUsecase(ctrl),
		User:  mock.NewMockUserUsecase(ctrl),
	}
}
