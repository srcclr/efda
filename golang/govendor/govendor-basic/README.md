# Golang Project using Govendor

This project is an example of a Golang project using Govendor to vendor its dependencies.

`vendor.json` declares the dependencies:
- github.com/google/go-github (version 27c7c32b6d369610435bd2ad7b4d8554f235eb01)
- github.com/google/go-querystring (version 53e6ce116135b80d037921a7fdd5138cf32d7a8a)
- github.com/square/go-jose (version 7f0dd807d3f3d73bb536898cb7980ddf638ce69a)

Total number of dependencies:
- 3 direct dependencies.
- 1 transitive dependency. `github.com/google/go-querystring` is a direct and also transitive dependency (of github.com/google/go-github).
- 3 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.go` and send us a PR.

#### What's next?
By running the security scanner of your choice against `govendor-basic` project, you should get results that are coherent with the details above if the tool supports scanning Golang projects using Govendor as their package manager.
