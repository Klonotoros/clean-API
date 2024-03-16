package utils

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HashUtils", func() {

	Describe("HashPassword", func() {
		Context("when password is provided", func() {
			It("should return hashed password and no error", func() {
				password := "password"
				hashedPassword, err := HashPassword(password)
				Expect(err).To(BeNil())
				Expect(hashedPassword).NotTo(BeEmpty())
			})
		})

		Context("when password is empty", func() {
			It("should return hashed password and no error", func() {
				password := ""
				hashedPassword, err := HashPassword(password)
				Expect(err).To(BeNil())
				Expect(hashedPassword).NotTo(BeEmpty())
			})
		})
	})

	Describe("CheckPasswordHash", func() {
		Context("when comparing valid password with its hash", func() {
			It("should return true", func() {
				// given
				password := "password"
				hashedPassword, err := HashPassword(password)
				Expect(err).To(BeNil())
				// when
				match := CheckPasswordHash(password, hashedPassword)
				// then
				Expect(match).To(BeTrue())
			})
		})

		Context("when comparing invalid password with its hash", func() {
			It("should return false", func() {
				password := "password"
				invalidPassword := "invalid_password"
				hashedPassword, err := HashPassword(password)
				Expect(err).To(BeNil())
				match := CheckPasswordHash(invalidPassword, hashedPassword)
				Expect(match).To(BeFalse())
			})
		})
	})

})
