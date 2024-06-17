package xml_data

import (
	"encoding/xml"
)

type ArchiveRequest struct {
	XMLName     xml.Name `xml:"UVPinfo"`
	Text        string   `xml:",chardata"`
	GetArchives struct {
		Text    string `xml:",chardata"`
		Options struct {
			Text          string `xml:",chardata"`
			ReturnMeasure string `xml:"ReturnMeasure,attr"`
			ReturnHeader  string `xml:"ReturnHeader,attr"`
			ReturnComment string `xml:"ReturnComment,attr"`
			UserMeasure   string `xml:"UserMeasure,attr"`
			StartTime     string `xml:"StartTime,attr"`
			StopTime      string `xml:"StopTime,attr"`
			PipeNumber    string `xml:"PipeNumber,attr"`
		} `xml:"Options"`
	} `xml:"GetArchives"`
}
