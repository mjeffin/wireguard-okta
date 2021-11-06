# Wireguard Okta

[Wireguard](https://www.wireguard.com/) is an extremely fast, secure and modern VPN solution. It is now included in linux kernel
and is used by VPN providers like mozilla vpn and mullad. Wireguard follows the linux philosophy of doing one thing well. It deals with 
cryptography and network routing and doesn't handle identity management. It explicitly leaves identity management to application layer.

One major deterrent for wireguard adoption by enterprises is lack of 2FA. There are awesome products like [tailscale](https://tailscale.com/), which
offers a zero config VPN solution built on top of wireguard. They even have open sourced their [code](https://github.com/tailscale/tailscale).
You could creat islands of intranet within your corporate network and is a great option for large enterprises. 

The aim of this project is to explore wireguard in detail and create a simple system to link wireguard and okta together.
Once okta is integrated, other OIDC providers will be integrated in a generic way and will be exposed as a go package. 

# High Level Objectives

## Phase 1
- Periodically sync okta users and wireguard peers and add/delete peers as needed, using either
  - event webhooks
  - cronjob

Webhooks are more efficient but cron have lesser attack surface and is more robust. Hence, we'll use cron first.

## Phase 2
- Figure out a way to do key rotation and provide means for users to download new keys after they authenticate with okta 2FA

# Configuration 

Configuration settings required by the project so far. Subject to change

```shell
OKTASERVER_API_TOKEN=
OKTASERVER_WIREGUARD_GROUP_ID=
OKTASERVER_ORG_URL=
WG_INTERFACE_IP=10.49.0.1/24
ALLOWED_IPS=10.0.0.0/8
```

# Libraries Used

- [wgctrl-go](https://github.com/WireGuard/wgctrl-go)  - Package wgctrl enables control of WireGuard interfaces on multiple platforms.
- [wireguard-go](https://github.com/WireGuard/wireguard-go) - Go implementation of wireguard protocol. Used by wgctrl-go internally
- [go-sqlite3](github.com/mattn/go-sqlite3 ) 

# Inspirations

Would like to thank contributors of below open source projects, who has travelled before me in this path

- [tailscale](https://tailscale.com/)
- [wired-vpn](https://github.com/jbauers/wired-vpn)

# Progress so far
 - [x] Creation of wireguard client programmatically in the server using wgctrl-go package
 - [x] Fetching of users in okta and comparison of that list to the list of users in the sqlite db

# Next Steps
- [ ] Save peer configuration as text file, to share with user 