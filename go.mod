module github.com/virtual-kubelet/cri

go 1.12

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.0
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/golang/groupcache v0.0.0-20190129154638-5b532d6fd5ef // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/virtual-kubelet/node-cli v0.1.2
	github.com/virtual-kubelet/virtual-kubelet v1.1.0
	go.opencensus.io v0.22.0
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/grpc v1.23.1
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0-20190615170205-3722cb685593 // indirect
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/cri-api v0.0.0-20191114225051-381f557272b4
	k8s.io/klog v0.3.3
	k8s.io/kubernetes v1.15.2
	k8s.io/utils v0.0.0-20190607212802-c55fbcfc754a // indirect
)

replace k8s.io/api => k8s.io/api v0.0.0-20190606204050-af9c91bd2759

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d

replace k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190606204521-b8faab9c5193+incompatible

replace k8s.io/kubernetes => k8s.io/kubernetes v1.14.3

replace go.opencensus.io => go.opencensus.io v0.19.3
