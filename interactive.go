package whatsapp

const (
	InteractiveTypeList = "list"
	InteractiveTypeBtn  = "button"
)

type Interactive struct {
	Header HeaderText `json:"header"`
	Body   Body       `json:"body"`
	Footer Footer     `json:"footer"`
	Action Action     `json:"action"`
	Type   string     `json:"type"`
	api    *API
}

type InteractiveBtnReq struct {
	Body struct {
		Text string `json:"text"`
	} `json:"body"`
	Action struct {
		Buttons []Button `json:"buttons"`
	} `json:"action"`
	Type string `json:"type"`
	api  *API
}

type HeaderText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type HeaderMediaLink struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

type HeaderMediaId struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
type Body struct {
	Text string `json:"text"`
}
type Footer struct {
	Text string `json:"text"`
}
type Action struct {
	Button   string    `json:"button"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Title string `json:"title"`
	Rows  []Row  `json:"rows"`
}

type Row struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Button struct {
	Type  string `json:"type"`
	Reply Reply  `json:"reply"`
}

type Reply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
