package harlogconverter

type Har struct {
	Log HarLog `json:"log"`
}

type HarLog struct {
	Version string     `json:"version"`
	Entries []HarEntry `json:"entries"`
}

type HarEntry struct {
	Request  HarRequest  `json:"request"`
	Response HarResponse `json:"response"`
}

type HarRequest struct {
	Method      string             `json:"method"`
	Url         string             `json:"url"`
	Headers     []HarRequestHeader `json:"headers"`
	QueryString []QueryString      `json:"queryString"`
	PostData    HarRequestPostData `json:"postData"`
}

type HarResponse struct {
	Status  int                `json:"status"`
	Headers []HarRequestHeader `json:"headers"`
	Content HarResponseContent `json:"content"`
}

type HarResponseContent struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}

type HarRequestHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type QueryString struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type HarRequestPostData struct {
	MimeType string `json:"mimeType"`
	Text     string `json:"text"`
}
