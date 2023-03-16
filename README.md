# home-k8s

Kubernetes kustomization manifests run on my home k8s.

## Prerequisite

Before this repo's manifest automatically synced, we need to do things below.

- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/) deployed on k8s
- Setup vault

### Setup vault

Private configuration is required to generate vault manifst, so we execute secret replace tool to generate manifest.

```bash
kubectl apply -f ./base/vault/vault-server.yaml
# need to change variable according to your environment
cp sample.envrc .envrc
vim .envrc
source .envrc
go run tools/cmd/replace.go ./base/vault/vault-persistent-volume.yaml | kubectl apply --dry-run='server' -f -
go run tools/cmd/replace.go ./base/argocd/argocd-vault-plugin-secret.yaml | kubectl apply --dry-run='server' -f -
# also set the same secret value to deployed vault so argocd-vault-plugin can complement the values
vault write /home-k8s/vault/persistent-volume-nfs-ip 192.168.x.x
vault write /home-k8s/vault/persistent-volume-nfs-volume-path /volume1/xxx/xxx
vault write /home-k8s/argocd-vault-plugin/vault-address 192.168.x.x 
vault write /home-k8s/argocd-vault-plugin/vault-token xxxxxx
```

After a vault started on k8s and argocd-vault-plugin is setup, private data are read from vault.

## Setup ArgoCD and sync

```bash
# you can access to UI via port-forwarding. https://argo-cd.readthedocs.io/en/stable/getting_started/
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

Access to UI and set this repo and done :tada:!
