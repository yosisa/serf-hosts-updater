# serf-hosts-updater
serf-hosts-updater is a [Serf](http://www.serfdom.io/) event handler for automatic updating ``/etc/hosts`` file to represent Serf networking.
This is especially useful when used in combination with [Docker](https://www.docker.io/), because docker containers are normally assigned a random IP address.

Compared to a similar software, this handler has several benefits.

* It can use with setuid, so no need to execute serf as root user.

This software developed using Go.

## Installation
To build serf-hosts-updater, you must install Go. Please refer to Go's official documentation.

Then, you can install this handler by:

    $ go get github.com/yosisa/serf-hosts-updater

A single binary file is installed to ``$GOPATH/bin/serf-hosts-updater``.

## License
This software is distributed under the MIT license.
