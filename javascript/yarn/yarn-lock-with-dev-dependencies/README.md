# JavaScript Project with yarn.lock

This project is an example of a JavaScript project with a yarn.lock to lockdown the version of its dependencies.

`yarn.lock` at the root directory declares the dependencies:
- dependencies:
 - jquery 3.0.0-alpha1
 - boom
 - console-io 2.6.3
 - cookie-signature 1.0.2
- devDependencies:
 - coffee-script ~1.6.3

Total number of dependencies:
- 5 direct dependencies.
- 147 transitive dependencies.
- 17 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `yarn.lock`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `yarn` project, you should get results that are coherent with the details above if the tool supports using yarn.lock instead of package.json to resolve dependencies and versions lockdown.
