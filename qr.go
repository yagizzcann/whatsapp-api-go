package whatsapp

import (
	"encoding/json"
	"fmt"
)

func (api *API) CreateQr(phoneId, message, format string) (*Qr, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls", phoneId)

	params := map[string]interface{}{}
	params["prefilled_message"] = message
	params["generate_qr_image"] = format
	params["access_token"] = api.Token

	res, status, err := api.request(endpoint, "POST", params, nil)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response Qr
	json.Unmarshal(res, &response)
	return &response, nil

}

func (api *API) ListQr(phoneId string) (*ListQrResponse, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls", phoneId)

	params := map[string]interface{}{}
	params["access_token"] = api.Token

	res, status, err := api.request(endpoint, "GET", params, nil)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response ListQrResponse
	json.Unmarshal(res, &response)
	return &response, nil

}

func (api *API) GetQr(phoneId, qrId string) (*GetQrResponse, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls/%s", phoneId, qrId)

	params := map[string]interface{}{}
	params["access_token"] = api.Token

	res, status, err := api.request(endpoint, "GET", params, nil)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response GetQrResponse
	json.Unmarshal(res, &response)
	return &response, nil

}

func (api *API) UpdateQr(phoneId, qrId, prefilledMessage string) (*Qr, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls/%s", phoneId, qrId)

	params := map[string]interface{}{}
	params["access_token"] = api.Token
	params["prefilled_message"] = prefilledMessage

	res, status, err := api.request(endpoint, "POST", params, nil)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response Qr
	json.Unmarshal(res, &response)
	return &response, nil

}

func (api *API) DeleteQr(phoneId, qrId string) (*SuccessResponse, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls/%s", phoneId, qrId)

	params := map[string]interface{}{}
	params["access_token"] = api.Token

	res, status, err := api.request(endpoint, "DELETE", params, nil)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response SuccessResponse
	json.Unmarshal(res, &response)
	return &response, nil

}

type ListQrResponse struct {
	Data []Qr `json:"data"`
}

type GetQrResponse struct {
	Data Qr `json:"data"`
}

type SuccessResponse struct {
	Success bool `json:"cuccess"`
}

type Qr struct {
	Code             string `json:"code"`
	PrefilledMessage string `json:"prefilled_message"`
	DeepLinkUrl      string `json:"deep_link_url"`
	QrImageUrl       string `json:"qr_image_url"`
}
