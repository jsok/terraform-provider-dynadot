package dynadot

import (
	"encoding/xml"
	"fmt"
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
          <Name>exmple.com</Name>
          <Expiration>1546061090000</Expiration>
          <Registration>1451366690000</Registration>
          <NameServerSettings>
            <Type>Name Servers</Type>
            <NameServers>
              <ServerId>8</ServerId>
              <ServerName>ns-1111.awsdns-11.org</ServerName>
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
	err := xml.Unmarshal([]byte(example), &resp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
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
	err := xml.Unmarshal([]byte(example), &resp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}
