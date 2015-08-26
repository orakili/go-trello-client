package trello

type Label struct {
	Id      string `json:"id"`
	IdBoard string `json:"idBoard"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Uses    int    `json:"uses"`
}
