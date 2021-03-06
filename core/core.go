// Agouti core is a general-purpose WebDriver API for Golang
package core

import (
	"fmt"
	"github.com/sclevine/agouti/core/internal/browser"
	"github.com/sclevine/agouti/core/internal/service"
	"github.com/sclevine/agouti/core/internal/types"
	"net"
	"strings"
	"time"
)

type Selection types.Selection
type Page types.Page

// Browser represents a Selenium, PhantomJS, or Chrome (via ChromeDriver) WebDriver process
type Browser interface {
	// Start launches the WebDriver process
	Start() error

	// Stop ends all sessions and stops the WebDriver process
	Stop() (nonFatal error)

	// Page returns a new WebDriver session.
	// For Selenium, browserName is the type of browser ("firefox", "safari", "chrome", etc.)
	Page(browserName ...string) (types.Page, error)
}

// Chrome returns an instance of a Chrome Browser via ChromeDriver
func Chrome() (Browser, error) {
	address, err := freeAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to locate a free port: %s", err)
	}

	port := strings.SplitN(address, ":", 2)[1]
	url := fmt.Sprintf("http://%s", address)
	command := []string{"chromedriver", "--silent", "--port=" + port}
	service := &service.Service{URL: url, Timeout: 5 * time.Second, Command: command}

	return &browser.Browser{Service: service}, nil
}

// PhantomJS returns an instance of a PhantomJS Browser
func PhantomJS() (Browser, error) {
	address, err := freeAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to locate a free port: %s", err)
	}

	url := fmt.Sprintf("http://%s", address)
	command := []string{"phantomjs", fmt.Sprintf("--webdriver=%s", address)}
	service := &service.Service{URL: url, Timeout: 5 * time.Second, Command: command}

	return &browser.Browser{Service: service}, nil
}

// Selenium returns an instance of a Selenium Browser
func Selenium() (Browser, error) {
	address, err := freeAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to locate a free port: %s", err)
	}

	port := strings.SplitN(address, ":", 2)[1]
	url := fmt.Sprintf("http://%s/wd/hub", address)
	command := []string{"selenium-server", "-port", port}
	service := &service.Service{URL: url, Timeout: 5 * time.Second, Command: command}

	return &browser.Browser{Service: service}, nil
}

func freeAddress() (string, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer listener.Close()
	return listener.Addr().String(), nil
}
