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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis/huozj.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FufuLister helps list Fufus.
// All objects returned here must be treated as read-only.
type FufuLister interface {
	// List lists all Fufus in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fufu, err error)
	// Fufus returns an object that can list and get Fufus.
	Fufus(namespace string) FufuNamespaceLister
	FufuListerExpansion
}

// fufuLister implements the FufuLister interface.
type fufuLister struct {
	indexer cache.Indexer
}

// NewFufuLister returns a new FufuLister.
func NewFufuLister(indexer cache.Indexer) FufuLister {
	return &fufuLister{indexer: indexer}
}

// List lists all Fufus in the indexer.
func (s *fufuLister) List(selector labels.Selector) (ret []*v1alpha1.Fufu, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fufu))
	})
	return ret, err
}

// Fufus returns an object that can list and get Fufus.
func (s *fufuLister) Fufus(namespace string) FufuNamespaceLister {
	return fufuNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FufuNamespaceLister helps list and get Fufus.
// All objects returned here must be treated as read-only.
type FufuNamespaceLister interface {
	// List lists all Fufus in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fufu, err error)
	// Get retrieves the Fufu from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Fufu, error)
	FufuNamespaceListerExpansion
}

// fufuNamespaceLister implements the FufuNamespaceLister
// interface.
type fufuNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fufus in the indexer for a given namespace.
func (s fufuNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Fufu, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fufu))
	})
	return ret, err
}

// Get retrieves the Fufu from the indexer for a given namespace and name.
func (s fufuNamespaceLister) Get(name string) (*v1alpha1.Fufu, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fufu"), name)
	}
	return obj.(*v1alpha1.Fufu), nil
}
