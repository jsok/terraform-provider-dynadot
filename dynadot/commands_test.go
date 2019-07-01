package dynadot

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListDomain(t *testing.T) {
	example := `
<ListDomainInfoResponse>
  <ListDomainInfoHeader>
    <ResponseCode>0</ResponseCode>
    <Status>success</Status>
  </ListDomainInfoHeader>
  <ListDomainInfoContent>
    <DomainInfoList>
      <DomainInfo>
        <Domain>
          <Name>example.com</Name>
          <Expiration>1546061090000</Expiration>
          <Registration>1451366690000</Registration>
          <NameServerSettings>
            <Type>Name Servers</Type>
            <NameServers>
              <ServerId>8</ServerId>
              <ServerName>ns-55.awsdns-555.org</ServerName>
              <ServerId>7</ServerId>
              <ServerName>ns-333.awsdns-17.com</ServerName>
              <ServerId>9</ServerId>
              <ServerName>ns-2222.awsdns-48.co.uk</ServerName>
              <ServerId>6</ServerId>
              <ServerName>ns-444.awsdns-32.net</ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
              <ServerId></ServerId>
              <ServerName></ServerName>
            </NameServers>
          </NameServerSettings>
          <Whois>
            <Registrant>
              <ContactId>263900</ContactId>
            </Registrant>
            <Admin>
              <ContactId>263900</ContactId>
            </Admin>
            <Technical>
              <ContactId>263900</ContactId>
            </Technical>
            <Billing>
              <ContactId>263900</ContactId>
            </Billing>
          </Whois>
          <Locked>yes</Locked>
          <Disabled>no</Disabled>
          <UdrpLocked>no</UdrpLocked>
          <RegistrantUnverified>no</RegistrantUnverified>
          <Hold>no</Hold>
          <Privacy>full</Privacy>
          <isForSale>no</isForSale>
          <RenewOption>no renew option</RenewOption>
          <Note></Note>
          <Folder>
            <FolderId>-1</FolderId>
            <FolderName>(no folder)</FolderName>
          </Folder>
        </Domain>
      </DomainInfo>
    </DomainInfoList>
  </ListDomainInfoContent>
</ListDomainInfoResponse>
`
	var resp ListDomainInfoResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, "", resp.Error)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, 1, len(resp.DomainInfoList))
	firstDomainInfo := resp.DomainInfoList[0]
	assert.Equal(t, "example.com", firstDomainInfo.Name)
	assert.Equal(t, 4, len(firstDomainInfo.NameServers))
	assert.Contains(t, firstDomainInfo.NameServers, NameServer{Id: "8", Name: "ns-55.awsdns-555.org"})
}

func TestDomainInfo(t *testing.T) {
	example := `
<DomainInfoResponse>
    <DomainInfoResponseHeader>
        <SuccessCode>0</SuccessCode>
        <Status>success</Status>
    </DomainInfoResponseHeader>
    <DomainInfoContent>
        <Domain>
            <Name>lggccbbj.com</Name>
            <Expiration>1574494426000</Expiration>
            <Registration>1542958426000</Registration>
            <NameServerSettings>
                <Type>Name Servers</Type>
                <NameServers>
                    <ServerId>1338851</ServerId>
                    <ServerName>ns-1351.awsdns-40.org</ServerName>
                    <ServerId>1338852</ServerId>
                    <ServerName>ns-943.awsdns-53.net</ServerName>
                    <ServerId>1338853</ServerId>
                    <ServerName>ns-422.awsdns-52.com</ServerName>
                    <ServerId>1338854</ServerId>
                    <ServerName>ns-1653.awsdns-14.co.uk</ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                    <ServerId></ServerId>
                    <ServerName></ServerName>
                </NameServers>
            </NameServerSettings>
            <Whois>
                <Registrant>
                    <ContactId>371939</ContactId>
                </Registrant>
                <Admin>
                    <ContactId>371939</ContactId>
                </Admin>
                <Technical>
                    <ContactId>371939</ContactId>
                </Technical>
                <Billing>
                    <ContactId>371939</ContactId>
                </Billing>
            </Whois>
            <Locked>yes</Locked>
            <Disabled>no</Disabled>
            <UdrpLocked>no</UdrpLocked>
            <RegistrantUnverified>no</RegistrantUnverified>
            <Hold>no</Hold>
            <Privacy>none</Privacy>
            <isForSale>no</isForSale>
            <RenewOption>manual renewal</RenewOption>
            <Note>flagged dsp domain</Note>
            <Folder>
                <FolderId>-1</FolderId>
                <FolderName>(no folder)</FolderName>
            </Folder>
        </Domain>
    </DomainInfoContent>
</DomainInfoResponse>
`
	var resp DomainInfoResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, 0, resp.SuccessCode)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, 4, len(resp.DomainInfo.NameServers))
	assert.Equal(t, NameServer{Id: "1338851", Name: "ns-1351.awsdns-40.org"}, resp.DomainInfo.NameServers[0])
}

