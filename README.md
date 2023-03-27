Multiple ways to create a Kubernetes controller:

* k8s http api
* k8s go api
* k8s client-go informer
* k8s client-go shared informer
* k8s go-generator for custom CRD

Note: If you intend to generate code then you will need the code-generator repo to exist in an old-style location.
One easy way to do this is to use the command ```go mod vendor``` to create and populate the vendor directory.
