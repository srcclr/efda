# Golang Project without package manager

This project is an example of a Golang project which does not use any package manager and does not vendor its dependencies. It uses `go get` to always fetch the latest versions of its dependencies.

The dependencies required by this project are:
- github.com/google/go-github
- github.com/google/go-querystring
- github.com/square/go-jose

Total number of dependencies:
- 3 direct dependencies.
- 1 transitive dependency. `github.com/google/go-querystring` is a direct and also transitive dependency (of github.com/google/go-github).
- The number of vulnerabilities depends on the version of the dependencies being used. As of 02 November 2017, the latest versions for each of the dependencies do not have known vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.go` and send us a PR.

#### What's next?
By running the security scanner of your choice against `goget-basic` project, you should get results that are coherent with the details above if the tool supports scanning Golang projects which does not use any package managers.
