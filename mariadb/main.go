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
		Name:      "sample-mariadb-0",
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

	// k exec -it sample-mariadb-0 -c mariadb -- bash -c 'mysql -u"$MYSQL_ROOT_USERNAME" -p"$MYSQL_ROOT_PASSWORD"'
	out, err := exec.Exec(
		cfg,
		pod,
		exec.TTY(true),
		exec.Container("mariadb"),
		// exec.Command("bash", "-c", `curl -k https://localhost:9200`),
		exec.Command("bash", "-c", `mysql -u"$MYSQL_ROOT_USERNAME" -p"$MYSQL_ROOT_PASSWORD"`),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
