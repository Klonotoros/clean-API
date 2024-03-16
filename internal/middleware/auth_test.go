package middleware

import (
	"clean-API/internal/dto"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuthenticateMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authenticate Middleware Suite")
}

var _ = Describe("Authenticate", func() {
	var (
		authMiddleware gin.HandlerFunc
		w              *httptest.ResponseRecorder
		ctx            *gin.Context
		config         dto.Config
		token          string
	)

	BeforeEach(func() {
		var err error
		w = httptest.NewRecorder()
		ctx, _ = gin.CreateTestContext(w)
		config = dto.Config{SigningSecret: "sample-secret"}
		authMiddleware = Authenticate(config)

		// Generowanie przykładowego tokena
		token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":  "test@gmail.com",
			"userId": int64(10),
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		}).SignedString([]byte(config.SigningSecret))

		Expect(err).To(BeNil())
	})

	Describe("Authenticate middleware", func() {
		Context("when a valid token is provided", func() {
			It("should set userId in context and call the next handler", func() {
				// given
				ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
				ctx.Request.Header.Set("Authorization", token)

				// when
				authMiddleware(ctx)

				// then
				userID, exists := ctx.Get("userId")
				Expect(exists).To(BeTrue())
				Expect(userID).To(Equal(int64(10)))
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})

		Context("when no token is provided", func() {
			It("should return 401 Unauthorized", func() {
				// given
				expectedResponse := dto.ErrorResponse{Message: "No token provided"}
				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())
				ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
				ctx.Request.Header.Set("Authorization", "")

				// when
				authMiddleware(ctx)

				// then
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(w.Body).To(MatchJSON(expectedJSON))
			})
		})

		Context("when an invalid token is provided", func() {
			It("should return 401 Unauthorized", func() {
				// given
				expectedResponse := dto.ErrorResponse{Message: "Wrong token"}
				expectedJSON, err := json.Marshal(expectedResponse)
				Expect(err).ToNot(HaveOccurred())
				ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)
				ctx.Request.Header.Set("Authorization", "Bearer invalid-token")

				// when
				authMiddleware(ctx)

				// then
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(w.Body).To(MatchJSON(expectedJSON))
			})
		})
	})
})
