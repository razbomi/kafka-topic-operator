# kafka-topic-operator
Manage kafka topics and connector configuration along with your deployments.

# Developing

Run `kind create cluster --name local`
Run `kubectl cluster-info --context kind-local`
Run `make run`

Using the [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)
Framework [kubebuilder](https://book.kubebuilder.io/quick-start.html)
Runtime [kind](https://github.com/kubernetes-sigs/kind)