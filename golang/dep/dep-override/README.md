# [:] Example Go project using Dep.

An example Go project using Golang's Dep to demonstrate [SourceClear](https://www.sourceclear.com) scans. Note that this project is based on https://github.com/srcclr/example-go-golangdep but with an override in `Gopkg.toml`.

## Install and activate SourceClear
Follow the instructions under the section "Setup and Configuration" in https://www.sourceclear.com/docs/command-line-interface/ to install and activate our SourceClear agent.

## Scan this project
To scan this project:

```
1. git clone https://github.com/srcclr/efda/ $GOPATH/src/github.com/srcclr/efda/
2. srcclr scan $GOPATH/src/github.com/srcclr/efda/golang/dep/dep-override/
```
