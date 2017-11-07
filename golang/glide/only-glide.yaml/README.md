# Golang Project using Glide with only glide.yaml

This project is an example of a Golang project using Glide as its package manager and only glide.yaml file in its folder. It does not check-in glide.lock to lock down the versions.

`glide.yaml` at the root directory declares the dependencies:
- github.com/google/go-github (will always fetch the latest commit)
- github.com/google/go-querystring (will always fetch the latest commit)
- github.com/square/go-jose (version 7f0dd807d3f3d73bb536898cb7980ddf638ce69a)

Total number of dependencies:
- 3 direct dependencies.
- 1 transitive dependency. `github.com/google/go-querystring` is a direct and also transitive dependency (of github.com/google/go-github).
- 3 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.go` and send us a PR.

#### What's next?
By running the security scanner of your choice against `only-glide.yaml` project, you should get results that are coherent with the details above if the tool supports scanning Golang projects using Glide as their package manager.
