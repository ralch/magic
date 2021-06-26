# magic

The package detects a file's mime-type by looking for magic numbers in its content.

[![Go Reference](https://pkg.go.dev/badge/github.com/ralch/magic.svg)](https://pkg.go.dev/github.com/ralch/magic)

## Getting Started

First you should make sure that in your environment [libmagic(3)](https://linux.die.net/man/3/libmagic) library is presented. You
can install it for most of the platforms in the following way:

```bash
# Mac OS X
$ brew install libmagic
```

```bash
# Debian or Ubuntu
$ apt-get install libmagic-dev
```

```bash
# RHEL, CentOS or Fedora
$ yum install file-devel
```

Then you can install the package:

```bash
$ go get github.com/ralch/magic
```

# License
