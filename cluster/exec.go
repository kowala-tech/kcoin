package cluster

import (
	"bytes"
	"log"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

type ExecResponse struct {
	StdOut string
	StdErr string
}

func (client *cluster) Exec(podName, command string) (*ExecResponse, error) {
	executor, err := client.getExecExecutor(podName, command)
	if err != nil {
		return nil, err
	}

	var (
		stdOut bytes.Buffer
		stdErr bytes.Buffer
	)

	err = executor.Stream(remotecommand.StreamOptions{
		Stdout: &stdOut,
		Stderr: &stdErr,
	})
	if err != nil {
		return nil, err
	}

	return &ExecResponse{
		StdOut: stdOut.String(),
		StdErr: stdErr.String(),
	}, nil
}

func (client *cluster) getExecExecutor(podName, command string) (remotecommand.Executor, error) {
	log.Printf("Executing command in pod named `%v`\n", podName)
	req := client.Clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(Namespace).
		SubResource("exec")

	req.VersionedParams(&apiv1.PodExecOptions{
		Container: podName,
		Command: []string{
			`./kcoin`,
			`attach`,
			`--exec`,
			command,
		},
		Stdout: true,
		Stderr: true,
	}, scheme.ParameterCodec)

	config, err := client.Backend.RestConfig()
	if err != nil {
		return nil, err
	}

	return remotecommand.NewSPDYExecutor(config, "POST", req.URL())
}
