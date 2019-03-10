package network

type NetworkDriver interface {
	Name() string
	Create(subnet string, name string) (*Network, error)
	Load(name string) (*Network, error)
	Delete(*Network) error
	Connect(endpointId string, networkName string, portMappings []string) (*Endpoint, error)
	Disconnect(endpoint *Endpoint) error
}
