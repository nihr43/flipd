# flipd

flipd (floating ip daemon) is the pure-L3, broadcast domain / availability zone agnostic alternative to keepalived.

## use cases

flipd provides an imperfect (there will be tcp resets on failover) but simple solution for:

- HA unicast or anycast reachability for cluster control planes - k8s, lxd, cockroachdb, etc.
- HA unicast or anycast reachability for HAproxy instances
- floating default gateways for north/south vxlan routing

## motivation and differences with keepalived

Keepalived et al. (VRRP, CARP) essentially rely on ARP for failover convergence.  If we wanted to use keepalived to provide failover for administrative kubernetes api access, we would have to to run all our k8s control plane nodes (kube-apiserver primarily) in the same broadcast domain; likely in one rack.  By using higher-level tools to achieve convergence (and then advertising the ips with something like FRR), we are able to forgo L2 limitations entirely and achieve availability-zone agnostic address reachability.

## build and install

The Makefile has a target to build a .deb:

```
$ make deb
$ sudo dpkg -i flipd
```

## todo

- http watchdog
- systemd unit watchdog
- cluster membership - hashicorp/memberlist?
