package controller

import (
	"log"
	"k8s.io/api/core/v1"
	hzjv1alpha1 "github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis/huozj.io/v1alpha1"
)

type Handler interface {
	Created(item interface{})
	Updated(item interface{})
	Deleted(item interface{})
}

// default handler simply print out the pod's basic information
type DefaultHandler struct {}

func (dh *DefaultHandler) Created(item interface{}) {
	p := item.(*v1.Pod)
	log.Printf("Pod %v/%v created!\n", p.Namespace, p.Name)
}

func (dh *DefaultHandler) Updated(item interface{}) {
	p := item.(*v1.Pod)
	log.Printf("Pod %v/%v updated!\n", p.Namespace, p.Name)
}

func (dh *DefaultHandler) Deleted(item interface{}) {
	log.Printf("Pod deleted!\n")
}

// print out type fufu's information
type HzjHandler struct {}

func (hh *HzjHandler) Created(item interface{}) {
	f := item.(*hzjv1alpha1.Fufu)
	log.Printf("Fufu %v/%v created!\n", f.Namespace, f.Name)
}

func (hh *HzjHandler) Updated(item interface{}) {
	f := item.(*hzjv1alpha1.Fufu)
	log.Printf("Fufu %v/%v updated!\n", f.Namespace, f.Name)
}

func (hh *HzjHandler) Deleted(item interface{}) {
	log.Printf("Fufu deleted!\n")
}
