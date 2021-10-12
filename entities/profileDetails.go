package entities

type ProfileDetails struct {
	ProfileUrl     string 	`json:"profile_url"`
	ProfileName    string   `json:"profile_name"`
	ProfileHandle  string   `json:"profile_handle"`
	ProfileIconUrl string   `json:"profile_icon_url"`
	TagLine        string   `json:"tag_line"`
	Followers      string   `json:"followers"`
	PostUrls       []string `json:"post_urls"`
}
