package mynmap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func Nmap() {
	target := "192.168.1.8/24"

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts("80, 443"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatal("Error : ", err)
	}

	results, warnings, err := scanner.Run()
	if err != nil {
		log.Fatal("Error : ", err)
	}

	if warnings != nil {
		log.Fatalf("Warnings %s: ", warnings)
	}

	if len(results.Hosts) > 0 {
		for _, host := range results.Hosts {
			if len(host.Ports) == 0 || len(host.Addresses) == 0 {
				continue
			}
			fmt.Printf("IP: %q\n", host.Addresses[0])
			if len(host.Addresses) > 1 {
				fmt.Printf("\tMAC %v\n", host.Addresses[1])
			}

			for _, port := range host.Ports {
				fmt.Printf("\t Port %d %s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			}
		}
	}

}
