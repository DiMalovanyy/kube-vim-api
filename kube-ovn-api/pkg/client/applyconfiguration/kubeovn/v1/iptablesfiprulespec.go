// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IptablesFIPRuleSpecApplyConfiguration represents a declarative configuration of the IptablesFIPRuleSpec type for use
// with apply.
type IptablesFIPRuleSpecApplyConfiguration struct {
	EIP        *string `json:"eip,omitempty"`
	InternalIP *string `json:"internalIp,omitempty"`
}

// IptablesFIPRuleSpecApplyConfiguration constructs a declarative configuration of the IptablesFIPRuleSpec type for use with
// apply.
func IptablesFIPRuleSpec() *IptablesFIPRuleSpecApplyConfiguration {
	return &IptablesFIPRuleSpecApplyConfiguration{}
}

// WithEIP sets the EIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EIP field is set to the value of the last call.
func (b *IptablesFIPRuleSpecApplyConfiguration) WithEIP(value string) *IptablesFIPRuleSpecApplyConfiguration {
	b.EIP = &value
	return b
}

// WithInternalIP sets the InternalIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InternalIP field is set to the value of the last call.
func (b *IptablesFIPRuleSpecApplyConfiguration) WithInternalIP(value string) *IptablesFIPRuleSpecApplyConfiguration {
	b.InternalIP = &value
	return b
}
