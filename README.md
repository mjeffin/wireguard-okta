# Wireguard Okta

[Wireguard](https://www.wireguard.com/) is an extremely fast, secure and modern VPN solution. It is mow included in linux kernel
and is used by VPN providers like mozilla vpn and mullad. Wireguard follows the linux philosphy of doing one thing well. It deals with 
cryptography and network routing and doesn't handle identity management. It explicitly leaves identity management to application layer.

One major deterrent of wireguard adoption by enterprises is lack of 2FA. There are awesome products like [tailscale](https://tailscale.com/) 
offers a zero config VPN solution built on top of wireguard. They even have open sourced their [client code](https://github.com/tailscale/tailscale).
It is great for creating islands of intranet within your corporate network and is a good option for large enterprises. Becuase it does so many 
things, it might take some time for a network admin to know what actually goes under the hood. 

The easiest way to install wireguard is using [algo](https://github.com/trailofbits/algo), a set of Ansible scripts that simplify the setup of a 
personal WireGuard and IPsec VPN. It's a great and simple way to setup wireguard for a number of users, but the only trouble is key rotation and lack 
of 2FA. This project is aimed at provisioning and key management for enterprises simple and tied to their Identity provider system. 

Okta will be integrated first and the learnings from builing that will be used to create a generic system with integration to multiple OIDC providers. 


# High Level Requirements

- Provide webhooks/api endpoints to add and remove wireguard user accounts automatically based on Okta events
- Provide means to authenticate via okta and refresh the certificate once every day

# Inspirations

Would like to thank contributors of below open source projects, who has travelled before me in this path

- [tailscale](https://tailscale.com/) 
- [wired-vpn](https://github.com/jbauers/wired-vpn)
- 
