package model

type Input struct {
	Ev     string `json:"ev"`
	Et     string `json:"et"`
	ID     string `json:"id"`
	UID    string `json:"uid"`
	MID    string `json:"mid"`
	T      string `json:"t"`
	P      string `json:"p"`
	L      string `json:"l"`
	SC     string `json:"sc"`
	ATRK1  string `json:"atrk1"`
	ATRV1  string `json:"atrv1"`
	ATRT1  string `json:"atrt1"`
	ATRK2  string `json:"atrk2"`
	ATRV2  string `json:"atrv2"`
	ATRT2  string `json:"atrt2"`
	UATRK1 string `json:"uatrk1"`
	UATRV1 string `json:"uatrv1"`
	UATRT1 string `json:"uatrt1"`
	UATRK2 string `json:"uatrk2"`
	UATRV2 string `json:"uatrv2"`
	UATRT2 string `json:"uatrt2"`
	UATRK3 string `json:"uatrk3"`
	UATRV3 string `json:"uatrv3"`
	UATRT3 string `json:"uatrt3"`
}

type Output struct {
	Event           string      `json:"event"`
	EventType       string      `json:"event_type"`
	AppID           string      `json:"app_id"`
	UserID          string      `json:"user_id"`
	MessageID       string      `json:"message_id"`
	PageTitle       string      `json:"page_title"`
	PageURL         string      `json:"page_url"`
	BrowserLanguage string      `json:"browser_language"`
	ScreenSize      string      `json:"screen_size"`
	Attributes      interface{} `json:"attributes"`
	Traits          interface{} `json:"traits"`
}
