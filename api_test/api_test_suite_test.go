package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApiTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApiTest Suite")
}
