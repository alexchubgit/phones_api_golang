package keys

type Key struct {
	IDTOKEN int    `json:""`
	Number  string `json:"number"`
	Idowner int    `json:"idowner"`
	Status  string `json:"status"`
	Comment string `json:"comment"`
}
