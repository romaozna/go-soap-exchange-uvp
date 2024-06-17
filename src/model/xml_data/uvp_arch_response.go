package xml_data

import "encoding/xml"

type UvpInfo struct {
	XMLName             xml.Name            `xml:"UVPinfo"`
	Text                string              `xml:",chardata"`
	GetArchivesResponse GetArchivesResponse `xml:"GetArchivesResponse"`
}

type GetArchivesResponse struct {
	XMLName   xml.Name  `xml:"GetArchivesResponse"`
	Text      string    `xml:",chardata"`
	ArchData  ArchData  `xml:"ArchData"`
	ArchItems ArchItems `xml:"ArchItems"`
}

type ArchData struct {
	XMLName    xml.Name `xml:"ArchData"`
	Text       string   `xml:",chardata"`
	BeginTime  string   `xml:"BeginTime,attr"`
	EndTime    string   `xml:"EndTime,attr"`
	NextTime   string   `xml:"NextTime,attr"`
	PipeNumber string   `xml:"PipeNumber,attr"`
}

type ArchItems struct {
	XMLName xml.Name `xml:"ArchItems"`
	Text    string   `xml:",chardata"`
	Value   []Value  `xml:"Value"`
}

type Value struct {
	XMLName xml.Name `xml:"Value"`
	Text    string   `xml:",chardata"`
	Index   string   `xml:"Index,attr"`
	Comment string   `xml:"Comment,attr"`
	Measure string   `xml:"Measure,attr"`
}
