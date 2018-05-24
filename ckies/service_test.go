package ckies_test

import (
	"path"

	. "github.com/ckies/ckies/ckies"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var (
		s   *Service
		err error
	)

	Describe("GetFromPrefix()", func() {
		Context("Valid service identifier", func() {
			BeforeEach(func() {
				prefix := path.Join("fixtures", "service")
				s, err = GetServiceFromPrefix("valid", prefix)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should have Name", func() {
				Expect(s.Name).To(Equal("Valid Service Name"))
			})

			It("should have Info", func() {
				Expect(s.Info).To(Equal("A short description"))
			})

			It("should have Website", func() {
				Expect(s.Website).To(Equal("https://example.com"))
			})

			It("should have four Cookies", func() {
				Expect(len(s.Cookies)).To(Equal(4))
			})

			It("should have one Necessary Cookie", func() {
				Expect(len(s.Necessary())).To(Equal(1))

				first := s.Necessary()[0]
				Expect(first.Name).To(Equal("test_necessary"))
				Expect(first.Info).To(Equal("Some cookie. IDK what purpose."))
				Expect(first.Type).To(Equal(CookieTypeNecessary))
				Expect(first.Expires).To(Equal("10y"))
			})

			It("should have one Function Cookie", func() {
				Expect(len(s.Functional())).To(Equal(1))

				first := s.Functional()[0]
				Expect(first.Name).To(Equal("test_functional"))
				Expect(first.Info).To(Equal("Some cookie. IDK what purpose."))
				Expect(first.Type).To(Equal(CookieTypeFunctional))
				Expect(first.Expires).To(Equal("1d"))
			})

			It("should have one Performance Cookie", func() {
				Expect(len(s.Performance())).To(Equal(1))

				first := s.Performance()[0]
				Expect(first.Name).To(Equal("test_performance"))
				Expect(first.Info).To(Equal("Some cookie. IDK what purpose."))
				Expect(first.Type).To(Equal(CookieTypePerformance))
				Expect(first.Expires).To(Equal("1h"))
			})

			It("should have one Marketing Cookie", func() {
				Expect(len(s.Marketing())).To(Equal(1))

				first := s.Marketing()[0]
				Expect(first.Name).To(Equal("test_marketing"))
				Expect(first.Info).To(Equal("Some cookie. IDK what purpose."))
				Expect(first.Type).To(Equal(CookieTypeMarketing))
				Expect(first.Expires).To(Equal("1m"))
			})
		})

		Context("InValid service name", func() {
			BeforeEach(func() {
				s, err = GetServiceFromPrefix("invalid", "")
			})

			It("should return the zero-value for the service", func() {
				Expect(s).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
