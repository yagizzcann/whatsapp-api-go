package go_whatsapp_official

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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response Qr
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response ListQrResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response GetQrResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response Qr
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response SuccessResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
