package trello

type Action struct {
	Id              string      `json:"id"`
	IdMemberCreator string      `json:"idMemberCreator"`
	Data            interface{} `json:"data"`
	Type            string      `json:"type"`
	Date            string      `json:"date"`
	Member          []Member    `json:"member,omitempty"`
	MemberCreator   Member      `json:"memberCreator"`
}
