package whatsapp

import (
	"encoding/json"
	"fmt"
)

func (api *API) ListPhoneNumbers(businessAccountId string) (*ListPhoneResponse, error) {

	endpoint := fmt.Sprintf("/%s/phone_numbers", businessAccountId)

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

	var response ListPhoneResponse
	json.Unmarshal(res, &response)
	return &response, nil

}

func (api *API) SinglePhoneNumber(phoneId string) (*PhoneNumber, error) {

	endpoint := fmt.Sprintf("/%s", phoneId)

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

	var response PhoneNumber
	json.Unmarshal(res, &response)
	return &response, nil

}

type ListPhoneResponse struct {
	Data []PhoneNumber `json:"data"`
}

type GetPhoneResponse struct {
	Data PhoneNumber `json:"data"`
}

type PhoneNumber struct {
	VerifiedName       string `json:"verified_name"`
	DisplayPhoneNumber string `json:"display_phone_number"`
	Id                 string `json:"id"`
	QualityRating      string `json:"quality_rating"`
}
