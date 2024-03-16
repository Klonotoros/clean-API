package service

import (
	"clean-API/internal/dto"
	"clean-API/internal/model"
	"clean-API/internal/repository"
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("UserService", func() {

	var (
		userService  UserService
		userRepoCtrl *gomock.Controller
		userRepoMock *repository.MockUserRepository
	)

	BeforeEach(func() {
		userRepoCtrl = gomock.NewController(GinkgoT())
		userRepoMock = repository.NewMockUserRepository(userRepoCtrl)
		userService = newUserService(userRepoMock, dto.Config{})

	})

	AfterEach(func() {
		userRepoCtrl.Finish()
	})

	Describe("Login", func() {
		Context("when user provides correct credentials", func() {
			It("should return JWT token, and no error", func() {
				// given
				userRepoMock.EXPECT().GetUserByEmail("test@email.com").Return(model.User{Password: "$2a$14$5lZMVIy0ESQ2cgEcOgLN1u59r48PZlV2Mq2YJK1Bs98kxNzGqNrum"}, nil)
				// when
				token, err := userService.Login("test@email.com", "password")
				// then
				Expect(err).To(BeNil())
				Expect(token).NotTo(BeEmpty())
				fmt.Println(token)
			})
		})

		Context("when user provides incorrect email", func() {
			It("should not return jwt token, and return error", func() {
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any()).Return(model.User{}, errors.New("user not found"))

				token, err := userService.Login("incorrect@email.com", "password")
				Expect(err.Error()).To(Equal("user not found"))
				Expect(token).To(BeEmpty())
			})
		})

		Context("when user provides incorrect password", func() {
			It("should not return jwt token and return error", func() {
				userRepoMock.EXPECT().GetUserByEmail("correct@email.com").Return(model.User{Password: "password"}, nil)

				token, err := userService.Login("correct@email.com", "incorrectPassword")
				Expect(err.Error()).To(Equal("invalid password"))
				Expect(token).To(BeEmpty())
			})
		})

	})

	Describe("Register", func() {
		Context("when email is empty", func() {
			It("should return error", func() {
				token, err := userService.Register("", "password!123")
				Expect(err.Error()).To(Equal("email is required"))
				Expect(token).To(BeEmpty())
			})
		})

		Context("when password is empty", func() {
			It("should return error", func() {
				token, err := userService.Register("test@email.com", "")
				Expect(err.Error()).To(Equal("password is required"))
				Expect(token).To(BeEmpty())
			})
		})

		Context("when password is too short", func() {
			It("should return error", func() {
				token, err := userService.Register("test@email.com", "test")
				Expect(err.Error()).To(Equal("password must be at least 8 characters"))
				Expect(token).To(BeEmpty())
			})
		})

		Context("when user with provided email already exists", func() {
			It("should return error", func() {

				existingUser := model.User{
					ID:    10,
					Email: "existing@example.com",
				}

				userRepoMock.EXPECT().GetUserByEmail("existing@example.com").Return(existingUser, nil)
				token, err := userService.Register("existing@example.com", "password")

				Expect(err.Error()).To(Equal("user with email existing@example.com already exists"))
				Expect(token).To(BeEmpty())
			})
		})

		Context("when provided data is correct", func() {
			It("should return JWT token and no error", func() {

				userRepoMock.EXPECT().GetUserByEmail("test@test.com").Return(model.User{}, errors.New("user not found"))
				userRepoMock.EXPECT().Save(gomock.Any()).DoAndReturn(func(user model.User) {
					Expect(user.Email).To(Equal("test@test.com"))
				}).Return(model.User{
					Email:    "test@test.com",
					Password: "password123!",
					ID:       10,
				}, nil)

				token, err := userService.Register("test@test.com", "password123!")

				Expect(err).To(BeNil())
				Expect(token).NotTo(BeEmpty())
			})
		})
	})
})
