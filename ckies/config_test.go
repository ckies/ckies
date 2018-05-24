package ckies_test

import (
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ckies/ckies/ckies"
)

var _ = Describe("Config", func() {
	var (
		c   *Config
		err error
	)

	Describe("Load()", func() {
		Context("Valid config", func() {
			BeforeEach(func() {
				file := path.Join("fixtures", "config", "valid.toml")
				c, err = Load(file)
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should have Name", func() {
				Expect(c.Name).To(Equal("Valid Name"))
			})

			It("should have Info", func() {
				Expect(c.Info).To(Equal("Example Description"))
			})

			It("should have Website", func() {
				Expect(c.Website).To(Equal("https://example.com"))
			})

			It("should have Links", func() {
				Expect(c.Links.Policy).To(Equal("/custom/policy.html"))
				Expect(c.Links.Settings).To(Equal("/custom/settings.html"))
			})

			It("should have one Cookie", func() {
				Expect(len(c.Cookies)).To(Equal(1))

				first := c.Cookies[0]
				Expect(first.Name).To(Equal("default_functional"))
				Expect(first.Info).To(Equal("Information about cookie"))
				Expect(first.Type).To(Equal(CookieTypeFunctional))
				Expect(first.Expires).To(Equal("1y"))
			})

			It("should have one Services", func() {
				Expect(len(c.Services)).To(Equal(1))

				Expect(c.Services[0]).To(Equal("valid"))
			})
		})

		Context("InValid config file", func() {
			BeforeEach(func() {
				file := path.Join("fixtures", "config", "invalid.toml")
				c, err = Load(file)
			})

			It("should return the zero-value for the config", func() {
				Expect(c).To(BeZero())
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
