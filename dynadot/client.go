package dynadot

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NewClient(apiUrl, apiKey string) (Client, error) {
	return &client{apiUrl, apiKey}, nil
}

type Client interface {
	GetDomainInfo(name string) (*DomainInfo, error)
	SetDomainNameServers(name string, ns []string) error
	ParkDomain(name string) error
}

type client struct {
	url string
	key string
}

// https://www.dynadot.com/domain/api3.html#set_ns
// https://api.dynadot.com/api3.xml?key=mykey&command=set_ns&domain=domain1.com,domain2.com&ns0=ns1.hostns.com&ns1=ns2.hostns.com
func (c *client) SetDomainNameServers(domain string, nameServers []string) error {
	request, e := c.newApiRequest("set_ns")
	if e != nil {
		return e
	}
	withParam(request, "domain", domain)
	for index, nameServer := range nameServers {
		withParam(request, fmt.Sprintf("ns%d", index), nameServer)
	}
	rawResponse, e := c.doRequest(request)
	if e != nil {
		return e
	}
	var setNameServerResponse SetNameServerResponse
	err := xml.Unmarshal(rawResponse, &setNameServerResponse)
	if err != nil {
		return e
	}
	if setNameServerResponse.SuccessCode != 0 {
		return fmt.Errorf("operation failed with successCode=%d, API error message=%s", setNameServerResponse.SuccessCode, setNameServerResponse.Error)
	}
	return nil
}

//https://www.dynadot.com/domain/api3.html#set_parking
//https://api.dynadot.com/api3.xml?key=mykey&command=set_parking&domain=domain1.com&with_ads=no
func (c *client) ParkDomain(domain string) error {
	request, e := c.newApiRequest("set_parking")
	if e != nil {
		return e
	}
	withParam(request, "domain", domain)
	rawResponse, e := c.doRequest(request)
	if e != nil {
		return e
	}
	var parkDomainResponse ParkDomainResponse
	err := xml.Unmarshal(rawResponse, &parkDomainResponse)
	if err != nil {
		return err
	}
	if parkDomainResponse.SuccessCode != 0 {
		return fmt.Errorf("operation failed with successCode=%d, API error message=%s", parkDomainResponse.SuccessCode, parkDomainResponse.Error)
	}
	return nil
}

// https://www.dynadot.com/domain/api3.html#domain_info
// https://api.dynadot.com/api3.xml?key=mykey&command=domain_info&domain=domain1.com
func (c *client) GetDomainInfo(domain string) (*DomainInfo, error) {
	request, e := c.newApiRequest("domain_info")
	if e != nil {
		return nil, e
	}
	withParam(request, "domain", domain)
	rawResponse, e := c.doRequest(request)
	if e != nil {
		return nil, e
	}
	var domainInfoResponse DomainInfoResponse
	err := xml.Unmarshal(rawResponse, &domainInfoResponse)
	if err != nil {
		return nil, e
	}
	if domainInfoResponse.SuccessCode != 0 {
		return nil, fmt.Errorf("operation failed with successCode=%d, API error message=%s", domainInfoResponse.SuccessCode, domainInfoResponse.Error)
	}
	return &domainInfoResponse.DomainInfo, nil
}

func withParam(request *http.Request, paramName string, paramValue string) *http.Request {
	query := request.URL.Query()
	query.Set(paramName, paramValue)
	request.URL.RawQuery = query.Encode()
	return request
}

func (c *client) newApiRequest(command string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.url, nil)
	withParam(req, "key", c.key)
	withParam(req, "command", command)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *client) doRequest(req *http.Request) ([]byte, error) {
	log.Printf("[DEBUG] Dynadot Request %+v\n", req)
	httpClient := &http.Client{Timeout: time.Second * 10}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Dynadot Response %+v\n", res)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		log.Printf("[DEBUG] Dynadot response body %s\n", string(body))

		var maybeGeneralError GeneralErrorResponse
		err := xml.Unmarshal(body, &maybeGeneralError)
		if err != nil {
			return nil, err
		}
		if maybeGeneralError.ResponseCode != 0 {
			return nil, fmt.Errorf("operation failed with response code=%d, API error message=%s", maybeGeneralError.ResponseCode, maybeGeneralError.Error)
		}

		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
