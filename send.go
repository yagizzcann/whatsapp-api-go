package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (api *API) NewText(body string, previewUrl bool) *Text {
	return &Text{api: api, Type: "text", Body: body, PreviewUrl: previewUrl}
}

func (api *API) NewMediaId(file, _type string) *Media {
	return &Media{api: api, File: file, IsItAnID: true, Type: _type}
}

func (api *API) NewMediaLink(file, _type string) *Media {
	return &Media{api: api, File: file, IsItAnID: false, Type: _type}
}

func (api *API) NewLocation(lng, ltd, name, address string) *Location {
	return &Location{api: api, Type: "location", Longitude: lng, Latitude: ltd, Name: name, Address: address}
}

func (api *API) NewContacts(c []Contacts) *ContactsReq {
	return &ContactsReq{api: api, Type: "contacts", Contacts: c}
}

func (api *API) NewInteractive() *Interactive {
	return &Interactive{api: api, Type: "interactive"}
}
func (api *API) NewInteractiveBtnReq() *InteractiveBtnReq {
	return &InteractiveBtnReq{api: api, Type: "interactive"}
}

func (api *API) NewTextBasedTemplate() *TextBasedTemplate {
	return &TextBasedTemplate{api: api, Type: "template"}
}

func (api *API) NewMultiBasedTemplate() *MultiBasedTemplate {
	return &MultiBasedTemplate{api: api, Type: "template"}
}

func (obj *Media) Send(phoneId, to string) (*MessageResponse, error) {

	var rObj interface{}
	if obj.IsItAnID {
		rObj = obj.ToId()
	} else {
		rObj = obj.ToLink()
	}

	r, err := obj.api.send(phoneId, to, obj.Type, rObj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	err = json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *Text) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "text", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *Location) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "text", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (api *API) send(phoneId, to, _type string, obj interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/%s/messages", phoneId)

	body := map[string]interface{}{}
	body[_type] = obj
	body["type"] = _type
	body["messaging_product"] = "whatsapp"
	body["recipient_type"] = "individual"
	body["to"] = to

	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(body)

	if err != nil {
		return nil, err
	}

	res, status, err := api.request(endpoint, "POST", nil, buf)

	if err != nil {
		return nil, err
	}

	if status != 200 || status != 201 || status > 300 {
		e := ErrorResponse{}
		err = json.NewDecoder(res).Decode(&e)
		return nil, &e
	}

	r := map[string]interface{}{}
	err = json.NewDecoder(res).Decode(&r)

	fmt.Printf("things are: %v\n\n", r)

	return &r, nil
}

func (obj *ContactsReq) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *Interactive) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *InteractiveBtnReq) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *TextBasedTemplate) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

func (obj *MultiBasedTemplate) Send(phoneId, to string) (*MessageResponse, error) {
	r, err := obj.api.send(phoneId, to, "contacts", obj)
	var res MessageResponse
	jsonString, _ := json.Marshal(r)
	json.Unmarshal(jsonString, &res)
	return &res, err
}

type MessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaId  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		Id bool `json:"id"`
	} `json:"messages"`
}
