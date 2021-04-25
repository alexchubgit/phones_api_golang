package certs

type Cert struct {
	IDCERT    int    `json:"idcert"`
	Filename  string `json:"filename"`
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
}
