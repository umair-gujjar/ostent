[![GoDoc](https://godoc.org/github.com/ostrost/ostent?status.svg)](https://godoc.org/github.com/ostrost/ostent)
[![Travis CI](https://travis-ci.org/ostrost/ostent.svg?branch=master)](https://travis-ci.org/ostrost/ostent)

ostent collects system metrics to display and relay to

- Graphite
- InfluxDB
- Librato

The interactive UI ([**demo**](http://demo.ostrost.com/)):

![Screenshot](https://www.ostrost.com/ostent/screenshot.png)

The metrics collected and exported:
- RAM, swap usage
- CPU usage, load average
- Disk space usage in bytes and inodes
- Network ins and outs in bytes, packets, drops and errors

The processes top is on-display only.

The exporting is kept on par with [collectd](https://collectd.org/)
[reporting](https://collectd.org/wiki/index.php/Plugin:Write_Graphite),
although the metrics naming is slightly different.

## Install

ostent is a single executable (everything is builtin) without dependecies.
Release binaries self-upgrade whenever there's new stable release.
Binaries builds by courtesy of [Travis CI](https://travis-ci.org/ostrost/ostent),
distributed by [GitHub Releases](https://github.com/ostrost/ostent/releases).

Install & run with `curl -sSL https://github.com/ostrost/ostent/raw/master/ostent.sh | sh`

Platforms:

   - [Mac OS X](https://github.com/ostrost/ostent/releases/download/v0.5.0/Darwin.x86_64)
   - Linux [64-bit](https://github.com/ostrost/ostent/releases/download/v0.5.0/Linux.x86_64) | [32-bit](https://github.com/ostrost/ostent/releases/download/v0.5.0/Linux.i686)
   - FreeBSD [64-bit](https://github.com/ostrost/ostent/releases/download/v0.5.0/FreeBSD.amd64) | [32-bit](https://github.com/ostrost/ostent/releases/download/v0.5.0/FreeBSD.i386)

## Usage

```
To continuously export collected metrics to --graphite, --influxdb and/or --librato
specify it like an URL with host part pointing at the server and query being parameters.
E.g. --graphite localhost\?delay=30s

Flags:
  -b, --bind address            Bind address (default :8050)
      --graphite URL            Graphite exporting URL
      --influxdb URL            InfluxDB exporting URL
      --librato URL             Librato exporting URL
      --max-delay delay         Collection and display maximum delay (default 10m)
  -d, --min-delay delay         Collection and display minimum delay (default 1s)
      --version                 Print version and exit
```

Here's how it goes:

```
$ ostent                                     ________________
[ostent]    -------------                   < Spot the links >
[ostent]  / server ostent \                  ----------------
[ostent] +------------------------------+           \   ^__^
[ostent] | http://127.0.0.1:8050        |            \  (oo)\_______
[ostent] |------------------------------|               (__)\       )\/\
[ostent] | http://192.168.1.2:8050      |                   ||----w |
[ostent] +------------------------------+                   ||     ||
```

## Running the code

1. **`go get github.com/ostrost/ostent`**
2. `ostent` to run.

### Rebuilding

1. `cd $(go list -f {{.Dir}} github.com/ostrost/ostent)`
2. `make init` once.
3. Optionally, for asset and template rebuilding only (once):
   - `npm install`: the package list comes from package.json
   - `node_modules/.bin/bower install`: the list from bower.json
4. `gulp watch` or `make` after changes.

`make` rebuilds these **commited to the repo** files:
- `share/assets/bindata.*.go`
- `share/assets/css/*.css`
- `share/assets/js/*/*.js`
- `share/templates/*.html`
- `share/templates/bindata.*.go`
- `share/js/*.jsx`

If you don't change source files, content re-generated
should not differ from the commited.

`gulp watch`

- watches share/{js,style,templatesorigin} and rebuilds dependants on changes
- does live-reloading `ostent` code run
- acceps all ostent flags e.g. `gulp watch -b 127.0.0.1:8080`

### Two kinds of builds

Standalone and release binaries produced by `make` are built with `bin` tag.
These binaries include

- embeded template and assets
- [goagain](https://github.com/rcrowley/goagain) recovering
- self-upgrading via [go-update](https://github.com/inconshreveable/go-update)

This builds also contain `extractassets` subcommand to copy assets on disk.

Non-bin builds have
- `main.dev.go` as entry point in main package
- different set of flags: facilitating debugging etc.
- serving from actual template and asset files as is

This kind (`!bin`) is produced by `gulp watch` and `go get`.
