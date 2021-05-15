package _struct

type Kakao struct {
	Title     string   `json:"Title"`
	EndDate   string   `json:"EndDate"`
	Location  string   `json:"Location"`
	JobGroups []string `json:"JobGroups"`
	Company   string   `json:"Company"`
	JobType   string   `json:"JobType"`
}

type Indeed struct {
	Id       string
	Title    string
	Company  string
	Location string
}