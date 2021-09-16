/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	huozjiov1alpha1 "github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis/huozj.io/v1alpha1"
	versioned "github.com/ZhengjunHUO/k8s-custom-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/ZhengjunHUO/k8s-custom-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ZhengjunHUO/k8s-custom-controller/pkg/client/listers/huozj.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FufuInformer provides access to a shared informer and lister for
// Fufus.
type FufuInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FufuLister
}

type fufuInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFufuInformer constructs a new informer for Fufu type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFufuInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFufuInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFufuInformer constructs a new informer for Fufu type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFufuInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HuozjV1alpha1().Fufus(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HuozjV1alpha1().Fufus(namespace).Watch(context.TODO(), options)
			},
		},
		&huozjiov1alpha1.Fufu{},
		resyncPeriod,
		indexers,
	)
}

func (f *fufuInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFufuInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fufuInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&huozjiov1alpha1.Fufu{}, f.defaultInformer)
}

func (f *fufuInformer) Lister() v1alpha1.FufuLister {
	return v1alpha1.NewFufuLister(f.Informer().GetIndexer())
}
