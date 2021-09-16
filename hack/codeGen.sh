bash generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/ZhengjunHUO/k8s-custom-controller/pkg/client \
  github.com/ZhengjunHUO/k8s-custom-controller/pkg/apis \
  huozj.io:v1alpha1 \
  --go-header-file boilerplate.go.txt
