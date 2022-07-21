package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"kmodules.xyz/client-go/tools/exec"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	fmt.Println("Using Generated client")
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	pod := types.NamespacedName{
		Namespace: "demo",
		Name:      "es-quickstart-0",
	}
	/*
		// k exec -it es-quickstart-0 -c elasticsearch -- curl -k https://localhost:9200
		out, err := exec.Exec(
			cfg,
			pod,
			exec.TTY(true),
			exec.Container("elasticsearch"),
			exec.Command("curl", "-k", "https://localhost:9200"),
		)
	*/
	// k exec -it es-quickstart-0 -c elasticsearch -- bash -c 'curl -XGET -k -u "$ELASTIC_USER:$ELASTIC_PASSWORD" "https://localhost:9200/_cluster/health?pretty"'
	out, err := exec.Exec(
		cfg,
		pod,
		exec.TTY(true),
		exec.Container("elasticsearch"),
		// exec.Command("bash", "-c", `curl -k https://localhost:9200`),
		exec.Command("bash", "-c", `curl -XGET -k -u "$ELASTIC_USER:$ELASTIC_PASSWORD" "https://localhost:9200/_cluster/health?pretty"`),

		// exec.Command("curl -k https://localhost:9200"),
		// exec.Command(`/bin/bash -c 'curl -XGET -k -u "$ELASTIC_USER:$ELASTIC_PASSWORD" "https://localhost:9200/_cluster/health?pretty"'`),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
