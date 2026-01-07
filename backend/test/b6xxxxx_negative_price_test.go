package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/mikedeni/sut-final-lab/backend/entity"
	"github.com/onsi/gomega"
)

func NegativePriceTest(t *testing.T) {
	g := gomega.NewWithT(t)

	book := entity.Books{
		Title: "Hello, World!",
		Price: -2,
		Code:  "BK123456",
	}

	result, err := govalidator.ValidateStruct(book)

	g.Expect(result).To(gomega.BeFalse())
	g.Expect(err).Should(gomega.HaveOccurred())
}
