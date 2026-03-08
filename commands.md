`brew install kind`
`kind create cluster --config=k8s/kind.yaml --name=test1-go`
`kubectl cluster-info --context kind-test1-go`
`kubectl config get-clusters`
`kubectl config use-context test1-go`
`kubectl get nodes`
`docker build -t asoncs/test1_go .`
`docker run --name test1_go --rm -p 80:80 asoncs/test1_go`
`kubectl apply -f k8s/pod.yaml`
`kubectl get pods`
`kubectl port-forward pod/test1-go 8000:80`
`kubectl apply -f k8s/replicaset.yaml`
`kubectl get rs`
`kubectl get replicasets`
`kubectl delete replicaset test1-go`
`kubectl apply -f k8s/deployment.yaml`
`kubectl get deploy`
`kubectl get deployments`
`kubectl rollout history deployment test1-go`
// Rollback to previous version
`kubectl rollout undo deployment test1-go`
// Rollback to specific version
`kubectl rollout undo deployment test1-go --to-revision=1`
// Check the status of the rollout
`kubectl rollout status deployment test1-go`
// Stop and delete cluster
`kind delete cluster --name=test1-go`
