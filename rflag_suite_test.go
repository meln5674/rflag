package rflag_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRflag(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rflag Suite")
}
