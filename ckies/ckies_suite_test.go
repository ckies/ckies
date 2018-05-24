package ckies_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCkies(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ckies Suite")
}
