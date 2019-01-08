#! /bin/bash

# e2e tests basic for operator
KIND_KUBECONFIG=`kind get kubeconfig-path`

kubectl get pods --kubeconfig="$KIND_KUBECONFIG"
echo " -- cluster info --"
kubectl cluster-info -v5 --kubeconfig="$KIND_KUBECONFIG"
echo 

echo "deploying operator"
kubectl apply -f deployments/registries-operator-full.yaml --kubeconfig="$KIND_KUBECONFIG"
echo "--> reg-operator deployed"


