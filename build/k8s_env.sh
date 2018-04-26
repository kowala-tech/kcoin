#!/bin/sh

set -e

if [[ -z "$K8S_CLUSTER_ADDRESS" ]]; then
	echo "Kubernetes cluster not configured. Using minikube."

    # Export minikube docker environment
    eval "$(minikube docker-env -p testing --shell bash | grep DOCKER | sed "s/export\ /export\ K8S_/g")"

    # Export minikube k8s config
    export K8S_CLUSTER_IP="`minikube ip -p testing`"
    export K8S_CLUSTER_MASTER_URL=$(kubectl config view | grep server | cut -f 2- -d ":" | tr -d " " | grep $K8S_CLUSTER_IP)
    export K8S_CLUSTER_TOKEN=$(kubectl describe --cluster testing secret $(kubectl get --cluster testing secrets | grep default | cut -f1 -d ' ') | grep -E '^token' | cut -f2 -d':' | tr -d '\ ' | tr -d '\t')
fi

# Launch the arguments with the configured environment.
exec ./build/env.sh "$@"
