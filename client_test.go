package client

import (
	"errors"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

const (
	testURL      = "https://10.0.0.1"
	testToken    = "test_token"
	testEndpoint = "/test"
)

// ErrReader implements the io.Reader interface and fails on Read.
type ErrReader struct{}

// Read mocks failing io.Reader test cases.
func (r ErrReader) Read(buf []byte) (int, error) {
	return 0, errors.New("fail")
}

func testClient() Client {
	client, err := NewClient(testURL, testToken)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	gock.InterceptClient(client.HttpClient)
	return *client
}

func TestClient_Get(t *testing.T) {
	defer gock.Off()
	client := testClient()

	var err error

	// Success
	gock.New(testURL).Get(testEndpoint).Reply(200)
	_, err = client.Get(testEndpoint)
	assert.NoError(t, err)

	// HTTP error
	gock.New(testURL).Get(testEndpoint).ReplyError(errors.New("fail"))
	_, err = client.Get(testEndpoint)
	assert.Error(t, err)

	// Invalid HTTP status code
	gock.New(testURL).Get(testEndpoint).Reply(405)
	_, err = client.Get(testEndpoint)
	assert.Error(t, err)

	// Error decoding response body
	gock.New(testURL).Get(testEndpoint).Reply(200).Map(
		func(res *http.Response) *http.Response {
			res.Body = io.NopCloser(ErrReader{})
			return res
		})
	_, err = client.Get(testEndpoint)
	assert.Error(t, err)
}

func TestClient_Post(t *testing.T) {
	defer gock.Off()
	client := testClient()

	var err error

	// Success
	gock.New(testURL).Post(testEndpoint).Reply(200)
	_, err = client.Post(testEndpoint, "{}")
	assert.NoError(t, err)

	// HTTP error
	gock.New(testURL).Post(testEndpoint).ReplyError(errors.New("fail"))
	_, err = client.Post(testEndpoint, "{}")
	assert.Error(t, err)

	// Invalid HTTP status code
	gock.New(testURL).Post(testEndpoint).Reply(405)
	_, err = client.Post(testEndpoint, "{}")
	assert.Error(t, err)

	// Error decoding response body
	gock.New(testURL).Post(testEndpoint).Reply(200).Map(
		func(res *http.Response) *http.Response {
			res.Body = io.NopCloser(ErrReader{})
			return res
		})
	_, err = client.Post(testEndpoint, "{}")
	assert.Error(t, err)
}

func TestClient_Put(t *testing.T) {
	defer gock.Off()
	client := testClient()

	var err error

	// Success
	gock.New(testURL).Put(testEndpoint).Reply(200)
	_, err = client.Put(testEndpoint, "{}")
	assert.NoError(t, err)

	// HTTP error
	gock.New(testURL).Put(testEndpoint).ReplyError(errors.New("fail"))
	_, err = client.Put(testEndpoint, "{}")
	assert.Error(t, err)

	// Invalid HTTP status code
	gock.New(testURL).Put(testEndpoint).Reply(405)
	_, err = client.Put(testEndpoint, "{}")
	assert.Error(t, err)

	// Error decoding response body
	gock.New(testURL).Put(testEndpoint).Reply(200).Map(
		func(res *http.Response) *http.Response {
			res.Body = io.NopCloser(ErrReader{})
			return res
		})
	_, err = client.Put(testEndpoint, "{}")
	assert.Error(t, err)
}

func TestClient_Patch(t *testing.T) {
	defer gock.Off()
	client := testClient()

	var err error

	// Success
	gock.New(testURL).Patch(testEndpoint).Reply(200)
	_, err = client.Patch(testEndpoint, "{}")
	assert.NoError(t, err)

	// HTTP error
	gock.New(testURL).Patch(testEndpoint).ReplyError(errors.New("fail"))
	_, err = client.Patch(testEndpoint, "{}")
	assert.Error(t, err)

	// Invalid HTTP status code
	gock.New(testURL).Patch(testEndpoint).Reply(405)
	_, err = client.Patch(testEndpoint, "{}")
	assert.Error(t, err)

	// Error decoding response body
	gock.New(testURL).Patch(testEndpoint).Reply(200).Map(
		func(res *http.Response) *http.Response {
			res.Body = io.NopCloser(ErrReader{})
			return res
		})
	_, err = client.Patch(testEndpoint, "{}")
	assert.Error(t, err)
}

func TestClient_Delete(t *testing.T) {
	defer gock.Off()
	client := testClient()

	var err error

	// Success
	gock.New(testURL).Delete(testEndpoint).Reply(204)
	_, err = client.Delete(testEndpoint)
	assert.NoError(t, err)

	// HTTP error
	gock.New(testURL).Delete(testEndpoint).ReplyError(errors.New("fail"))
	_, err = client.Delete(testEndpoint)
	assert.Error(t, err)
}
