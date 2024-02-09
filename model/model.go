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
	ATRK3  string `json:"atrk3"`
	ATRV3  string `json:"atrv3"`
	ATRT3  string `json:"atrt3"`
	ATRK4  string `json:"atrk4"`
	ATRV4  string `json:"atrv4"`
	ATRT4  string `json:"atrt4"`
	ATRK5  string `json:"atrk5"`
	ATRV5  string `json:"atrv5"`
	ATRT5  string `json:"atrt5"`
	ATRK6  string `json:"atrk6"`
	ATRV6  string `json:"atrv6"`
	ATRT6  string `json:"atrt6"`
	UATRK1 string `json:"uatrk1"`
	UATRV1 string `json:"uatrv1"`
	UATRT1 string `json:"uatrt1"`
	UATRK2 string `json:"uatrk2"`
	UATRV2 string `json:"uatrv2"`
	UATRT2 string `json:"uatrt2"`
	UATRK3 string `json:"uatrk3"`
	UATRV3 string `json:"uatrv3"`
	UATRT3 string `json:"uatrt3"`
	UATRK4 string `json:"uatrk4"`
	UATRV4 string `json:"uatrv4"`
	UATRT4 string `json:"uatrt4"`
	UATRK5 string `json:"uatrk5"`
	UATRV5 string `json:"uatrv5"`
	UATRT5 string `json:"uatrt5"`
	UATRK6 string `json:"uatrk6"`
	UATRV6 string `json:"uatrv6"`
	UATRT6 string `json:"uatrt6"`
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
