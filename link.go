package charm

type LinkStatus int

const (
	LinkStatusInit LinkStatus = iota
	LinkStatusTokenCreated
	LinkStatusRequested
	LinkStatusRequestDenied
	LinkStatusSameAccount
	LinkStatusDifferentAccount
	LinkStatusSuccess
	LinkStatusTimedOut
	LinkStatusError
	LinkStatusValidTokenRequest
	LinkStatusInvalidTokenRequest
)

type Link struct {
	Token         string     `json:"token"`
	RequestPubKey string     `json:"request_pub_key"`
	RequestAddr   string     `json:"request_addr"`
	Host          string     `json:"host"`
	Port          int        `json:"port"`
	Status        LinkStatus `json:"status"`
}

type LinkerMessage struct {
	Message string `json:"message"`
}

type LinkHandler interface {
	DisplayCode(*Link)
	ConfirmRequest(*Link) bool
	DisplayFinalStatus(*Link)
}
