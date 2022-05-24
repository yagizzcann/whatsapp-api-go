package whatsapp

import (
	"encoding/json"
	"fmt"
)

type Status struct {
	MessageId string `json:"message_id"`
	api       *API
}

func (api *API) NewStatus(mId string) *Status {
	return &Status{api: api, MessageId: mId}
}

func (obj *Status) MakeQr(phoneId, message, format string) (*StatusResponse, error) {

	endpoint := fmt.Sprintf("/%s/message_qrdls", phoneId)

	body := make(map[string]interface{})
	body["status"] = "read"
	body["messaging_product"] = "whatsapp"
	body["message_id"] = obj.MessageId

	res, status, err := obj.api.request(endpoint, "POST", nil, body)
	if err != nil {

		return nil, err
	}

	if status != 200 {
		e := ErrorResponse{}
		json.Unmarshal(res, &e)
		return nil, &e
	}

	var response StatusResponse
	json.Unmarshal(res, &response)
	return &response, nil

}

type StatusResponse struct {
	Success bool `json:"success"`
}
