# serf-hosts-updater
serf-hosts-updater is a [Serf](http://www.serfdom.io/) event handler for automatic updating ``/etc/hosts`` file to represent Serf networking.
This is especially useful when used in combination with [Docker](https://www.docker.io/), because docker containers are normally assigned a random IP address.

Compared to a similar software, this handler has several benefits.

* It can use with setuid, so no need to execute serf as root user.
* It can send SIGHUP to dnsmasq in order to enable new entries.

This software developed using Go.

## Installation
To build serf-hosts-updater, you must install Go. Please refer to Go's official documentation.

Then, you can install this handler by:

    $ go get github.com/yosisa/serf-hosts-updater

A single binary file is installed to ``$GOPATH/bin/serf-hosts-updater``.

If you plan to execute serf as a non-root user, you must set proper permissions to this program.
It is recommended to copy this program from ``$GOPATH/bin`` to ``/usr/local/bin``.

    $ sudo cp $GOPATH/bin/serf-hosts-updater /usr/local/bin
    $ sudo chmod u+s,g+s /usr/local/bin/serf-hosts-updater

## Usage
You can use this program as a serf event handler.
When you configure serf by a json file, append path to the program to the list of ``event_handlers``:

```json
{
    "event_handlers": [
        "/usr/local/bin/serf-hosts-updater"
    ]
}
```

Or, enable dnsmasq reloader if dnsmasq is running in your host:

```json
{
    "event_handlers": [
        "/usr/local/bin/serf-hosts-updater -dnsmasq"
    ]
}
```

## License
This software is distributed under the MIT license.
