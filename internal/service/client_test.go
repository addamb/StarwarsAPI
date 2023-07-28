package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/addamb/starwarsapi/internal/service"
	"github.com/addamb/starwarsapi/mocks"
	"github.com/addamb/starwarsapi/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_Client(t *testing.T) {
	t.Parallel()
	url := "http://test.com/"
	assert := assert.New(t)

	swapiResponse := models.SwapiResponse{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results:  []models.SwapiHuman{},
	}

	var out []byte
	out, _ = json.Marshal(swapiResponse)

	clientResp := ioutil.NopCloser(bytes.NewReader(out))

	tests := map[string]struct {
		ClientResp    io.ReadCloser
		ClientError   error
		ExpectedError bool
	}{
		"error swapi request": {
			ClientError:   errors.New("test error client"),
			ExpectedError: true,
		},
		"success": {
			ClientResp:    clientResp,
			ExpectedError: false,
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)
			swapi := service.NewSwapiClient(client)

			request, _ := http.NewRequest(http.MethodGet, url, nil)
			client.EXPECT().Do(request).Return(
				&http.Response{
					Body: tcase.ClientResp,
				},
				tcase.ClientError,
			)

			resp, err := swapi.SendSwapiRequest(url)

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.ClientResp, resp)
		})
	}
}
