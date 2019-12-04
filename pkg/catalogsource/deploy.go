package catalogsource

import (
	"fmt"
	"github.com/cardil/operator-e2e-harness/pkg/config"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Deploy will deploy catalog source to configured cluster
func Deploy() {
	catalogsourcePath := path.Join(rundir, "catalogsource-e2e.yaml")
	source := Generate()
	err := ioutil.WriteFile(catalogsourcePath, []byte(source), 0644)
	if err != nil {
		panic(err)
	}
	command := fmt.Sprintf("oc apply -n %s -f %s",
		config.OperatorsNamespace, catalogsourcePath)
	log.Print("Deploying catalog source...")
	execute(command)
}

// Undeploy will undeploy a catalog source from configured cluster
func Undeploy() {
	catalogsourcePath := path.Join(rundir, "catalogsource-e2e.yaml")
	command := fmt.Sprintf("oc delete -n %s -f %s",
		config.OperatorsNamespace, catalogsourcePath)
	log.Print("Undeploying catalog source...")
	execute(command)
	err := os.RemoveAll(rundir)
	if err != nil {
		panic(err)
	}
}
