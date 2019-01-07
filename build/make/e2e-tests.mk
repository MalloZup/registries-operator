#############################################################
# End To End Tests k8s operator
#############################################################

e2e-tests:
	echo "$$KUBECONFIG"
	kubectl get pods --kubeconfig=`kind get kubeconfig-path`
	cat `kind get kubeconfig-path`
	echo "-- cluster info --"
	kubectl cluster-info -v5
	kubectl config view
	echo
	echo "deploying operator"
	kubectl apply -f deployments/registries-operator-full.yaml
