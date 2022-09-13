本周作业
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
1. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
1. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
1. 当访问 `localhost/healthz` 时，应返回 200

# kubeadm init

```shell
$ kubeadm init \
 --image-repository registry.aliyuncs.com/google_containers \
 --pod-network-cidr=192.168.0.0/16 \
 --apiserver-advertise-address=192.168.172.129
```

# Copy kubeconfig

```shell
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

# Install calico cni plugin

https://docs.projectcalico.org/getting-started/kubernetes/quickstart

```shell
$ kubectl create -f https://docs.projectcalico.org/manifests/tigera-operator.yaml
$ kubectl create -f https://docs.projectcalico.org/manifests/custom-resources.yaml
```

# install ingress
```shell
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
```

# install cert-manager
```shell
helm repo add jetstack https://charts.jetstack.io
helm repo updatekubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.crds.yaml
helm install \ 
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \  
  --version v1.7.1
```
