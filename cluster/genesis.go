package cluster

import (
	"log"
	"path/filepath"

	"encoding/json"

	"github.com/kowala-tech/kcoin/common"
	"github.com/kowala-tech/kcoin/kcoin/genesis"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (client *cluster) generateGenesis(seedAccount common.Address) error {
	log.Println("Generating and storing genesis configmap")
	configMaps := client.Clientset.CoreV1().ConfigMaps(Namespace)

	// Remove existing genesis
	err := configMaps.DeleteCollection(nil, metav1.ListOptions{
		LabelSelector: "type = genesis",
	})
	if err != nil {
		return err
	}

	newGenesis, err := genesis.GenerateGenesis(
		genesis.Options{
			Network:                        "test",
			MaxNumValidators:               "1",
			UnbondingPeriod:                "0",
			AccountAddressGenesisValidator: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
			PrefundedAccounts: []genesis.PrefundedAccount{
				{
					AccountAddress: "0xd6e579085c82329c89fca7a9f012be59028ed53f",
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
				{
					AccountAddress: seedAccount.Hex(),
					Balance:        "0x200000000000000000000000000000000000000000000000000000000000000",
				},
				{
					AccountAddress: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
					Balance:        "1000000000000000000",
				},
			},
			SmartContractsOwner: "0x259be75d96876f2ada3d202722523e9cd4dd917d",
		},
	)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(newGenesis, "", "  ")
	if err != nil {
		return err
	}

	// Add new genesis
	_, err = configMaps.Create(&apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "genesis",
			Labels: map[string]string{
				"network": "testnet",
				"type":    "genesis",
			},
		},
		Data: map[string]string{
			"genesis.json": string(out),
		},
	})
	return err
}

func useGenesisFromConfigmap(spec *apiv1.PodSpec) {
	volume := apiv1.Volume{
		Name: "genesis-v",
		VolumeSource: apiv1.VolumeSource{
			ConfigMap: &apiv1.ConfigMapVolumeSource{
				LocalObjectReference: apiv1.LocalObjectReference{
					Name: "genesis",
				},
			},
		},
	}

	volumeMount := apiv1.VolumeMount{
		Name:      "genesis-v",
		MountPath: filepath.Join("/kcoin", "genesis.json"),
		SubPath:   "genesis.json",
	}
	addVolume(spec, volume)
	addVolumeMount(&spec.Containers[0], volumeMount)
}
