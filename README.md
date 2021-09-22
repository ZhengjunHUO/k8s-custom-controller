# k8s-custom-controller
Put theory into practice, build a custom controller myself in order to know better about Kubernetes' mechanism.

## About kubernetes/code-generator:
- Unfortunately it doesn't support go mod well, the output will write under $GOPATH/src/<MODULE_NAME>, if the --output-base option is set it generate the <OUTPUT_PATH>/<MODULE_NAME> , in a word it just can't merge the generated code in the current module under the current path we're working on (if i'm not mistaken).
- A possible way to make it correct is to prepared a code-generator container with deps installed, mount the workdir under $GOPATH/src in the conainer and run the script.
- the folder's name under <apis-package> we specifie when running the codegen script should match strictly the \<groups-versions\>, eg.
> \<apis-package\>:     github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis
>
> \<groups-versions\>:  huozj.io:v1alpha1
>
> the path should be: pkg/apis/huozj.io/v1alpha1/*.go

## Quick out-cluster test

```bash
$ go run main.go --kubeconfig ~/.kube/config
```

## in-cluster deployment

```bash
# Create namespace
$ kubectl create ns controller
# Add credential if pull from private register
$ kubectl create -n controller secret generic regcred --from-file=.dockerconfigjson=<PATH/TO/.docker/config> --type=kubernetes.io/dockerconfigjson
# Create serviceaccount with list/watch/get privileges on pod
$ kubectl apply -f kubernetes/rbac.yaml
# Change the image's value to your registry in deployment.yaml
$ kubectl apply -f kubernetes/deployment.yaml
# Check the controller pod's output
$ kubectl logs -f custom-controller-xxx
```

## TODO: helm chart
