// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// OvnFipLister helps list OvnFips.
// All objects returned here must be treated as read-only.
type OvnFipLister interface {
	// List lists all OvnFips in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.OvnFip, err error)
	// Get retrieves the OvnFip from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.OvnFip, error)
	OvnFipListerExpansion
}

// ovnFipLister implements the OvnFipLister interface.
type ovnFipLister struct {
	listers.ResourceIndexer[*kubeovnv1.OvnFip]
}

// NewOvnFipLister returns a new OvnFipLister.
func NewOvnFipLister(indexer cache.Indexer) OvnFipLister {
	return &ovnFipLister{listers.New[*kubeovnv1.OvnFip](indexer, kubeovnv1.Resource("ovnfip"))}
}
