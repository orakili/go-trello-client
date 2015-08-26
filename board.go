package trello

type Board struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	DescData struct {
		Emoji struct {
		} `json:"emoji"`
	} `json:"descData"`
	Closed         bool   `json:"closed"`
	IdOrganization string `json:"idOrganization"`
	Invited        bool   `json:"invited"`
	Pinned         bool   `json:"pinned"`
	Starred        bool   `json:"starred"`
	Url            string `json:"url"`
	Prefs          struct {
		PermissionLevel       string `json:"permissionLevel"`
		Voting                string `json:"voting"`
		Comments              string `json:"comments"`
		Invitations           string `json:"invitations"`
		SelfJoin              bool   `json:"selfJoin"`
		CardCovers            bool   `json:"cardCovers"`
		CardAging             string `json:"cardAging"`
		CalendarFeedEnabled   bool   `json:"calendarFeedEnabled"`
		Background            string `json:"background"`
		BackgroundColor       string `json:"backgroundColor"`
		BackgroundImage       string `json:"backgroundImage"`
		BackgroundImageScaled bool   `json:"backgroundImageScaled"`
		BackgroundTile        bool   `json:"backgroundTile"`
		BackgroundBrightness  string `json:"backgroundBrightness"`
		CanBePublic           bool   `json:"canBePublic"`
		CanBeOrg              bool   `json:"canBeOrg"`
		CanBePrivate          bool   `json:"canBePrivate"`
		CanInvite             bool   `json:"canInvite"`
	} `json:"prefs"`
	Invitations []interface{} `json:"invitations"`
	memberships []struct {
		Id          string `json:"id"`
		IdMember    string `json:"idMember"`
		MemberType  string `json:"memberType"`
		Unconfirmed bool   `json:"unconfirmed"`
		Deactivated bool   `json:"deactivated"`
	} `json:"memberships"`
	ShortLink  string `json:"shortLink"`
	Subscribed bool   `json:"subscribed"`
	LabelNames struct {
		Green  string `json:"green"`
		Yellow string `json:"yellow"`
		Orange string `json:"orange"`
		Red    string `json:"red"`
		Purple string `json:"purple"`
		Blue   string `json:"blue"`
		Sky    string `json:"sky"`
		Lime   string `json:"lime"`
		Pink   string `json:"pink"`
		Black  string `json:"black"`
	} `json:"labelNames"`
	PowerUps         []string `json:"powerUps"`
	DateLastActivity string   `json:"dateLastActivity"`
	DateLastView     string   `json:"dateLastView"`
	ShortUrl         string   `json:"shortUrl"`
	Lists            []List   `json:"lists"`
	Cards            []Card   `json:"cards"`
	Labels           []Label  `json:"labels"`
	Actions          []Action `json:"actions"`
}
