// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IPPoolSpecApplyConfiguration represents a declarative configuration of the IPPoolSpec type for use
// with apply.
type IPPoolSpecApplyConfiguration struct {
	Subnet     *string  `json:"subnet,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
	IPs        []string `json:"ips,omitempty"`
}

// IPPoolSpecApplyConfiguration constructs a declarative configuration of the IPPoolSpec type for use with
// apply.
func IPPoolSpec() *IPPoolSpecApplyConfiguration {
	return &IPPoolSpecApplyConfiguration{}
}

// WithSubnet sets the Subnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subnet field is set to the value of the last call.
func (b *IPPoolSpecApplyConfiguration) WithSubnet(value string) *IPPoolSpecApplyConfiguration {
	b.Subnet = &value
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *IPPoolSpecApplyConfiguration) WithNamespaces(values ...string) *IPPoolSpecApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *IPPoolSpecApplyConfiguration) WithIPs(values ...string) *IPPoolSpecApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}
