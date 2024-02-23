//go:build !darwin

package cri

const (
	CriSocketPath      = "/run/containerd/containerd.sock"
	PodVolRoot         = "/run/vk-cri/volumes/"
	PodSecretVolDir    = "/secrets"
	PodConfigMapVolDir = "/configmaps"
)
