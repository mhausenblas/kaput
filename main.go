package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/mhausenblas/kubecuddler"
)

var (
	releaseVersion string
	kubectlbin     string
	verbose        bool
)

func main() {
	// if we have an argument, we interpret it as the namespace:
	if len(os.Args) > 1 {
		if os.Args[1] == "version" {
			fmt.Printf("This is the Kubernetes attack & policy underminer tool in version %v\n", releaseVersion)
			os.Exit(0)
		}
	}
	// get params and env variables:
	flag.Parse()
	if kb := os.Getenv("KAPUT_KUBECTL_BIN"); kb != "" {
		kubectlbin = kb
	}
	if v := os.Getenv("KAPUT_VERBOSE"); v != "" {
		verbose = true
	}
	if err := gather(); err != nil {
		log(err)
	}
	// attack()
}

// gather runs a set of Kubernetes security scanners and
// dumps the data to /tmp/kaput_UNIXNANOTIMENOW
func gather() (err error) {
	res, err := kubecuddler.Kubectl(verbose, verbose, kubectlbin,
		"get", "all", "--all-namespaces", "--output=json")
	if err != nil {
		return
	}
	resfile := fmt.Sprintf("kaput_%v", time.Now().UnixNano())
	fn := filepath.Join("/tmp", resfile)
	err = ioutil.WriteFile(fn, []byte(res), 0644)
	if err != nil {
		return
	}
	return nil
}

func log(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
}

func info(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[92m%v\x1b[0m\n", msg)
}
