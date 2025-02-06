package client

import (
	"net/http"
	"testing"
)

type MockClient struct {
	numberOfGetRequests int
	lastRequested       string
}

func (m *MockClient) RoundTrip(req *http.Request) (*http.Response, error) {
	m.numberOfGetRequests++
	m.lastRequested = req.URL.String()

	// Return a fake response
	return &http.Response{
		StatusCode: 200,
		Body:       http.NoBody,
		Header:     make(http.Header),
	}, nil
}

func TestClient(t *testing.T) {
	t.Run("it makes a http get request to the input string", func(t *testing.T) {
		mock := &MockClient{}
		client := &Client{
			http.Client{Transport: mock},
		}

		websiteAddress := "test"

		client.GetData(websiteAddress)

		if mock.numberOfGetRequests != 1 {
			t.Fatalf("Expected get to be called 1 time")
		}

		if mock.lastRequested != websiteAddress {
			t.Fatalf("Expected last request to be %s but called %s", websiteAddress, mock.lastRequested)
		}
	})
}
