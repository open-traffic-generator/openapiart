import (
	"encoding/xml"
)

type ObjectRequest struct {
	XMLName   xml.Name   `xml:"object-request"`
	RequestId uint       `xml:"request-id,attr"`
	Address   string     `xml:"address,attr"`
	Type      string     `xml:"type,attr"`
	Version   string     `xml:"version,attr"`
	Method    string     `xml:"method"`
	Argument  []Argument `xml:"argument"`
	ArgXml    string     `xml:",innerxml"`
}

type Argument struct {
	XMLName   xml.Name          `xml:"argument"`
	Name      string            `xml:"name,attr"`
	Direction ArgumentDirection `xml:"direction,attr"`
	Type      string            `xml:"type,attr"`
	Member    []Member          `xml:"member"`
	Item      []Item            `xml:"item"`
	Value     string            `xml:",chardata"`
}

type Member struct {
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Member  []Member `xml:"member"`
	Item    []Item   `xml:"item"`
	Value   string   `xml:",chardata"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Type    string   `xml:"type,attr"`
	Member  []Member `xml:"member"`
	Value   string   `xml:",chardata"`
}

type ObjectResponse struct {
	XMLName   xml.Name `xml:"object-response"`
	RequestId uint     `xml:"request-id,attr"`
	RetVal    []RetVal `xml:"retval"`
	ResultXml string   `xml:",innerxml"`
}

type RetVal struct {
	XMLName xml.Name `xml:"retval"`
	Type    string   `xml:"type,attr"`
	Member  []Member `xml:"member"`
	Item    []Item   `xml:"item"`
	Value   string   `xml:",chardata"`
}
