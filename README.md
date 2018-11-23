# kaput

`kaput` is the Kubernetes Attack and Policy Underminer Tool— 
think of it as a toolbox exploiting known weaknesses and vulnerabilities in the Kubernetes control and data plane. In a nutshell, `kaput` mainly coordinates a collection of existing tools to probe Kubernetes clusters and respectively creates a report for the cluster admin to act on.

I plan to use the following tools to gather candidate attack paths:

- [aquasecurity/kube-hunter](https://github.com/aquasecurity/kube-hunter) 
- [aquasecurity/kube-bench](https://github.com/aquasecurity/kube-bench)
- [banyanops/collector](https://github.com/banyanops/collector)
- [coreos/clair/](https://github.com/coreos/clair/)
- [docker/docker-bench-security](https://github.com/docker/docker-bench-security)
- [nicholasjackson/cnitch](https://github.com/nicholasjackson/cnitch) 
- [OpenSCAP](https://www.open-scap.org/resources/documentation/security-compliance-of-rhel7-docker-containers/)

In addition to the above tools, `kaput` will (at some point in time) implement some simple attacks itself, including but not limited to:

- some of the low-hanging fruits demonstrated in [Hacking & Hardening Kubernetes By Example](https://schd.ws/hosted_files/kccncna17/d8/Hacking%20and%20Hardening%20Kubernetes%20By%20Example%20v2.pdf)
- pod-level: check if cluster is RBAC enabled, check if default SA is used and/or locked down
- service-level: check if it pod can see and communicate with other services in same/different namespaces (`NetworkPolicy` check)
- node-level: poisoning of a node via pod running on the node
- system-level: check if one can get to stuff into the `kube-system` namespace

## Use

Something like:

```
$ kaput --cluster=https://192.168.64.14:8443 --profile=generic,po,svc
Summary: found 12 potential vulnerabilities of which 3 are exploitable
Control plane:
...
```