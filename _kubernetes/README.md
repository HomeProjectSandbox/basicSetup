create cluster with kind

`go install sigs.k8s.io/kind@v0.26.0`

set kubeconfig (not mess the default one)

manual
`kind create cluster --name kind-test0`
`kind delete cluster --name kind-test0`

load image to the cluster:
`kind load docker-image my-custom-image-0 my-custom-image-1`


kind create cluster --config=kind/clusterconfig.yaml
`cd /path/to/your/helmchart`
`helm install my-app . --namespace app --debug --dry-run`

`helm install <release-name> . --namespace <namespace> --create-namespace`
`helm install myapp . --namespace app --create-namespace`
```
helm list --namespace <namespace>
kubectl get all --namespace <namespace>
```