# KubeCon 2022 Dell Booth Talk Code & Notes

## Prerequisites
- [go](https://golang.org/dl/) version v1.19.0+
- [docker](https://docs.docker.com/install/) version 17.03+.
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster. (I use [kind](https://kind.sigs.k8s.io/))

## Kind Cluster Setup
1. Install kind via your chosen method - https://kind.sigs.k8s.io/docs/user/quick-start/#installation
2. Use the local registry helper script (https://kind.sigs.k8s.io/docs/user/local-registry/)
3. Run `./kind-with-registry.sh`

## Kubebuilder 
1. Install Kubebuilder
   `curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)`

   `chmod +x kubebuilder && mv kubebuilder /usr/local/bin/`
2. (Optional) Clear the controller directory so you can start from scratch with `rm -rf controller/*`
3. `cd controller`
4. `kubebuilder init --domain kubecon.dell.com --repo github.com/king-jam/kubecon2022/controller`
5. Check out what was built with `tree`.
6. Create your first API `kubebuilder create api --group dell --version v1 --kind Server`
7. Open an IDE to start creating your code.
   1. Go to `api/v1/server_types.go` and make edits.
   2. Go to `controllers/server_controller.go` and make edits.
   3. Run `make`
   4. Run `make install`
8. (Optional) Edit the Makefile to change the IMG to a `localhost:5001/controller:latest`.
9.  `make docker-build docker-push`
10. Check cluster pods before deployment with `kubectl get pods -A`
11. `make deploy`
12. Check cluster pods after deployment with `kubectl get pods -A`
13. 