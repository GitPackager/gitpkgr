package main

import (
	"flag"
	"fmt"
	"github.com/GitPackager/gitpkgr/internals/net"
)

var (
	sync   bool
	update bool
)

func init() {
	flag.BoolVar(&sync, "sync", false, "Sync all packages")
	flag.BoolVar(&update, "update", false, "Update all packages")
	flag.Parse()
}

func main() {
	res, err := net.GetRepoHash("acidicneko/karui", "main")
	if err != nil {
		fmt.Println("Error fetching repository hash:", err)
	} else {
		fmt.Println("Repository Hash:", res)
	}
	if sync {
		fmt.Println("Syncing all packages...")
	}

	if update {
		fmt.Println("Updating all packages...")
	}

}
