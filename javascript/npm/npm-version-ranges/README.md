# JavaScript Project with version ranges in package.json

This project is an example of a JavaScript project that uses package.json to manage its dependencies. The dependencies listed in package.json use semver ranges to specify their versions.

`package.json` at the root directory declares the dependencies:
- jquery: ^3.0.0 <3.1.1
  - This should resolve to 3.1.0.
- backbone: >=0.2.0 <0.3.3
  - This should resolve to 0.3.2.
- hawk: <2.3.1
  - This should resolve to 2.3.0.

Total number of dependencies:
- 3 direct dependencies.
- 5 transitive dependencies.
- 3 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `package.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `npm-version-ranges` project, you should get results that are coherent with the details above if the tool supports version ranges in package.json to manage JavaScript dependencies.
