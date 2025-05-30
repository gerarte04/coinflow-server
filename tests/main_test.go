package tests

import (
	"coinflow/coinflow-server/pkg/testutils"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"testing"
)

func setup() error {
	var err error
	cookieJar, err = cookiejar.New(&cookiejar.Options{
		PublicSuffixList: &testutils.PublicSuffixList{},
	})

	if err != nil {
		return err
	}

	cli = &http.Client{Jar: cookieJar}

	return nil
}

func teardown() error {
	return nil
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatalf("setup: %s", err.Error())
	}

	code := m.Run()

	if err := teardown(); err != nil {
		log.Fatalf("teardown: %s", err.Error())
	}

	os.Exit(code)
}
