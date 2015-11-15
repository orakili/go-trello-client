package trello

type List struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Closed     bool    `json:"closed"`
	IdBoard    string  `json:"idBoard"`
	Pos        float32 `json:"pos"`
	Subscribed bool    `json:"subscribed"`
	Cards      []Card  `json:"cards,omitempty"`
}
