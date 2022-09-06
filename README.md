# k8s-custom-controller
Put theory into practice, build a custom controller myself in order to know better about Kubernetes' mechanism.

### About kubernetes/code-generator:
- Unfortunately it doesn't support go mod well, the output will write under $GOPATH/src/<MODULE_NAME>, if the --output-base option is set it generate the <OUTPUT_PATH>/<MODULE_NAME> , in a word it just can't merge the generated code in the current module under the current path we're working on (if i'm not mistaken).
- A possible way to make it correct is to prepared a code-generator container with deps installed, mount the workdir under $GOPATH/src in the conainer and run the script.
- the folder's name under <apis-package> we specifie when running the codegen script should match strictly the \<groups-versions\>, eg.
> \<apis-package\>:     github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis
>
> \<groups-versions\>:  huozj.io:v1alpha1
>
> the path should be: pkg/apis/huozj.io/v1alpha1/*.go

### code-generator update
```bash
# Prepare code-generator docker image
$ cd hack
$ docker build -t kubernetes-codegen:latest --build-arg REPO_NAME=github.com/ZhengjunHUO/k8s-custom-controller .
# Generate code in a seperate env (mounting the current project in container under ${GOPATH}/src/)
$ bash runInContainer.sh
```

### Code structure
```
├── hack			# code generator
├── kubernetes			# manifests used for testing
└── pkg
    ├── apis			# custome type definition, with runtime.Object implementation
    │				# expose its AddToScheme
    ├── client			# define clienset, informers, listers
    └── controller		# custom controller business logic
```
### Implemented by following the diagram from kubernetes/sample-controller
![diagram from kubernetes/sample-controller](https://raw.githubusercontent.com/kubernetes/sample-controller/master/docs/images/client-go-controller-interaction.jpeg)

## Custome Ressource
```bash
# Create crd & cr
$ kubectl apply -f kubernetes/crd.yaml
$ kubectl apply -f kubernetes/fufu.yaml
# Check
$ go run kubernetes/getCRD.go --kubeconfig ~/.kube/config
$ kubectl get crd
$ kubectl get fufu
```

## Custom Controller
### Quick out-cluster test
```bash
$ go run main.go --kubeconfig ~/.kube/config
# In a new terminal, create/delete a pod and watch the output
$ kubectl run alpine --image=alpine --restart=Never --command -- sleep infinity
$ kubectl delete alpine
# In a new terminal, create/update/delete a CR and watch the output
$ kubectl apply -f kubernetes/crd.yaml
$ kubectl apply -f kubernetes/fufu.yaml
$ vim kubernetes/fufu.yaml
$ kubectl apply -f kubernetes/fufu.yaml
$ kubectl delete -f kubernetes/fufu.yaml
```

### In-cluster deployment
```bash
# Create namespace
$ kubectl create ns controller
# Create serviceaccount with list/watch/get privileges on pod
$ kubectl apply -f kubernetes/rbac.yaml
# Build custom controller image and push it to your registry
$ docker build -t <IMAGE_REGISTRY/custom-controller:v1> .
$ docker push <IMAGE_REGISTRY/custom-controller:v1>
# Add credential if pull from private register
$ kubectl create -n controller secret generic regcred --from-file=.dockerconfigjson=<PATH/TO/.docker/config> --type=kubernetes.io/dockerconfigjson
# Change the image's value to your registry in deployment.yaml before apply
$ kubectl apply -f kubernetes/deployment.yaml
# Check the controller pod's output
$ kubectl logs -f custom-controller-xxx
```

## TODO: helm chart
