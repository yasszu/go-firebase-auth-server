package registry

import (
	"github.com/golang/mock/gomock"

	"go-firebase-auth-server/application/usecase/mock"

	"go-firebase-auth-server/application/usecase"
)

func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		IndexUsecase: mock.NewMockIndexUsecase(ctrl),
		UserUsecase:  mock.NewMockUserUsecase(ctrl),
	}
}

type MockUsecase struct {
	IndexUsecase usecase.IndexUsecase
	UserUsecase  usecase.UserUsecase
}

func (m MockUsecase) NewIndex() usecase.IndexUsecase {
	return m.IndexUsecase
}

func (m MockUsecase) NewUser() usecase.UserUsecase {
	return m.UserUsecase
}