func TestServerList(t *testing.T) {
	example := `
<ServerListResponse>
  <ServerListHeader>
    <ResponseCode>0</ResponseCode>
    <Status>success</Status>
  </ServerListHeader>
  <ServerListContent>
    <NameServerList>
      <List>
        <Server>
          <ServerId>9</ServerId>
          <ServerName>ns-2222.awsdns-48.co.uk</ServerName>
          <ServerIp></ServerIp>
        </Server>
        <Server>
          <ServerId>8</ServerId>
          <ServerName>ns-1111.awsdns-11.org</ServerName>
          <ServerIp></ServerIp>
        </Server>
        <Server>
          <ServerId>7</ServerId>
          <ServerName>ns-333.awsdns-17.com</ServerName>
          <ServerIp></ServerIp>
        </Server>
        <Server>
          <ServerId>6</ServerId>
          <ServerName>ns-444.awsdns-32.net</ServerName>
          <ServerIp></ServerIp>
        </Server>
      </List>
    </NameServerList>
  </ServerListContent>
</ServerListResponse>
`
	var resp ServerListResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, 4, len(resp.NameServers))
	assert.Contains(t, resp.NameServers, NameServer{Id: "9", Name: "ns-2222.awsdns-48.co.uk"})
	assert.Contains(t, resp.NameServers, NameServer{Id: "8", Name: "ns-1111.awsdns-11.org"})
	assert.Contains(t, resp.NameServers, NameServer{Id: "7", Name: "ns-333.awsdns-17.com"})
	assert.Contains(t, resp.NameServers, NameServer{Id: "6", Name: "ns-444.awsdns-32.net"})
}

func TestNameServers(t *testing.T) {
	example := `
<SetNsResponse>
  <SetNsHeader>
    <SuccessCode>0</SuccessCode>
    <Status>success</Status>
  </SetNsHeader>
</SetNsResponse>
`
	var resp SetNameServerResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, 0, resp.SuccessCode)
	assert.Equal(t, "success", resp.Status)
}

func TestNameServersError(t *testing.T) {
	example := `
<SetNsResponse>
    <SetNsHeader>
        <SuccessCode>-1</SuccessCode>
        <Status>error</Status>
        <Error>Cannot set nameservers to these domains ( gglpsd.com).</Error>
    </SetNsHeader>
</SetNsResponse>
`
	var resp SetNameServerResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, -1, resp.SuccessCode)
	assert.Equal(t, "error", resp.Status)
	assert.Equal(t, "Cannot set nameservers to these domains ( gglpsd.com).", resp.Error)
}

func TestParkDomain(t *testing.T) {
	example := `
<SetParkingResponse>
    <SetParkingHeader>
        <SuccessCode>-1</SuccessCode>
        <Status>failure</Status>
    </SetParkingHeader>
</SetParkingResponse>
`
	var resp ParkDomainResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, -1, resp.SuccessCode)
	assert.Equal(t, "failure", resp.Status)
}

func TestGeneralErrorResponse(t *testing.T) {
	example := `
<Response>
    <ResponseHeader>
        <ResponseCode>-1</ResponseCode>
        <Error>invalid key</Error>
    </ResponseHeader>
</Response>

`
	var resp GeneralErrorResponse
	unmarshal(example, &resp, t)
	assert.Equal(t, -1, resp.ResponseCode)
	assert.Equal(t, "invalid key", resp.Error)
}

func unmarshal(xmlString string, resp interface{}, t *testing.T) {
	err := xml.Unmarshal([]byte(xmlString), &resp)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", resp)
}
