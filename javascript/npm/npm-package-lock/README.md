# JavaScript Project with package-lock.json

This project is an example of a JavaScript project with a package-lock.json to lockdown the version of its dependencies.

`package-lock.json` at the root directory locked down the following dependencies' versions:
- d3@3.5.17
- hawk@2.3.1
- boom@2.10.1
- cryptiles@2.0.5
- hoek@2.16.3
- sntp@1.0.9
- minimatch@2.0.10
- brace-expansion@1.1.8
- balanced-match@0.4.2
- concat-map@0.0.1

Total number of dependencies:
- 10 direct dependencies (7 of these direct dependencies are also transitives)

Vulnerabilities:
- 2 vulnerabilities

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `package-lock.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `efda-npm-package-lock` project, you should get results that are coherent with the details above if the tool supports using `package-lock.json` to resolve dependencies and lock down versions.
