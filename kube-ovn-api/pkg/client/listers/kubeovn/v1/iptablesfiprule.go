// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// IptablesFIPRuleLister helps list IptablesFIPRules.
// All objects returned here must be treated as read-only.
type IptablesFIPRuleLister interface {
	// List lists all IptablesFIPRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.IptablesFIPRule, err error)
	// Get retrieves the IptablesFIPRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.IptablesFIPRule, error)
	IptablesFIPRuleListerExpansion
}

// iptablesFIPRuleLister implements the IptablesFIPRuleLister interface.
type iptablesFIPRuleLister struct {
	listers.ResourceIndexer[*kubeovnv1.IptablesFIPRule]
}

// NewIptablesFIPRuleLister returns a new IptablesFIPRuleLister.
func NewIptablesFIPRuleLister(indexer cache.Indexer) IptablesFIPRuleLister {
	return &iptablesFIPRuleLister{listers.New[*kubeovnv1.IptablesFIPRule](indexer, kubeovnv1.Resource("iptablesfiprule"))}
}
