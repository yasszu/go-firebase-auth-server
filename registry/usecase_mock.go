package registry

import (
	"github.com/golang/mock/gomock"
	"go-firebase-auth-server/application/usecase"
	"go-firebase-auth-server/application/usecase/mock"
)

func NewMockUsecase(ctrl *gomock.Controller) Usecase {
	return &mockUsecase{ctrl: ctrl}
}

type mockUsecase struct {
	ctrl *gomock.Controller
}

func (m mockUsecase) NewIndex() usecase.IndexUsecase {
	return mock.NewMockIndexUsecase(m.ctrl)
}

func (m mockUsecase) NewUser() usecase.UserUsecase {
	return mock.NewMockUserUsecase(m.ctrl)
}
