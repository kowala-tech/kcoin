package cluster

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *cluster) addKeys() error {
	log.Println("Adding keys configmaps")
	configMaps := client.Clientset.CoreV1().ConfigMaps(Namespace)

	// Remove existing keys
	err := configMaps.DeleteCollection(nil, metav1.ListOptions{
		LabelSelector: "type = key",
	})
	if err != nil {
		return err
	}

	box := packr.NewBox("./keys")
	keys := box.List()
	for _, key := range keys {
		idx := strings.LastIndex(key, "-")
		address := key[idx+1:]

		content, err := box.MustString(key)
		if err != nil {
			return err
		}

		_, err = configMaps.Create(&apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("key-%v", address),
				Labels: map[string]string{
					"network": "testnet",
					"type":    "key",
				},
			},
			Data: map[string]string{
				key: content,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (client *cluster) addKeysPassword() error {
	log.Println("Adding password configmap")
	configMaps := client.Clientset.CoreV1().ConfigMaps(Namespace)
	// Remove existing password
	err := configMaps.DeleteCollection(nil, metav1.ListOptions{
		LabelSelector: "type = password",
	})
	if err != nil {
		return err
	}

	_, err = configMaps.Create(&apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "password",
			Labels: map[string]string{
				"network": "testnet",
				"type":    "password",
			},
		},
		Data: map[string]string{
			"password.txt": "test",
		},
	})
	return err
}

func usePasswordFromConfigmap(spec *apiv1.PodSpec) {
	volume := apiv1.Volume{
		Name: "password-v",
		VolumeSource: apiv1.VolumeSource{
			ConfigMap: &apiv1.ConfigMapVolumeSource{
				LocalObjectReference: apiv1.LocalObjectReference{
					Name: "password",
				},
			},
		},
	}

	volumeMount := apiv1.VolumeMount{
		Name:      "password-v",
		MountPath: filepath.Join("/kusd", "password.txt"),
		SubPath:   "password.txt",
	}
	addVolume(spec, volume)
	addVolumeMount(&spec.Containers[0], volumeMount)
}

func useKeyFromConfigmap(spec *apiv1.PodSpec, pub_key, pub_key_file string) {
	name := fmt.Sprintf("key-%v", pub_key)

	volume := apiv1.Volume{
		Name: name,
		VolumeSource: apiv1.VolumeSource{
			ConfigMap: &apiv1.ConfigMapVolumeSource{
				LocalObjectReference: apiv1.LocalObjectReference{
					Name: name,
				},
			},
		},
	}

	volumeMount := apiv1.VolumeMount{
		Name:      name,
		MountPath: filepath.Join("/root", ".kUSD", "keystore", pub_key_file),
		SubPath:   pub_key_file,
	}
	addVolume(spec, volume)
	addVolumeMount(&spec.Containers[0], volumeMount)
}
