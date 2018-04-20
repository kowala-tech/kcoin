package cluster

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/pkg/errors"
	"path/filepath"
)

type PodBuilder interface {
	Build() (*apiv1.Pod, error)
	Network(string) PodBuilder
	Port(int32) PodBuilder
	Name(string) PodBuilder
	Bootnode(string) PodBuilder
}

var availablePort int32 = 31301

type Builder struct {
	network   string
	port      int32
	name      string
	namespace string
	bootnode  string
}

func NewPodBuilder() *Builder {
	return &Builder{}
}

func (builder *Builder) Network(name string) PodBuilder {
	builder.network = name
	return builder
}

func (builder *Builder) Port(port int32) PodBuilder {
	builder.port = port
	return builder
}

func (builder *Builder) Name(name string) PodBuilder {
	builder.name = name
	return builder
}

func (builder *Builder) Bootnode(address string) PodBuilder {
	builder.bootnode = address
	return builder
}

func (builder *Builder) Build() (*apiv1.Pod, error) {
	if builder.name == "" {
		return nil, errors.New("cant build pod without name")
	}
	if builder.bootnode == "" {
		return nil, errors.New("cant build pod without bootnode address")
	}

	// no port provided, get next available
	if builder.port == 0 {
		builder.port = availablePort
		availablePort += 1
	}

	return builder.build(), nil
}

func (builder *Builder) build() *apiv1.Pod {
	args := []string{
		"--syncmode", "fast",
		"--bootnodes", builder.bootnode,
		"--networkid", builder.network,
		"--verbosity", "6",
	}
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: builder.name,
			Labels: map[string]string{
				"network": builder.network,
				"app":     "node",
				"name":    builder.name,
			},
		},
		Spec: apiv1.PodSpec{
			Volumes: []apiv1.Volume{volume()},
			Containers: []apiv1.Container{
				{
					Name:            builder.name,
					Image:           "kowalatech/kusd:dev",
					ImagePullPolicy: apiv1.PullNever,
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: builder.port,
							Protocol:      apiv1.ProtocolTCP,
						},
					},
					Args:         args,
					VolumeMounts: []apiv1.VolumeMount{volumeMount()},
				},
			},
		},
	}
}

func volume() apiv1.Volume {
	return apiv1.Volume{
		Name: "genesis-v",
		VolumeSource: apiv1.VolumeSource{
			ConfigMap: &apiv1.ConfigMapVolumeSource{
				LocalObjectReference: apiv1.LocalObjectReference{
					Name: "genesis",
				},
			},
		},
	}
}

func volumeMount() apiv1.VolumeMount {
	return apiv1.VolumeMount{
		Name:      "genesis-v",
		MountPath: filepath.Join("/kcoin", "genesis.json"),
		SubPath:   "genesis.json",
	}
}
