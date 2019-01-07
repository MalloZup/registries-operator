#############################################################
# End To End Tests k8s operator
#############################################################

e2e-tests:
	cp `kind get kubeconfig-path --name="1"` ~/.kube/config
	cat  ~/.kube/config
	echo "-- cluster info --"
	kubectl cluster-info -v5
	kubectl config view
	echo
	echo "deploying operator"
	kubectl apply -f deployments/registries-operator-full.yaml
