package structure

type KeyStruct struct {
	Code int     `json:"code"`
	Data KeyData `json:"data"`
}

type KeyData struct {
	Command string `json:"command"`
}
