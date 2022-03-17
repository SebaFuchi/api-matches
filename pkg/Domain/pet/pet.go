package pet

type Pet struct {
	Token      string `json:"token,omitempty"`
	OwnerToken string `json:"ownerToken,omitempty"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Sex        string `json:"sex"`
	Image      string `json:"image"`
}
