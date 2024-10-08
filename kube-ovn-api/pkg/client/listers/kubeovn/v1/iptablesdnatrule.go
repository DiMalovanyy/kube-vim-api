// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// IptablesDnatRuleLister helps list IptablesDnatRules.
// All objects returned here must be treated as read-only.
type IptablesDnatRuleLister interface {
	// List lists all IptablesDnatRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.IptablesDnatRule, err error)
	// Get retrieves the IptablesDnatRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.IptablesDnatRule, error)
	IptablesDnatRuleListerExpansion
}

// iptablesDnatRuleLister implements the IptablesDnatRuleLister interface.
type iptablesDnatRuleLister struct {
	listers.ResourceIndexer[*kubeovnv1.IptablesDnatRule]
}

// NewIptablesDnatRuleLister returns a new IptablesDnatRuleLister.
func NewIptablesDnatRuleLister(indexer cache.Indexer) IptablesDnatRuleLister {
	return &iptablesDnatRuleLister{listers.New[*kubeovnv1.IptablesDnatRule](indexer, kubeovnv1.Resource("iptablesdnatrule"))}
}
