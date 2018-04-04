package cluster

import apiv1 "k8s.io/api/core/v1"

func addVolume(spec *apiv1.PodSpec, volume apiv1.Volume) {
	if spec.Volumes == nil {
		spec.Volumes = []apiv1.Volume{volume}
	} else {
		spec.Volumes = append(spec.Volumes, volume)
	}
}

func addVolumeMount(spec *apiv1.Container, volumeMount apiv1.VolumeMount) {
	if spec.VolumeMounts == nil {
		spec.VolumeMounts = []apiv1.VolumeMount{volumeMount}
	} else {
		spec.VolumeMounts = append(spec.VolumeMounts, volumeMount)
	}
}
