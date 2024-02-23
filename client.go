package cri

import (
	"fmt"

	"github.com/virtual-kubelet/node-cli/manager"
	"github.com/virtual-kubelet/virtual-kubelet/errdefs"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	"golang.org/x/net/context"
	v1 "k8s.io/api/core/v1"
	criapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

// Call RunPodSandbox on the CRI client
func runPodSandbox(ctx context.Context, client criapi.RuntimeServiceClient, config *criapi.PodSandboxConfig) (string, error) {
	ctx, span := trace.StartSpan(ctx, "cri.getPodSandboxes")
	defer span.End()

	request := &criapi.RunPodSandboxRequest{Config: config}
	log.G(ctx).Debug("RunPodSandboxRequest")

	r, err := client.RunPodSandbox(context.Background(), request)
	log.G(ctx).Debug("RunPodSandboxResponse")
	if err != nil {
		span.SetStatus(err)
		return "", err
	}
	log.G(ctx).Debug("New pod sandbox created")
	return r.PodSandboxId, nil
}

// Call StopPodSandbox on the CRI client
func stopPodSandbox(ctx context.Context, client criapi.RuntimeServiceClient, id string) error {
	ctx, span := trace.StartSpan(ctx, "cri.getPodSandboxes")
	defer span.End()
	if id == "" {
		err := errdefs.InvalidInput("ID cannot be empty")
		span.SetStatus(err)
		return err
	}
	request := &criapi.StopPodSandboxRequest{PodSandboxId: id}
	log.G(ctx).Debug("StopPodSandboxRequest")
	_, err := client.StopPodSandbox(context.Background(), request)
	log.G(ctx).Debug("StopPodSandboxResponse")
	if err != nil {
		span.SetStatus(err)
		return err
	}

	log.G(ctx).Debugf("Stopped sandbox %s", id)
	return nil
}

// Call RemovePodSandbox on the CRI client
func removePodSandbox(ctx context.Context, client criapi.RuntimeServiceClient, id string) error {
	ctx, span := trace.StartSpan(ctx, "cri.getPodSandboxes")
	defer span.End()
	if id == "" {
		err := errdefs.InvalidInput("ID cannot be empty")
		span.SetStatus(err)
		return err
	}
	request := &criapi.RemovePodSandboxRequest{PodSandboxId: id}
	log.G(ctx).Debug("RemovePodSandboxRequest")
	_, err := client.RemovePodSandbox(context.Background(), request)
	log.G(ctx).Debug("RemovePodSandboxResponse")
	if err != nil {
		span.SetStatus(err)
		return err
	}
	log.G(ctx).Debugf("Removed sandbox %s", id)
	return nil
}

// Call ListPodSandbox on the CRI client
func getPodSandboxes(ctx context.Context, client criapi.RuntimeServiceClient) ([]*criapi.PodSandbox, error) {
	ctx, span := trace.StartSpan(ctx, "cri.getPodSandboxes")
	defer span.End()

	filter := &criapi.PodSandboxFilter{}
	request := &criapi.ListPodSandboxRequest{
		Filter: filter,
	}

	log.G(ctx).Debug("ListPodSandboxRequest")
	r, err := client.ListPodSandbox(context.Background(), request)

	log.G(ctx).Debug("ListPodSandboxResponse")
	if err != nil {
		span.SetStatus(err)
		return nil, err
	}
	return r.GetItems(), err
}

// Call PodSandboxStatus on the CRI client
func getPodSandboxStatus(ctx context.Context, client criapi.RuntimeServiceClient, psId string) (*criapi.PodSandboxStatus, error) {
	ctx, span := trace.StartSpan(ctx, "cri.getPodSandboxStatus")
	defer span.End()

	if psId == "" {
		err := errdefs.InvalidInput("Pod ID cannot be empty in GPSS")
		span.SetStatus(err)
		return nil, err
	}

	request := &criapi.PodSandboxStatusRequest{
		PodSandboxId: psId,
		Verbose:      false,
	}

	log.G(ctx).Debug("PodSandboxStatusRequest")
	r, err := client.PodSandboxStatus(context.Background(), request)
	log.G(ctx).Debug("PodSandboxStatusResponse")
	if err != nil {
		span.SetStatus(err)
		return nil, err
	}

	return r.Status, nil
}

// Call CreateContainer on the CRI client
func createContainer(ctx context.Context, client criapi.RuntimeServiceClient, config *criapi.ContainerConfig, podConfig *criapi.PodSandboxConfig, pId string) (string, error) {
	ctx, span := trace.StartSpan(ctx, "cri.createContainer")
	defer span.End()

	request := &criapi.CreateContainerRequest{
		PodSandboxId:  pId,
		Config:        config,
		SandboxConfig: podConfig,
	}
	log.G(ctx).Debug("CreateContainerRequest")
	r, err := client.CreateContainer(context.Background(), request)
	log.G(ctx).Debug("CreateContainerResponse")
	if err != nil {
		span.SetStatus(err)
		return "", err
	}
	log.G(ctx).Debugf("Container created: %s", r.ContainerId)
	return r.ContainerId, nil
}

// Call StartContainer on the CRI client
func startContainer(ctx context.Context, client criapi.RuntimeServiceClient, cId string) error {
	ctx, span := trace.StartSpan(ctx, "cri.startContainer")
	defer span.End()

	if cId == "" {
		err := errdefs.InvalidInput("ID cannot be empty")
		span.SetStatus(err)
		return err
	}
	request := &criapi.StartContainerRequest{
		ContainerId: cId,
	}
	log.G(ctx).Debug("StartContainerRequestv")
	_, err := client.StartContainer(context.Background(), request)
	log.G(ctx).Debug("StartContainerResponse")
	if err != nil {
		span.SetStatus(err)
		return err
	}
	log.G(ctx).Debugf("Container started: %s", cId)
	return nil
}

// Call ContainerStatus on the CRI client
func getContainerCRIStatus(ctx context.Context, client criapi.RuntimeServiceClient, cId string) (*criapi.ContainerStatus, error) {
	ctx, span := trace.StartSpan(ctx, "cri.getContainerCRIStatus")
	defer span.End()

	if cId == "" {
		err := errdefs.InvalidInput("Container ID cannot be empty in GCCS")
		span.SetStatus(err)
		return nil, err
	}

	request := &criapi.ContainerStatusRequest{
		ContainerId: cId,
		Verbose:     false,
	}
	log.G(ctx).Debug("ContainerStatusRequest")
	r, err := client.ContainerStatus(context.Background(), request)
	log.G(ctx).Debug("ContainerStatusResponsev")
	if err != nil {
		span.SetStatus(err)
		return nil, err
	}

	return r.Status, nil
}

// Call ListContainers on the CRI client
func getContainersForSandbox(ctx context.Context, client criapi.RuntimeServiceClient, psId string) ([]*criapi.Container, error) {
	ctx, span := trace.StartSpan(ctx, "cri.getContainersForSandbox")
	defer span.End()

	filter := &criapi.ContainerFilter{}
	filter.PodSandboxId = psId
	request := &criapi.ListContainersRequest{
		Filter: filter,
	}
	log.G(ctx).Debug("ListContainerRequest")
	r, err := client.ListContainers(context.Background(), request)
	log.G(ctx).Debug("ListContainerResponse")
	if err != nil {
		span.SetStatus(err)
		return nil, err
	}
	return r.Containers, nil
}

// Pull and image on the CRI client and return the image ref
func pullImage(ctx context.Context, client criapi.ImageServiceClient, image string) (string, error) {
	ctx, span := trace.StartSpan(ctx, "cri.pullImage")
	defer span.End()

	request := &criapi.PullImageRequest{
		Image: &criapi.ImageSpec{
			Image: image,
		},
	}
	log.G(ctx).Debug("PullImageRequest")
	r, err := client.PullImage(context.Background(), request)
	log.G(ctx).Debug("PullImageResponse")
	if err != nil {
		span.SetStatus(err)
		return "", err
	}

	return r.ImageRef, nil
}

// Generate the CRI ContainerConfig from the Pod and container specs
// TODO: Probably incomplete
func generateContainerConfig(ctx context.Context, container *v1.Container, pod *v1.Pod, imageRef, podVolRoot string, rm *manager.ResourceManager, attempt uint32) (*criapi.ContainerConfig, error) {
	// TODO: Probably incomplete
	config := &criapi.ContainerConfig{
		Metadata: &criapi.ContainerMetadata{
			Name:    container.Name,
			Attempt: attempt,
		},
		Image:       &criapi.ImageSpec{Image: imageRef},
		Command:     container.Command,
		Args:        container.Args,
		WorkingDir:  container.WorkingDir,
		Envs:        createCtrEnvVars(container.Env),
		Labels:      createCtrLabels(container, pod),
		Annotations: createCtrAnnotations(container, pod),
		Linux:       createCtrLinuxConfig(container, pod),
		LogPath:     fmt.Sprintf("%s-%d.log", container.Name, attempt),
		Stdin:       container.Stdin,
		StdinOnce:   container.StdinOnce,
		Tty:         container.TTY,
	}
	mounts, err := createCtrMounts(ctx, container, pod, podVolRoot, rm)
	if err != nil {
		return nil, err
	}
	config.Mounts = mounts
	return config, nil
}

// Greate CRI PodSandboxConfig from the Pod spec
// TODO: This is probably incomplete
func generatePodSandboxConfig(ctx context.Context, pod *v1.Pod, logDir string, attempt uint32) (*criapi.PodSandboxConfig, error) {
	podUID := string(pod.UID)
	config := &criapi.PodSandboxConfig{
		Metadata: &criapi.PodSandboxMetadata{
			Name:      pod.Name,
			Namespace: pod.Namespace,
			Uid:       podUID,
			Attempt:   attempt,
		},
		Labels:       createPodLabels(pod),
		Annotations:  pod.Annotations,
		LogDirectory: logDir,
		DnsConfig:    createPodDnsConfig(pod),
		Hostname:     createPodHostname(pod),
		PortMappings: createPortMappings(pod),
		Linux:        createPodSandboxLinuxConfig(pod),
	}
	return config, nil
}
