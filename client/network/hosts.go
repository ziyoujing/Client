package network

// Interface to edit the host file
import (
	"github.com/lextoumbourou/goodhosts"
)

func AddHost(destIp string, domainName string) {
	hosts, _ := goodhosts.NewHosts()
	if hosts.IsWritable() {
		if !hosts.Has(destIp, domainName) {
			hosts.Add(destIp, domainName)
			hosts.Flush()
		}
	}
}

func RemoveHost(destIp string, domainName string) {
	hosts, _ := goodhosts.NewHosts()
	if hosts.IsWritable() {
		if hosts.Has(destIp, domainName) {
			hosts.Remove(destIp, domainName)
			hosts.Flush()
		}
	}
}
