package main

import (
	"flag"
	"log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"

	"fmt"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var kubeconfigPath string
	// defaultPath := os.Getenv("HOME") + "/.kube/config"
	
	flag.StringVar(&kubeconfigPath, "kubeconfig", "", "Path to kubeconfig")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	pod, err := clientset.CoreV1().Pods("kube-system").Get(context.TODO(), "kube-apiserver-k8s-node-0", metav1.GetOptions{})
        fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", pod.Namespace, pod.Name, pod.Status.Phase)
}
