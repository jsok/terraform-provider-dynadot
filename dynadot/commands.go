package dynadot

import (
	"encoding/xml"
	"errors"
)

// Response from a server_list command
type ServerListResponse struct {
	name         xml.Name     `xml:"ServerListResponse"`
	ResponseCode int          `xml:"ServerListHeader>ResponseCode"`
	Status       string       `xml:"ServerListHeader>Status"`
	Error        string       `xml:"ServerListHeader>Error"`
	NameServers  []NameServer `xml:"ServerListContent>NameServerList>List>Server"`
}

// Response from a list_domain command
type ListDomainInfoResponse struct {
	name           xml.Name     `xml:"ListDomainInfoResponse"`
	ResponseCode   int          `xml:"ListDomainInfoHeader>ResponseCode"`
	Status         string       `xml:"ListDomainInfoHeader>Status"`
	Error          string       `xml:"ListDomainInfoHeader>Error"`
	DomainInfoList []DomainInfo `xml:"ListDomainInfoContent>DomainInfoList>DomainInfo"`
}

type DomainInfo struct {
	name        xml.Name    `xml:"DomainInfo"`
	Name        string      `xml:"Domain>Name"`
	NameServers NameServers `xml:"Domain>NameServerSettings>NameServers"`
}

type NameServers []NameServer

type NameServer struct {
	Id   string `xml:"ServerId"`
	Name string `xml:"ServerName"`
}

func (n *NameServers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var current *NameServer
	for {
		token, err := d.Token()
		if token == nil {
			break
		}
		if err != nil {
			return err
		}
		if t, ok := token.(xml.StartElement); ok {
			switch t.Name.Local {
			case "ServerId":
				{
					current = &NameServer{}
					if err := d.DecodeElement(&current.Id, &t); err != nil {
						return err
					}
				}
			case "ServerName":
				{
					if current == nil {
						errors.New("Encountered ServerName without first seeing ServerId")
					}
					if err := d.DecodeElement(&current.Name, &t); err != nil {
						return err
					}
					if current.Id != "" && current.Name != "" {
						*n = append(*n, *current)
					}
				}
			}
		}
	}
	return nil
}
