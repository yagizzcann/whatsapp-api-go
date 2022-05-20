package whatsapp

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

var (
	DefaultURI     = "https://graph.facebook.com"
	DefaultVersion = "v13.0"
)

type API struct {
	Token               string
	WebHookVerification string
	URI                 string `default:"https://graph.facebook.com"`
	Version             string `default:"v13.0"`
	client              *http.Client
}

type ErrorResponse struct {
	Data struct {
		Message      string `json:"message"`
		Type         string `json:"type"`
		Code         int32  `json:"code"`
		ErrorSubCode int32  `json:"error_subcode"`
		FbTraceId    string `json:"fbtrace_id"`
		ErrorData    struct {
			MessagingProduct string `json:"messaging_product"`
			Details          string `json:"details"`
		} `json:"error_data"`
	} `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return e.Data.ErrorData.Details
}

func (api *API) request(endpoint string, method string, params map[string]interface{}, body io.Reader) (result *bytes.Buffer, status int, err error) {
	if api.client == nil {
		api.client = &http.Client{}
	}

	uri := fmt.Sprintf("%s/%s%s", api.URI, api.Version, endpoint)

	if params != nil {
		uri = fmt.Sprintf("%s?", uri)

		for k, v := range params {
			uri = fmt.Sprintf("%s&%s=%s", uri, k, v)
		}
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.Token))

	resp, err := api.client.Do(req)

	if err != nil {
		return
	}

	status = resp.StatusCode

	result = &bytes.Buffer{}
	if _, err = io.Copy(result, resp.Body); err != nil {
		return
	}

	return
}
