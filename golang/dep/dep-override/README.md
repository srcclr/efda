# Golang Project using Dep

This project is an example of a Golang project using Dep to vendor its dependencies. It is based on https://github.com/srcclr/example-go-golangdep but with an override in `Gopkg.toml`.

`Gopkg.toml` declares the dependencies:
- github.com/google/go-github (version 173908142398084e14d7f14c39924f0004d43679)
- github.com/google/go-querystring (version 53e6ce116135b80d037921a7fdd5138cf32d7a8a)
- github.com/simeji/jid (version 0.7.2)
- github.com/square/go-jose (version 7f0dd807d3f3d73bb536898cb7980ddf638ce69a)

`main.go` declares the dependencies:
- github.com/google/go-querystring
- github.com/russellhaering/gosaml2

Total number of dependencies:
- 5 direct dependencies.
- 12 transitive dependencies.
- 1 vulnerability.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.go` and send us a PR.

#### What's next?
By running the security scanner of your choice against `dep-override` project, you should get results that are coherent with the details above if the tool supports scanning Golang projects using Dep as their package manager.

When using SourceClear agent, follow the instructions under the section "Setup and Configuration" in https://www.sourceclear.com/docs/command-line-interface/ to install and activate our SourceClear agent. To scan this project:
```
1. git clone https://github.com/srcclr/efda/ $GOPATH/src/github.com/srcclr/efda/
2. srcclr scan $GOPATH/src/github.com/srcclr/efda/golang/dep/dep-override/
```
