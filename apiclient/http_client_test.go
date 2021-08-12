package apiclient_test

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/http-api/apiclient"
	"github.com/http-api/apiclient/mocks"
)

func TestHashResponseSuccess(t *testing.T) {
	client := &mocks.MockClient{}
	addr := "given_url_sample"

	response := []byte(`{"sample_json_response"}`)
	responseBody := ioutil.NopCloser(bytes.NewReader(response))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       responseBody,
		}, nil
	}
	responseHash := md5.Sum(response)
	expectedHash := hex.EncodeToString(responseHash[:])

	_, hash, err := apiclient.HashResponse(addr, client)
	if hash == "" {
		t.Errorf("TestHashResponseSuccess failed, expected hash value %v, got %v.", expectedHash, hash)
	}

	if err != nil {
		t.Errorf("TestHashResponseSuccess failed, expected no error value, got %v.", err)
	}
}

func TestHashResponseFailToValidateURL(t *testing.T) {
	client := &mocks.MockClient{}
	addr := "}}}"

	_, _, err := apiclient.HashResponse(addr, client)

	if err == nil {
		t.Errorf("TestHashResponseFailToValidateURL failed, expected error, got nothing")
	}
}

func TestHashResponseFailToReach(t *testing.T) {
	client := &mocks.MockClient{}
	addr := "given_url_sample"

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New(
			"Error from web server",
		)
	}

	_, _, err := apiclient.HashResponse(addr, client)

	if err == nil {
		t.Errorf("TestHashResponseFailToValidateURL failed, expected error, got nothing")
	}
}
