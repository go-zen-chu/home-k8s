# home-k8s

Kubernetes kustomization manifests run on my home k8s.

## Prerequisite

Before this repo's manifest automatically synced, we need to do things below.

- ArgoCD deployed on k8s
- Setup vault

```bash
kubectl apply -f ./base/vault/vault-server.yaml
# need to change variable according to your environment
kubectl apply -f ./base/vault/vault-persistent-volume.yaml
kubectl apply -f ./base/argocd/argocd-vault-plugin-secret.yaml
# also set the same secret value to deployed vault so argocd-vault-plugin can complement the values
vault write /home-k8s/vault/persistent-volume-nfs-ip 192.168.x.x
vault write /home-k8s/vault/persistent-volume-nfs-volume-path /volume1/xxx/xxx
vault write /home-k8s/argocd-vault-plugin/vault-address 192.168.x.x 
vault write /home-k8s/argocd-vault-plugin/vault-token xxxxxx
```

## Setup ArgoCD and sync

```bash
# you can access to UI via port-forwarding. https://argo-cd.readthedocs.io/en/stable/getting_started/
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

Access to UI and set this repo and done :tada:!
