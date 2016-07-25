package dynadot

type DomainInfo struct {
	Name        string
	NameServers []string
}

func NewClient(apiUrl, apiKey string) (Client, error) {
	return &client{apiUrl, apiKey}, nil
}

type Client interface {
	// NameServer related calls
	ListNameServers() ([]string, error)
	AddNameServer(ns string) error
	DeleteNameServer(id string) error

	// Domain related calls
	GetDomainInfo(name string) (*DomainInfo, error)
	SetDomainNameServers(name string, ns []string) error
}

type client struct {
	url string
	key string
}

func (c *client) ListNameServers() ([]string, error) {
	return nil, nil
}

func (c *client) AddNameServer(ns string) error {
	return nil
}

func (c *client) DeleteNameServer(id string) error {
	return nil
}

func (c *client) GetDomainInfo(name string) (*DomainInfo, error) {
	return nil, nil
}

func (c *client) SetDomainNameServers(name string, ns []string) error {
	return nil
}
