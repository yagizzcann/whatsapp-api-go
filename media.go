package whatsapp

var (
	MediaAudio    = "audio"
	MediaDocument = "document"
	MediaImage    = "image"
	MediaSticker  = "sticker"
	MediaVideo    = "video"
)

type Media struct {
	File     string `json:"file"`
	IsItAnID bool   `json:"is_it_an_id"`
	Type     string `json:"type"`
	api      *API
}

type MediaId struct {
	Type string `json:"-"`
	Id   string `json:"id"`
	api  *API
}

type MediaLink struct {
	Type string `json:"-"`
	Link string `json:"link"`
	api  *API
}

func (obj *Media) ToId() *MediaId {
	return &MediaId{api: obj.api, Type: obj.Type, Id: obj.File}
}
func (obj *Media) ToLink() *MediaLink {
	return &MediaLink{api: obj.api, Type: obj.Type, Link: obj.File}
}

func (m *MediaLink) ToMedia() *Media {
	return &Media{api: m.api, Type: m.Type, File: m.Link, IsItAnID: false}
}

func (m *MediaId) ToMedia() *Media {
	return &Media{api: m.api, Type: m.Type, File: m.Id, IsItAnID: true}
}
