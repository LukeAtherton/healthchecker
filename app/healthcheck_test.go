package healthchecker_test

import (
	. "github.com/lukeatherton/healthchecker/app"

	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

/*
Convert JSON data into a slice.
*/
func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}

/*
Convert JSON data into a map.
*/
func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

/*
Server unit tests.
*/
var _ = Describe("(Cluster) Healthchecker App", func() {
	//var server *martini.Martini
	var server *gin.Engine
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		server = NewRouter()

		// Record HTTP responses.
		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {

	})

	Describe("GET /healthcheck", func() {

		// Set up a new GET request before every test
		// in this describe block.
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/healthcheck", nil)
		})

		Context("when service is running", func() {
			It("returns a status code of 200", func() {
				server.ServeHTTP(recorder, request)
				fmt.Printf("%v\n", recorder)
				Expect(recorder.Code).To(Equal(200))
			})
		})
	})

})
