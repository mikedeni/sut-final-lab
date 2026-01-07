package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/mikedeni/sut-final-lab/backend/entity"
	"github.com/onsi/gomega"
)

func NegativeCodeTest(t *testing.T) {
	g := gomega.NewWithT(t)

	book := entity.Books{
		Title: "Hello, World!",
		Price: 5000,
		Code:  "test",
	}

	result, err := govalidator.ValidateStruct(book)

	g.Expect(result).To(gomega.BeFalse())
	g.Expect(err).Should(gomega.HaveOccurred())
}
