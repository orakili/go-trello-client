package trello

type Organization struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Desc    string   `json:"desc"`
	Members []Member `json:"members"`
}
