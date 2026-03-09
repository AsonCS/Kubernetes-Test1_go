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
`kubectl delete pod test1-go`
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
`kubectl delete deployment test1-go`
// Stop and delete cluster
`kind delete cluster --name=test1-go`
`kubectl apply -f k8s/service.yaml`
`kubectl apply -f k8s/service.loadbalancer.yaml`
`kubectl get svc`
`kubectl get services`
`kubectl delete service test1-go-service`
`kubectl port-forward svc/test1-go-service 8000:80`
// http://localhost:8000/api/v1
// http://localhost:8000/api/v1/namespaces/default/services
// http://localhost:8000/api/v1/namespaces/default/services/test1-go-service
`kubectl proxy --port=8000`
`kubectl apply -f k8s/configmap-env.yaml`
`docker build -t asoncs/test1_go:v3 .`
`docker push asoncs/test1_go:v3`
`kubectl apply -f k8s/configmap-family.yaml`
`kubectl get configmap`
`kubectl exec -it test1-go-b5996b954-djmvg -- bash`
`kubectl exec -it test1-go-7f6ccfcbbb-24cqh -- bash`
`kubectl logs test1-go-7f6ccfcbbb-24cqh`
`kubectl apply -f k8s/secret.yaml`
`kubectl get secret`

