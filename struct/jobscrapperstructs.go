package _struct

type Kakao struct {
	LastExistDate string   `json:"LastExistDate"`
	Title         string   `json:"Title"`
	EndDate       string   `json:"EndDate"`
	StartDate     string   `json:"StartDate"`
	Location      string   `json:"Location"`
	JobGroups     []string `json:"JobGroups"`
	Company       string   `json:"Company"`
	JobType       string   `json:"JobType"`
	Url           string   `json:"Url"`
	Id            string   `json:"Id"`
}

type Line struct {
	LastExistDate string `json:"LastExistDate"`
	Title         string `json:"Title"`
	EndDate       string `json:"EndDate"`
	StartDate     string `json:"StartDate"`
	Location      string `json:"Location"`
	Company       string `json:"Company"`
	Url           string `json:"Url"`
	Id            string `json:"Id"`
}

type Indeed struct {
	Id       string
	Title    string
	Company  string
	Location string
}
