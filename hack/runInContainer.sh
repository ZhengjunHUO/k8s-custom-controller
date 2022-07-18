#!/bin/bash -e

REPO_DIR="/home/huo/Github/mine/k8s-custom-controller"
PROJECT_MODULE="github.com/ZhengjunHUO/k8s-custom-controller"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="huozj.io"
CUSTOM_RESOURCE_VERSION="v1alpha1"

echo "Generating client codes..."
docker run --rm -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}:Z" "${IMAGE_NAME}" ./generate-groups.sh all \
	$PROJECT_MODULE/pkg/client \
	$PROJECT_MODULE/pkg/apis \
	$CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION \
	--go-header-file ./hack/boilerplate.go.txt
