package xml_data

import "encoding/xml"

type ArchiveResponse struct {
	XMLName             xml.Name `xml_data:"UVPinfo"`
	Text                string   `xml_data:",chardata"`
	GetArchivesResponse struct {
		Text     string `xml_data:",chardata"`
		ArchData struct {
			Text       string `xml_data:",chardata"`
			BeginTime  string `xml_data:"BeginTime,attr"`
			EndTime    string `xml_data:"EndTime,attr"`
			NextTime   string `xml_data:"NextTime,attr"`
			PipeNumber string `xml_data:"PipeNumber,attr"`
		} `xml_data:"ArchData"`
		ArchItems struct {
			Text  string `xml_data:",chardata"`
			Value []struct {
				Text    string `xml_data:",chardata"`
				Index   string `xml_data:"Index,attr"`
				Comment string `xml_data:"Comment,attr"`
				Measure string `xml_data:"Measure,attr"`
			} `xml_data:"Value"`
		} `xml_data:"ArchItems"`
	} `xml_data:"GetArchivesResponse"`
}
