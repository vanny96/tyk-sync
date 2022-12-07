package dashboard

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/TykTechnologies/tyk-sync/clients/http_client/mocks"
	"github.com/TykTechnologies/tyk-sync/clients/objects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAPI(t *testing.T) {
	tcs := []struct {
		testName string

		givenAPI objects.DBApiDefinition

		setupCalls       func() *mocks.HTTPClient
		expectedMetaData string
		expectErr        error
	}{
		{
			testName: "API CREATION ok",
			givenAPI: objects.DBApiDefinition{},
			setupCalls: func() *mocks.HTTPClient {
				client := &mocks.HTTPClient{}

				resp := &http.Response{Status: "200", StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":"123","_id":"123","api_id": "ok"}`))}
				resp2 := &http.Response{Status: "200", StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewBufferString(`{"status": "ok"}`))}

				client.On("Do", mock.Anything, mock.Anything).Return(resp, nil)

				client.On("Do", mock.Anything, mock.Anything).Return(resp2, nil)

				return client
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.testName, func(t *testing.T) {
			c, err := NewDashboardClient("url", "secret", "org")

			assert.Nil(t, err)

			mockedClient := tc.setupCalls()
			c.SetHTTPClient(mockedClient)

			tc.givenAPI.APIID = "123"
			actualMeta, actualError := c.CreateAPI(&tc.givenAPI)

			assert.Equal(t, tc.expectErr, actualError)
			assert.Equal(t, tc.expectedMetaData, actualMeta)

			mockedClient.AssertExpectations(t)
		})
	}

}
