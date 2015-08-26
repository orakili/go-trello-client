package trello

type Member struct {
	Id         string `json:"id"`
	AvatarHash string `json:"avatarHash"`
	Bio        string `json:"string"`
	BioData    struct {
		Emoji struct {
		} `json:"emoji"`
	} `json:"bioData"`
	Confirmed       bool     `json:"confirmed"`
	FullName        string   `json:"fullName"`
	IdPremOrgsAdmin []string `json:"idPremOrgsAdmin"`
	Initials        string   `json:"initials"`
	MemberType      string   `json:"memberType"`
	Products        []int    `json:"products"`
	Status          string   `json:"status"`
	Url             string   `json:"url"`
	Username        string   `json:"username"`
}
