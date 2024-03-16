package utils

import (
	"github.com/golang-jwt/jwt/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("JwtUtils", func() {

	var (
		signingSecret string
		userId        int64
	)

	BeforeEach(func() {
		signingSecret = "secret"
		userId = 123
	})

	var _ = Describe("VerifyToken", func() {
		Context("when provided token and signing secret are valid", func() {
			It("should return user ID and no error", func() {

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"userId": userId,
				})

				tokenString, err := token.SignedString([]byte(signingSecret))
				Expect(err).To(BeNil())
				//when
				parsedUserId, err := VerifyToken(tokenString, signingSecret)
				//then
				Expect(err).To(BeNil())
				Expect(parsedUserId).To(Equal(userId))
			})
		})

		Context("when provided token is invalid", func() {
			It("should return error", func() {

				invalidToken := "invalid.token"

				_, err := VerifyToken(invalidToken, signingSecret)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("could not parse token"))
			})
		})

		Context("when provided signing secret is invalid", func() {
			It("should return error", func() {

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"userId": userId,
				})

				invalidSigningSecret := "invalid_secret"
				tokenString, err := token.SignedString([]byte(invalidSigningSecret))
				Expect(err).To(BeNil())
				//when
				_, err = VerifyToken(tokenString, signingSecret)
				//then
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
