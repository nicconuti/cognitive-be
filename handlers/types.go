package handlers

type TestResponse struct {
	Grid    [][]string `json:"grid"`
	Timeout int        `json:"timeout"`
}
