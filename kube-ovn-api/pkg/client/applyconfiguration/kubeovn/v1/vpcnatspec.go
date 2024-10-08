// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// VpcNatSpecApplyConfiguration represents a declarative configuration of the VpcNatSpec type for use
// with apply.
type VpcNatSpecApplyConfiguration struct {
	Vpc             *string                          `json:"vpc,omitempty"`
	Subnet          *string                          `json:"subnet,omitempty"`
	ExternalSubnets []string                         `json:"externalSubnets,omitempty"`
	LanIP           *string                          `json:"lanIp,omitempty"`
	Selector        []string                         `json:"selector,omitempty"`
	Tolerations     []corev1.Toleration              `json:"tolerations,omitempty"`
	Affinity        *corev1.Affinity                 `json:"affinity,omitempty"`
	QoSPolicy       *string                          `json:"qosPolicy,omitempty"`
	BgpSpeaker      *VpcBgpSpeakerApplyConfiguration `json:"bgpSpeaker,omitempty"`
}

// VpcNatSpecApplyConfiguration constructs a declarative configuration of the VpcNatSpec type for use with
// apply.
func VpcNatSpec() *VpcNatSpecApplyConfiguration {
	return &VpcNatSpecApplyConfiguration{}
}

// WithVpc sets the Vpc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Vpc field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithVpc(value string) *VpcNatSpecApplyConfiguration {
	b.Vpc = &value
	return b
}

// WithSubnet sets the Subnet field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subnet field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithSubnet(value string) *VpcNatSpecApplyConfiguration {
	b.Subnet = &value
	return b
}

// WithExternalSubnets adds the given value to the ExternalSubnets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ExternalSubnets field.
func (b *VpcNatSpecApplyConfiguration) WithExternalSubnets(values ...string) *VpcNatSpecApplyConfiguration {
	for i := range values {
		b.ExternalSubnets = append(b.ExternalSubnets, values[i])
	}
	return b
}

// WithLanIP sets the LanIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LanIP field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithLanIP(value string) *VpcNatSpecApplyConfiguration {
	b.LanIP = &value
	return b
}

// WithSelector adds the given value to the Selector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Selector field.
func (b *VpcNatSpecApplyConfiguration) WithSelector(values ...string) *VpcNatSpecApplyConfiguration {
	for i := range values {
		b.Selector = append(b.Selector, values[i])
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *VpcNatSpecApplyConfiguration) WithTolerations(values ...corev1.Toleration) *VpcNatSpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithAffinity(value corev1.Affinity) *VpcNatSpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithQoSPolicy sets the QoSPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QoSPolicy field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithQoSPolicy(value string) *VpcNatSpecApplyConfiguration {
	b.QoSPolicy = &value
	return b
}

// WithBgpSpeaker sets the BgpSpeaker field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BgpSpeaker field is set to the value of the last call.
func (b *VpcNatSpecApplyConfiguration) WithBgpSpeaker(value *VpcBgpSpeakerApplyConfiguration) *VpcNatSpecApplyConfiguration {
	b.BgpSpeaker = value
	return b
}
