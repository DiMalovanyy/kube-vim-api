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

// OvnDnatRuleInformer provides access to a shared informer and lister for
// OvnDnatRules.
type OvnDnatRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeovnv1.OvnDnatRuleLister
}

type ovnDnatRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOvnDnatRuleInformer constructs a new informer for OvnDnatRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOvnDnatRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOvnDnatRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOvnDnatRuleInformer constructs a new informer for OvnDnatRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOvnDnatRuleInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().OvnDnatRules().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeovnV1().OvnDnatRules().Watch(context.TODO(), options)
			},
		},
		&apiskubeovnv1.OvnDnatRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *ovnDnatRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOvnDnatRuleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ovnDnatRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeovnv1.OvnDnatRule{}, f.defaultInformer)
}

func (f *ovnDnatRuleInformer) Lister() kubeovnv1.OvnDnatRuleLister {
	return kubeovnv1.NewOvnDnatRuleLister(f.Informer().GetIndexer())
}
