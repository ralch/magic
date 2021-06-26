# magic

[![Go Reference](https://pkg.go.dev/badge/github.com/ralch/magic.svg)](https://pkg.go.dev/github.com/ralch/magic)

The package detects a file's mime-type by looking for magic numbers in its content.

## Getting Started

First you should make sure that in your environment,
[libmagic(3)](https://linux.die.net/man/3/libmagic) library is presented.
Otherwise, you can install it for most of the platforms in the following way:

### Mac OS X

```bash
$ brew install libmagic
```

### Debian or Ubuntu

```bash
$ apt-get install libmagic-dev
```

### RHEL, CentOS or Fedora

```bash
$ yum install file-devel
```

Then you can install the package:

```bash
$ go get github.com/ralch/magic
```
