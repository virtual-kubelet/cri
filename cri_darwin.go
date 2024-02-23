package cri

const (
	CriSocketPath      = "/var/run/containerd/containerd.sock"
	PodVolRoot         = "/var/run/vk-cri/volumes/"
	PodSecretVolDir    = "/var/lib/vk-cri/secrets"
	PodConfigMapVolDir = "/var/lib/vk-cri/configmaps"
)
