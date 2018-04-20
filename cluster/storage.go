package cluster

import (
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	NetworkIDKey = "network-id"
)

// StoreString stores a configuration value in the cluster
func (client *cluster) StoreString(key, value string) error {
	log.Printf("Storing config string with key `%v`\n", key)
	config_maps := client.Clientset.CoreV1().ConfigMaps(client.Namespace)

	current, err := config_maps.Get("config", metav1.GetOptions{})
	if err != nil {
		log.Println("Creating config store")
		// Assuming the configmap doesn't exist
		_, err = config_maps.Create(&apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: "config",
				Labels: map[string]string{
					"network": "testnet",
					"type":    "config",
				},
			},
			Data: map[string]string{
				key: value,
			},
		})
		if err != nil {
			return err
		}
		return nil
	}

	current.Data[key] = value
	_, err = config_maps.Update(current)
	return err
}

// GetString retrieves a configuration value from the cluster
func (client *cluster) GetString(key string) (string, error) {
	config_maps := client.Clientset.CoreV1().ConfigMaps(client.Namespace)

	current, err := config_maps.Get("config", metav1.GetOptions{})
	if err != nil {
		return "", err
	}
	return current.Data[key], nil
}
