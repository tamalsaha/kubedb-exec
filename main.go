package main

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/external-dns/source"
)

func main() {
	fmt.Println("Using Generated client")
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	kc, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}

	nodes, err := kc.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodes.Items {
		fmt.Println(node.GetName())
	}
	// os.Exit(1)

	ctx := context.Background()
	namespace := "default"
	annotationFilter := ""
	fqdnTemplate := "xyz.appscode.ninja"
	combineFqdnAnnotation := false
	ignoreHostnameAnnotation := false
	ignoreIngressTLSSpec := true // for nats
	ignoreIngressRulesSpec := false
	labelSelector := labels.SelectorFromSet(map[string]string{})

	s1, err := source.NewIngressSource(
		ctx,
		kc,
		namespace,
		annotationFilter,
		fqdnTemplate,
		combineFqdnAnnotation,
		ignoreHostnameAnnotation,
		ignoreIngressTLSSpec,
		ignoreIngressRulesSpec,
		labelSelector,
	)
	if err != nil {
		panic(err)
	}
	eps, err := s1.Endpoints(ctx)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(eps, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
