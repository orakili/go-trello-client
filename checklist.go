package trello

type Checklist struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	IdBoard    string  `json:"idBoard"`
	IdCard     string  `json:"idCard"`
	Pos        float32 `json:"pos"`
	CheckItems []struct {
		State    string      `json:"state"`
		Id       string      `json:"id"`
		Name     string      `json:"name"`
		NameData interface{} `json:"nameData"`
		Pos      float32     `json:"pos"`
	} `json:"checkItems"`
}
