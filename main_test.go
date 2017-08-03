package main

import "testing"

func TestValidateProxyHost(t *testing.T) {
	hosts := []string{
		"127.0.0.1",
		"http:127.0.0.1",
	}

	for _, host := range hosts {
		_, err := validateProxyHost(host)
		if err == nil {
			t.Fatal(host, " should error")
		}
	}
}

func TestComposeURL(t *testing.T) {
	host := "http://127.0.0.1"
	urls := []string{
		"a/b/c",
		"/a/b/c",
	}
	results := []string{
		"http://127.0.0.1/a/b/c",
		"http://127.0.0.1/a/b/c",
	}

	for i := range urls {
		// if strings.Compare()
		if results[i] != composeURL(host, results[i]) {
			t.Fatal(host, " + ", urls[i], " should = ", results[i])
		}
	}

}
