package registry

import (
	"github.com/golang/mock/gomock"
	"github.com/yasszu/go-firebase-auth-server/application/usecase"
	"github.com/yasszu/go-firebase-auth-server/application/usecase/mock"
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
