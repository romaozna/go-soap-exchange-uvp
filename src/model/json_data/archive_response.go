package json_data

type ArchiveResponse struct {
	StartTime  string `json:"BeginTime"`
	EndTime    string `json:"EndTime"`
	PipeNumber string `json:"PipeNumber"`
	Val        Values `json:"Data"`
}

type Values []struct {
	Index   string `json:"Index"`
	Value   string `json:"Value"`
	Comment string `json:"Comment"`
	Measure string `json:"Measure"`
}
