package main

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/yaml"
)

func main() {
	pod := corev1.Pod{}
	pod.Name = "example"
	pod.Spec.TerminationGracePeriodSeconds = ptr.To(int64(30))

	out, err := yaml.Marshal(pod)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
