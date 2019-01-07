#############################################################
# End To End Tests k8s operator
#############################################################

KUBECONFIG := $(shell kind get kubeconfig-path || echo "unknown")
e2e-tests:
	export KUBECONFIG=$(KUBECONFIG)
	echo "$$KUBECONFIG"
	kubectl get pods --kubeconfig="$(KUBECONFIG)"
	echo " -- cluster info --"
	kubectl cluster-info -v5
	kubectl config view
	echo
	echo "deploying operator"
	kubectl apply -f deployments/registries-operator-full.yaml
