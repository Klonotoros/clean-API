package controller

import (
	"bytes"
	"clean-API/internal/dto"
	"clean-API/internal/service"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("UserController", func() {
	var (
		userController  UserController
		userServiceCtrl *gomock.Controller
		userServiceMock *service.MockUserService
		w               *httptest.ResponseRecorder
		ctx             *gin.Context
	)

	BeforeEach(func() {
		userServiceCtrl = gomock.NewController(GinkgoT())
		userServiceMock = service.NewMockUserService(userServiceCtrl)
		w = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
		userController = newUserController(userServiceMock)
	})

	AfterEach(func() {
		userServiceCtrl.Finish()
	})

	Describe("Login", func() {
		Context("on successful login", func() {
			It("should return 200 OK and a token", func() {
				// given
				expectedToken := "sample-token"

				expectedResponse := dto.LoginResponse{
					Message: "Login successful!",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{"email":"test@example.com","password":"password"}`)

				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				userServiceMock.EXPECT().Login("test@example.com", "password").Return(expectedToken, nil)

				// when
				userController.Login(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))
			})
		})
		Context("on wrong request", func() {
			It("should return 400 StatusBadRequest and error message", func() {
				// given
				expectedToken := ""

				expectedResponse := dto.LoginResponse{
					Message: "Login fail",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{invalid}`)

				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				// when
				userController.Login(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))

			})
		})

		Context("when credentials are invalid", func() {
			It("should return 401 StatusUnauthorized and error message", func() {
				// given
				expectedToken := ""

				expectedResponse := dto.LoginResponse{
					Message: "Login fail",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{"email":"test@example.com","password":"password123"}`)

				req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				userServiceMock.EXPECT().Login("test@example.com", "password123").Return(expectedToken, errors.New("login error"))
				// when
				userController.Login(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))

			})
		})
	})

	Describe("Register", func() {
		Context("on successful register", func() {
			It("should return 201 status created and a token", func() {
				// given
				expectedToken := "sample-token"

				expectedResponse := dto.RegisterResponse{
					Message: "Register successful!",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{"email":"test@example.com","password":"password"}`)

				req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				userServiceMock.EXPECT().Register("test@example.com", "password").Return(expectedToken, nil)

				// when
				userController.Register(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))
			})
		})

		Context("on wrong request", func() {
			It("should return 400 StatusBadRequest and error message", func() {
				// given
				expectedToken := ""

				expectedResponse := dto.RegisterResponse{
					Message: "Register failed",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{invalid}`) //wrong request test

				req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				// when
				userController.Register(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))
			})
		})

		Context("when credentials are invalid", func() {
			It("should return 500 StatusInternalServerError and error message", func() {
				// given
				expectedToken := ""

				expectedResponse := dto.RegisterResponse{
					Message: "Register failed",
					Token:   expectedToken,
				}

				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())

				requestBody := []byte(`{"email":"test@example.com","password":"password123"}`)

				req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestBody))
				Expect(err).ToNot(HaveOccurred())

				ctx.Request = req
				ctx.Set("Content-Type", "application/json")
				ctx.Set("Accept", "application/json")
				ctx.Set("Accept-Encoding", "gzip")

				userServiceMock.EXPECT().Register("test@example.com", "password123").Return(expectedToken, errors.New("register error"))
				// when
				userController.Register(ctx)
				// then
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(MatchJSON(expectedJSON))
			})
		})
	})
})
