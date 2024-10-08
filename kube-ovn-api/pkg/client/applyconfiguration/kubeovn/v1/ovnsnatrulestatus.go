// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OvnSnatRuleStatusApplyConfiguration represents a declarative configuration of the OvnSnatRuleStatus type for use
// with apply.
type OvnSnatRuleStatusApplyConfiguration struct {
	Vpc        *string                                  `json:"vpc,omitempty"`
	V4Eip      *string                                  `json:"v4Eip,omitempty"`
	V6Eip      *string                                  `json:"v6Eip,omitempty"`
	V4IpCidr   *string                                  `json:"v4IpCidr,omitempty"`
	V6IpCidr   *string                                  `json:"v6IpCidr,omitempty"`
	Ready      *bool                                    `json:"ready,omitempty"`
	Conditions []OvnSnatRuleConditionApplyConfiguration `json:"conditions,omitempty"`
}

// OvnSnatRuleStatusApplyConfiguration constructs a declarative configuration of the OvnSnatRuleStatus type for use with
// apply.
func OvnSnatRuleStatus() *OvnSnatRuleStatusApplyConfiguration {
	return &OvnSnatRuleStatusApplyConfiguration{}
}

// WithVpc sets the Vpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vpc field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithVpc(value string) *OvnSnatRuleStatusApplyConfiguration {
	b.Vpc = &value
	return b
}

// WithV4Eip sets the V4Eip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4Eip field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithV4Eip(value string) *OvnSnatRuleStatusApplyConfiguration {
	b.V4Eip = &value
	return b
}

// WithV6Eip sets the V6Eip field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6Eip field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithV6Eip(value string) *OvnSnatRuleStatusApplyConfiguration {
	b.V6Eip = &value
	return b
}

// WithV4IpCidr sets the V4IpCidr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V4IpCidr field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithV4IpCidr(value string) *OvnSnatRuleStatusApplyConfiguration {
	b.V4IpCidr = &value
	return b
}

// WithV6IpCidr sets the V6IpCidr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V6IpCidr field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithV6IpCidr(value string) *OvnSnatRuleStatusApplyConfiguration {
	b.V6IpCidr = &value
	return b
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *OvnSnatRuleStatusApplyConfiguration) WithReady(value bool) *OvnSnatRuleStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *OvnSnatRuleStatusApplyConfiguration) WithConditions(values ...*OvnSnatRuleConditionApplyConfiguration) *OvnSnatRuleStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
