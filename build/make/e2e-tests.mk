#############################################################
# End To End Tests k8s operator
#############################################################

e2e-tests:
	echo "deploying operator"
	kubectl apply -f deployments/registries-operator-full.yaml
