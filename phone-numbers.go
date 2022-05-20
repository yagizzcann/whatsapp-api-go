package go_whatsapp_official

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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response ListPhoneResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	var response PhoneNumber
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &response)
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
