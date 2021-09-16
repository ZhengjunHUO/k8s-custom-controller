package main

import (
	"flag"
	"log"
	"fmt"
	"context"

	"github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis/huozj.io/v1alpha1"
	clientv1alpha1 "github.com/ZhengjunHUO/k8s-custom-controller/pkg/client/clientset/versioned/typed/huozj.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes/scheme"
)

func main() {
	var kubeconfigPath string
	
	flag.StringVar(&kubeconfigPath, "kubeconfig", "", "Path to kubeconfig")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	v1alpha1.AddToScheme(scheme.Scheme)

	clientset, err := clientv1alpha1.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

        fufus, err := clientset.Fufus("default").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                panic(err)
        }

        fmt.Printf("Fufus found: %+v\n", fufus)
}
