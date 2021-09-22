package controller

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fmt"
	"time"
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const MAX_RETRY int = 3

type Controller struct {
	clientset	kubernetes.Interface
	// List and watch the delta of certain resource and trigger the event handler
	// normally is to send the key to the queue
	informer	cache.SharedIndexInformer
	// dedicated to the controller to receive event(key)
	queue		workqueue.RateLimitingInterface
}

// Sent to queue by informer if match the condition 
type Event struct {
        key          string
        eventType    string
        resourceType string
}

func Start(client kubernetes.Interface) {
	ctlr := newController(client, "pod") 	

	chStop := make(chan struct{})
	defer close(chStop)
	defer fmt.Println("Receive interrupt signal, stop controller, cleanup ...")

	go ctlr.Run(chStop)

	// receive interrupt signal, close chStop channel to stop controller
	chIntrpt := make(chan os.Signal, 1)
        signal.Notify(chIntrpt, os.Interrupt, syscall.SIGTERM)

        <-chIntrpt
}

func newController(client kubernetes.Interface, resourceType string) *Controller {
	// Set the list watch functions, clientset is needed here
	lw := cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return client.CoreV1().Pods(metav1.NamespaceAll).List(context.TODO(), options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return client.CoreV1().Pods(metav1.NamespaceAll).Watch(context.TODO(), options)
		},
	}

	// use SharedIndexInformer
	informer := cache.NewSharedIndexInformer(&lw, &corev1.Pod{}, 0, cache.Indexers{})

	var event Event
	var err error
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// register event handler to the informer
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event.eventType, event.resourceType = "create", resourceType
			event.key, err = cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				log.Printf("Resource %v[type %v] created.\n", event.key, event.resourceType)
				queue.Add(event)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			event.eventType, event.resourceType = "update", resourceType
			event.key, err = cache.MetaNamespaceKeyFunc(oldObj)
			if err == nil {
				log.Printf("Resource %v[type %v] updated.\n", event.key, event.resourceType)
				queue.Add(event)
			}
		},
		DeleteFunc: func(obj interface{}) {
			event.eventType, event.resourceType = "delete", resourceType
			event.key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				log.Printf("Resource %v[type %v] deleted.\n", event.key, event.resourceType)
				queue.Add(event)
			}
		},
	})
	
	return &Controller {
		clientset: client,
		informer: informer,
		queue: queue,
	}
}

// Implement cache.Controller interface
func (c *Controller) Run(chStop <-chan struct{}) {
        defer utilruntime.HandleCrash()
        defer c.queue.ShutDown()

        go c.informer.Run(chStop)

        if !cache.WaitForCacheSync(chStop, c.HasSynced) {
                utilruntime.HandleError(fmt.Errorf("Waiting for caches to sync receive timeout!"))
                return
        }

        log.Println("Controller up!")

	// loop until interrupted
        wait.Until(c.workerUp, time.Second, chStop)
}


// Implement cache.Controller interface
func (c *Controller) HasSynced() bool {
        return c.informer.HasSynced()
}

// Implement cache.Controller interface
func (c *Controller) LastSyncResourceVersion() string {
        return c.informer.LastSyncResourceVersion()
}

func (c *Controller) workerUp() {
	for c.hasNext() {}
}

// Retrieve the event from the queue, and handle it
func (c *Controller) hasNext() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Done(item)
	if err := c.Process(item.(Event)); err == nil {
		// item processed ok
		c.queue.Forget(item)
	}else if c.queue.NumRequeues(item) < MAX_RETRY {
		// Process failed, still able to retry, add it back to queue
		log.Printf("[WARN] Failed to process %s: %v. Retrying ...\n", item.(Event).key, err)
		c.queue.AddRateLimited(item)
	}else{
		// Process failed, no quota to retry, throw it away
		log.Printf("[WARN] Failed to process %s: %v. No retry left, Abort !\n", item.(Event).key, err)
		c.queue.Forget(item)
		utilruntime.HandleError(err)
	}

	return true
}

func (c *Controller) Process(event Event) error {
	// send key in event to informer's indexer to retrieve item in shared cache 
	item, _ , err := c.informer.GetIndexer().GetByKey(event.key)
	if err != nil {
		return fmt.Errorf("Unable to get object[key %s] from store: %v", event.key, err)
	}

	handler := &DefaultHandler{}
	// Call handler depends on the event type
	switch event.eventType {
	case "create":
		handler.Created(item)
	case "update":
		handler.Updated(item)
	case "delete":
		handler.Deleted(item)
	}

	return nil
}
