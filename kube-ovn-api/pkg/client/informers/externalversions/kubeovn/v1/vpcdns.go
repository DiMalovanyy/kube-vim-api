// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiskubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	versioned "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/informers/externalversions/internalinterfaces"
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/listers/kubeovn/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VpcDnsInformer provides access to a shared informer and lister for
// VpcDnses.
type VpcDnsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.VpcDnsLister
}

type vpcDnsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVpcDnsInformer constructs a new informer for VpcDns type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVpcDnsInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVpcDnsInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVpcDnsInformer constructs a new informer for VpcDns type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVpcDnsInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().VpcDnses().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().VpcDnses().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.VpcDns{},
		resyncPeriod,
		indexers,
	)
}

func (f *vpcDnsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVpcDnsInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vpcDnsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.VpcDns{}, f.defaultInformer)
}

func (f *vpcDnsInformer) Lister() kubeovnv1.VpcDnsLister {
	return kubeovnv1.NewVpcDnsLister(f.Informer().GetIndexer())
}
