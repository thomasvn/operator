package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Parse the user's kubeconfig file
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	fmt.Printf("Kubeconfig path: %s\n", *kubeconfig)

	// Use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// DISCOVERY CLIENT
	// dclient, err := discovery.NewDiscoveryClientForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// apiresources, err := dclient.ServerPreferredNamespacedResources()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// for _, value := range apiresources {
	// 	fmt.Println(value)
	// }

	// DYNAMIC CLIENT
	// dyclient := dynamic.ConfigFor(config)
	// dynamic.ResourceInterface.Get(context.TODO())

	// Create the Kubernetes Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// DRAIN ALL NODES
		// https://pkg.go.dev/k8s.io/kubectl/pkg/drain

		// DELETE ALL PODS

		// LIST ALL PODS
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for _, value := range pods.Items {
			fmt.Println(value.Name)
		}

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		// namespace := "default"
		// pod := "example-xxxxx"
		// _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		// if errors.IsNotFound(err) {
		// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
		// 		pod, namespace, statusError.ErrStatus.Message)
		// } else if err != nil {
		// 	panic(err.Error())
		// } else {
		// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		// }

		time.Sleep(10 * time.Second)
	}
}
