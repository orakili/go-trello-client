package trello

type Attachment struct {
	Id        string `json:"id"`
	Bytes     int    `json:"bytes"`
	Date      string `json:"date"`
	EdgeColor string `json:"edgeColor"`
	IdMember  string `json:"idMember"`
	IsUpload  bool   `json:"isUpload"`
	MimeType  string `json:"mimeType"`
	Name      string `json:"name"`
	Previews  []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Url    string `json:"url"`
		Bytes  int    `json:"bytes"`
		Id     string `json:"_id"`
		Scaled bool   `json:"scaled"`
	} `json:"previews"`
	Url string `json:"url"`
}
