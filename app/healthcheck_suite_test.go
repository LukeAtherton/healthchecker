package healthchecker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

/*
Ginkgo generated this function
to kick off our unit tests. Hook into it
to define factories.
*/
func TestHealthcheckApp(t *testing.T) {
	defineFactories()
	RegisterFailHandler(Fail)
	RunSpecs(t, "(Cluster) Healthchecker App Suite")
}

func defineFactories() {

}
