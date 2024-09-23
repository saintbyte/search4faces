package search4faces

type JSONRPCRequeest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      string `json:"id"`
	Params  struct {
	} `json:"params"`
}
type JSONRPCResponse struct {
	Jsonrpc string   `json:"jsonrpc"`
	Result  struct{} `json:"result"`
	ID      string   `json:"id"`
}
type RateResult struct {
	Apikey    string   `json:"apikey"`
	Limit     int      `json:"limit"`
	Remaining int      `json:"remaining"`
	Enddate   string   `json:"enddate"`
	Speed     int      `json:"speed"`
	Allowed   []string `json:"allowed"`
	Disabled  string   `json:"disabled"`
}

type DetectRequestParams struct {
	Image string `json:"image"`
}

type DetectResponse struct {
	Image string           `json:"image"`
	Faces []DetectFaceItem `json:"faces"`
}
type DetectFaceItem struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	Lm1X   int `json:"lm1_x"`
	Lm1Y   int `json:"lm1_y"`
	Lm2X   int `json:"lm2_x"`
	Lm2Y   int `json:"lm2_y"`
	Lm3X   int `json:"lm3_x"`
	Lm3Y   int `json:"lm3_y"`
	Lm4X   int `json:"lm4_x"`
	Lm4Y   int `json:"lm4_y"`
	Lm5X   int `json:"lm5_x"`
	Lm5Y   int `json:"lm5_y"`
}
type SearchRequestParams struct {
	Image string `json:"image"`
	Face  struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Width  int `json:"width"`
		Height int `json:"height"`
		Lm1X   int `json:"lm1_x"`
		Lm1Y   int `json:"lm1_y"`
		Lm2X   int `json:"lm2_x"`
		Lm2Y   int `json:"lm2_y"`
		Lm3X   int `json:"lm3_x"`
		Lm3Y   int `json:"lm3_y"`
		Lm4X   int `json:"lm4_x"`
		Lm4Y   int `json:"lm4_y"`
		Lm5X   int `json:"lm5_x"`
		Lm5Y   int `json:"lm5_y"`
	} `json:"face"`
	Source  string `json:"source"`
	Hidden  bool   `json:"hidden"`
	Results int    `json:"results"`
	Lang    string `json:"lang"`
}
type SearchResponse struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	Score       string `json:"score"`
	Face        string `json:"face"`
	Profile     string `json:"profile"`
	Photo       string `json:"photo"`
	PhotoX      int    `json:"photo_x"`
	PhotoY      int    `json:"photo_y"`
	PhotoWidth  int    `json:"photo_width"`
	PhotoHeight int    `json:"photo_height"`
	Source      string `json:"source"`
	Age         int    `json:"age"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MaidenName  string `json:"maiden_name"`
	City        string `json:"city"`
	Country     string `json:"country"`
}
