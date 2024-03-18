#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

. ../hack/tools/vendor/k8s.io/code-generator/kube_codegen.sh


export KUBE_VERBOSE=3
SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/metal3.io/v1alpha1/generated" \
    --output-pkg "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1/generated" \
    --boilerplate "${SCRIPT_ROOT}/../hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}"
