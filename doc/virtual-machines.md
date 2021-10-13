# Virtual Machines

Since version 4.0, LXD supports virtual machines alongside containers. Thanks to a built-in agent, they can be used almost like containers.

They are implemented through the use of `qemu`.

Please note, currently not all features that are available with containers have been implemented for VMs,
however we continue to strive for feature parity with containers.

!!! Note
    Add some information from [https://linuxcontainers.org/lxd/#containers-and-virtual-machines](https://linuxcontainers.org/lxd/#containers-and-virtual-machines).

## Configuration
See [instance configuration](instances.md) for valid configuration options.

All categories and keys that contain the terms `virtual-machine` or `VM` are supported.

## Tutorial

See [Running virtual machines with LXD](https://discuss.linuxcontainers.org/t/running-virtual-machines-with-lxd-4-0/7519) for a tutorial on virtual machines, including a short howto for a Microsoft Windows VM.
