# Helm registry
A simple Helm registry for store and share Helm package
\
[![](https://img.shields.io/badge/Docker-compose?logo=docker&logoColor=white&color=blue)]()
[![](https://img.shields.io/badge/registry-helm?logo=helm&logoColor=white&label=Helm&labelColor=darkblue&color=white)]()
[![](https://img.shields.io/badge/Golang-1.23.3-grey?style=for-the-badge&logo=go&labelColor=cyan)]()


This project is for have a simple [Helm registry](https://helm.sh/docs/topics/chart_repository/) and a web admin UI for manage registry.
There are 2 parts :
- **Backend** : Build in Go, API for store Helm package
- **Frontend** : Build in ReactJS, for Web UI to manage Helm registry

These parts are Docker images, that you use and deploy him in Docker env.
![](architecture.png)

## Get stated :rocket:

### Environment Variables
It exists some environment variables to customize the registry.

| Var | Description                                                         |
| :-- |:--------------------------------------------------------------------|
| `REGISTRY_NAME` | Name of the registry                                                |
| `REGISTRY_DESCRIPTION` | A description of the registry                                       |
| `REGISTRY_VERSION` | Version of the registry                                             |
| `REGISTRY_MAINTAINER` | Name of the registry maintainer. Can be a person or an organisation |
| `REGISTRY_MAINTAINER_URL` | URL of the website or email address of the registry maintainer      |
| `REGISTRY_LABELS` | List (separed by ';') of labels. E.g : `env:prod;project:test`      |

## Version
[![](https://badgen.net/github/tag/BenjaminFourmaux/helm-registry?cache=600)](https://github.com/BenjaminFourmauxhelm-registry/tags) [![](https://badgen.net/github/release/BenjaminFourmaux/helm-registry?cache=600)](https://github.com/BenjaminFourmaux/helm-registry/releases)
- [coming soon][v1] First API version with basic actions

## Contributors 👪
[![](https://badgen.net/github/contributors/BenjaminFourmaux/helm-registry)](https://github.com/BenjaminFourmaux/helm-registry/graphs/contributors)
- :crown: [Benjamin Fourmaux](https://github.com/BenjaminFourmaux)

## Licence ⚖️
All files on this project is under [**Apache License v2**](https://www.apache.org/licenses/LICENSE-2.0).
You can:
- Reuse the code 
- Modified the code
- Build the code

You must **Mention** the © Copyright if you use and modified code for your own profit. Thank you

© 2004 - Benjamin Fourmaux -- Beruet - All right reserved
