package trello

type Card struct {
	Id     string `json:"id"`
	Badges struct {
		Votes              int    `json:"votes"`
		ViewingMemberVoted bool   `json:"viewingMemberVoted"`
		Subscribed         bool   `json:"subscribed"`
		Fogbugz            string `json:"fogbugz"`
		CheckItems         int    `json:"checkItems"`
		CheckItemsChecked  int    `json:"checkItemsChecked"`
		Comments           int    `json:"comments"`
		Attachments        int    `json:"attachments"`
		Description        bool   `json:"description"`
		Due                string `json:"due"`
	} `json:"badges"`
	CheckItemStates []struct {
		IdCheckItem string `json:"idCheckItem"`
		State       string `json:"state"`
	} `json:"checkItemStates"`
	Closed           bool   `json:"closed"`
	DateLastActivity string `json:"dateLastActivity"`
	Desc             string `json:"desc"`
	DescData         struct {
		Emoji struct {
		} `json:"emoji"`
	} `json:"descData"`
	Due                   string       `json:"due"`
	Email                 string       `json:"email"`
	IdBoard               string       `json:"idBoard"`
	IdChecklists          []string     `json:"idChecklists"`
	IdLabels              []string     `json:"idLabels"`
	IdList                string       `json:"idList"`
	IdMembers             []string     `json:"idMembers"`
	IdShort               int          `json:"idShort"`
	IdAttachmentCover     string       `json:"idAttachmentCover"`
	ManualCoverAttachment bool         `json:"manualCoverAttachment"`
	Labels                []Label      `json:"labels"`
	Name                  string       `json:"name"`
	Pos                   float32      `json:"post"`
	ShortUrl              string       `json:"shortUrl"`
	Url                   string       `json:"url"`
	Attachments           []Attachment `json:"attachments"`
	Members               []Member     `json:"members"`
	MembersVoted          []Member     `json:"membersVoted"`
	Actions               []Action     `json:"actions"`
}
