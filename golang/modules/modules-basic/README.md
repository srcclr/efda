# Golang Project using Go Modules

This project is an example of a Golang project using Go Modules to manage its dependencies.

`go.mod` declares the dependencies:
- github.com/google/go-github v29.0.2
- github.com/google/go-querystring v1.0.0
- gopkg.in/square/go-jose.v2 v2.4.1

Total number of dependencies:
- 3 direct dependencies.
- 8 transitive dependencies. `github.com/google/go-querystring` is a direct and also transitive dependency (of github.com/google/go-github).
- 0 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.go` and send us a PR.

#### What's next?
By running the security scanner of your choice against `modules-basic` project, you should get results that are coherent with the details above if the tool supports scanning Golang projects using Go Modules as their dependency manager.

