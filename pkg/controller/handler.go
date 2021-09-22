package controller

import (
	"log"
	"k8s.io/api/core/v1"
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
	p := item.(*v1.Pod)
	log.Printf("Pod %v/%v deleted!\n", p.Namespace, p.Name)
}
