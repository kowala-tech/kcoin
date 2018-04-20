package cluster

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/pkg/errors"
	"path/filepath"
	"strconv"
)

type KcoinPodBuilder interface {
	Build() (*apiv1.Pod, error)
	WithNetworkId(string) KcoinPodBuilder
	WithPort(int32) KcoinPodBuilder
	WithName(string) KcoinPodBuilder
	WithBootnode(string) KcoinPodBuilder
	WithSyncMode(string) KcoinPodBuilder
	WithLogLevel(int) KcoinPodBuilder
}

var availablePort int32 = 30301

type Builder struct {
	network   string
	port      int32
	name      string
	namespace string
	bootnode  string
	syncMode  string
	logLevel  int
}

func NewPodBuilder() *Builder {
	return &Builder{
		syncMode: "fast",
		logLevel: 3,
	}
}

func (builder *Builder) WithNetworkId(name string) KcoinPodBuilder {
	builder.network = name
	return builder
}

func (builder *Builder) WithPort(port int32) KcoinPodBuilder {
	builder.port = port
	return builder
}

func (builder *Builder) WithName(name string) KcoinPodBuilder {
	builder.name = name
	return builder
}

func (builder *Builder) WithBootnode(address string) KcoinPodBuilder {
	builder.bootnode = address
	return builder
}

func (builder *Builder) WithSyncMode(mode string) KcoinPodBuilder {
	builder.syncMode = mode
	return builder
}

func (builder *Builder) WithLogLevel(level int) KcoinPodBuilder {
	builder.logLevel = level
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
		"--syncmode", builder.syncMode,
		"--bootnodes", builder.bootnode,
		"--networkid", builder.network,
		"--verbosity", strconv.Itoa(builder.logLevel),
	}
	return &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: builder.name,
			Labels: map[string]string{
				"network": "testnet",
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
