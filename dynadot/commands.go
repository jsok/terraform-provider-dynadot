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

// Response from a domain_info command
type DomainInfoResponse struct {
	name        xml.Name   `xml:"DomainInfoContent"`
	SuccessCode int        `xml:"DomainInfoResponseHeader>SuccessCode"`
	Status      string     `xml:"DomainInfoResponseHeader>Status"`
	Error       string     `xml:"DomainInfoResponseHeader>Error"`
	DomainInfo  DomainInfo `xml:"DomainInfoContent"`
}

// Response from a set_ns command
type SetNameServerResponse struct {
	SuccessCode int    `xml:"SetNsHeader>SuccessCode"`
	Status      string `xml:"SetNsHeader>Status"`
	Error       string `xml:"SetNsHeader>Error"`
}

// Response from a set_parking command
type ParkDomainResponse struct {
	SuccessCode int    `xml:"SetParkingHeader>SuccessCode"`
	Status      string `xml:"SetParkingHeader>Status"`
	Error       string `xml:"SetParkingHeade>Error"`
}

// Response returned when some general error happens
type GeneralErrorResponse struct {
	ResponseCode int    `xml:"ResponseHeader>ResponseCode"`
	Error        string `xml:"ResponseHeader>Error"`
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
						return errors.New("encountered ServerName without first seeing ServerId")
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
