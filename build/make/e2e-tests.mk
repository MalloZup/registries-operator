#############################################################
# End To End Tests k8s operator
#############################################################

KIND_KUBECONFIG := $(shell kind get kubeconfig-path)
e2e-tests:
	kubectl get pods --kubeconfig="$(KIND_KUBECONFIG)"
	echo " -- cluster info --"
	kubectl cluster-info -v5 --kubeconfig="$(KIND_KUBECONFIG)"
	echo "deploying operator"
	kubectl apply -f deployments/registries-operator-full.yaml --kubeconfig="$(KIND_KUBECONFIG)"
