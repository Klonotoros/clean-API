package service

import (
	"clean-API/internal/dto"
	"clean-API/internal/model"
	"clean-API/internal/repository"
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("UserService", func() {

	var (
		userService UserService

		userRepoCtrl *gomock.Controller
		userRepoMock *repository.MockUserRepository
	)

	BeforeEach(func() {
		userRepoCtrl = gomock.NewController(GinkgoT())
		userRepoMock = repository.NewMockUserRepository(userRepoCtrl)

		userService = newUserService(userRepoMock, dto.Config{})
	})

	Describe("Login", func() {

		It("should return err if user repo failed to fetch user", func() {
			userRepoMock.EXPECT().GetUserByEmail("test@gmail.com").Return(model.User{}, errors.New("some-error"))

			token, err := userService.Login("test@gmail.com", "sample_password")
			Expect(err.Error()).To(Equal("some-error"))
			Expect(token).To(BeEmpty())
		})

		It("should return err if hash password is invalid", func() {
			userRepoMock.EXPECT().GetUserByEmail("test@gmail.com").Return(model.User{Password: "1234"}, nil)

			token, err := userService.Login("test@gmail.com", "sample_password")
			Expect(err.Error()).To(Equal("invalid password"))
			Expect(token).To(BeEmpty())
		})

	})

})
