package obs

import (
	"github.com/pivotal-cf/brokerapi"
)

// Bind implematation
func (b *OBSBroker) Bind(instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, nil
}
