package postman

type Collection struct {
	Info CollectionInfo   `json:"info"`
	Item []CollectionItem `json:"item"`
}

type CollectionInfo struct {
	Name       string `json:"name"`
	Schema     string `json:"schema"`
	PostmanId  string `json:"_postman_id"`
	ExporterId string `json:"_exporter_id"`
}

type CollectionItem struct {
	Name     string         `json:"name"`
	Request  ItemRequest    `json:"request"`
	Response []ItemResponse `json:"response"`
	Event    []ItemEvent    `json:"event"`
}

type ItemRequest struct {
	Method string          `json:"method"`
	Header []RequestHeader `json:"header"`
}

type RequestHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type RequestUrl struct {
	Raw   string            `json:"raw"`
	Host  []string          `json:"host"`
	Path  []string          `json:"path"`
	Query []RequestUrlQuery `json:"query"`
}

type RequestUrlQuery struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestBody struct {
	Mode    string             `json:"mode"`
	Raw     string             `json:"raw"`
	Options RequestBodyOptions `json:"options"`
}

type RequestBodyOptions struct {
	Raw map[string]string `json:"raw"`
}

type ItemEvent struct {
	Listen string        `json:"listen"`
	Script []EventScript `json:"script"`
}

type EventScript struct {
	Type string   `json:"type"`
	Exec []string `json:"exec"`
}

type ItemResponse struct {
}
