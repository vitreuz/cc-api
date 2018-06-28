package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetApps", func() {
	var (
		err error
	)

	It("should fetch all the apps", func() {
		Expect(err).ToNot(HaveOccurred())
	})
})
