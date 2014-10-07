package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/sclevine/agouti"
	"net/http"
	"net/http/httptest"
)

var server *httptest.Server

var _ = BeforeSuite(func() {
	server = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte(`<html><body><header><h1>Page Title</h1></header><a href="#new_page">Click Me</a></body></html>`))
	}))
})

var _ = AfterSuite(func() {
	server.Close()
})

var _ = Feature("Agouti", func() {
	Scenario("Loading a page with a cookie and clicking", func() {
		cookie := CreateCookie("theName", 42, "/my-path", "example.com", false, false, 1412358590)
		page := CreatePage()
		page.Navigate(server.URL).SetCookie(cookie)

		Step("finds text in a page", func() {
			page.ShouldContainText("Page Title")
			page.Within("header").ShouldContainText("Page Title")
		})

		Step("allows tests to be scoped by chaining", func() {
			page.Within("header").Within("h1").ShouldContainText("Page Title")
		})

		Step("allows tests to be scoped by functions", func() {
			page.Within("header h1", Do(func(h1 Selection) {
				h1.ShouldContainText("Page Title")
			}))
		})

		Step("allows clicking on a link", func() {
			page.Within("a").Click()
		})
	})
})