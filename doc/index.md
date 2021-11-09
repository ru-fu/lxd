```{toctree}
:hidden:
:titlesonly:

self
getting_started
configuration
images
operation
restapi_landing
internals
external_resources
```



[![LXD](https://linuxcontainers.org/static/img/containers.png)](https://linuxcontainers.org/lxd)
# LXD
LXD is a next generation system container and virtual machine manager.
It offers a unified user experience around full Linux systems running inside containers or virtual machines.

It's image based with pre-made images available for a [wide number of Linux distributions](https://images.linuxcontainers.org)
and is built around a very powerful, yet pretty simple, REST API.

To get a better idea of what LXD is and what it does, you can [try it online](https://linuxcontainers.org/lxd/try-it/)!
Then if you want to run it locally, take a look at our [getting started guide](https://linuxcontainers.org/lxd/getting-started-cli/).

Release announcements can be found here: <https://linuxcontainers.org/lxd/news/>

And the release tarballs here: <https://linuxcontainers.org/lxd/downloads/>

The documentation is here: <https://linuxcontainers.org/lxd/docs/master/>

## Status
Type                | Service               | Status
---                 | ---                   | ---
CI (client)         | GitHub                | [![Build Status](https://github.com/lxc/lxd/workflows/Client%20build%20and%20unit%20tests/badge.svg)](https://github.com/lxc/lxd/actions)
CI (server)         | Jenkins               | [![Build Status](https://jenkins.linuxcontainers.org/job/lxd-github-commit/badge/icon)](https://jenkins.linuxcontainers.org/job/lxd-github-commit/)
Go documentation    | Godoc                 | [![GoDoc](https://godoc.org/github.com/lxc/lxd/client?status.svg)](https://godoc.org/github.com/lxc/lxd/client)
Static analysis     | GoReport              | [![Go Report Card](https://goreportcard.com/badge/github.com/lxc/lxd)](https://goreportcard.com/report/github.com/lxc/lxd)
Translations        | Weblate               | [![Translation status](https://hosted.weblate.org/widgets/linux-containers/-/svg-badge.svg)](https://hosted.weblate.org/projects/linux-containers/lxd/)
Project status      | CII Best Practices    | [![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/1086/badge)](https://bestpractices.coreinfrastructure.org/projects/1086)
