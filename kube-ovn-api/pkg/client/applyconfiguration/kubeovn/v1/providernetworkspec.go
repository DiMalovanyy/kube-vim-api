// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ProviderNetworkSpecApplyConfiguration represents a declarative configuration of the ProviderNetworkSpec type for use
// with apply.
type ProviderNetworkSpecApplyConfiguration struct {
	DefaultInterface *string                             `json:"defaultInterface,omitempty"`
	CustomInterfaces []CustomInterfaceApplyConfiguration `json:"customInterfaces,omitempty"`
	ExcludeNodes     []string                            `json:"excludeNodes,omitempty"`
	ExchangeLinkName *bool                               `json:"exchangeLinkName,omitempty"`
}

// ProviderNetworkSpecApplyConfiguration constructs a declarative configuration of the ProviderNetworkSpec type for use with
// apply.
func ProviderNetworkSpec() *ProviderNetworkSpecApplyConfiguration {
	return &ProviderNetworkSpecApplyConfiguration{}
}

// WithDefaultInterface sets the DefaultInterface field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultInterface field is set to the value of the last call.
func (b *ProviderNetworkSpecApplyConfiguration) WithDefaultInterface(value string) *ProviderNetworkSpecApplyConfiguration {
	b.DefaultInterface = &value
	return b
}

// WithCustomInterfaces adds the given value to the CustomInterfaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CustomInterfaces field.
func (b *ProviderNetworkSpecApplyConfiguration) WithCustomInterfaces(values ...*CustomInterfaceApplyConfiguration) *ProviderNetworkSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCustomInterfaces")
		}
		b.CustomInterfaces = append(b.CustomInterfaces, *values[i])
	}
	return b
}

// WithExcludeNodes adds the given value to the ExcludeNodes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExcludeNodes field.
func (b *ProviderNetworkSpecApplyConfiguration) WithExcludeNodes(values ...string) *ProviderNetworkSpecApplyConfiguration {
	for i := range values {
		b.ExcludeNodes = append(b.ExcludeNodes, values[i])
	}
	return b
}

// WithExchangeLinkName sets the ExchangeLinkName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExchangeLinkName field is set to the value of the last call.
func (b *ProviderNetworkSpecApplyConfiguration) WithExchangeLinkName(value bool) *ProviderNetworkSpecApplyConfiguration {
	b.ExchangeLinkName = &value
	return b
}
