# Yarn Scopes

This project is an example of a project using Yarn's Dependencies Scopes feature.

`yarn.lock` at the root directory declares the dependencies:
- dependencies:
 - jquery 3.0.0-alpha1
 - boom
 - console-io 2.6.3
 - cookie-signature 1.0.2
- devDependencies:
 - coffee-script ~1.6.3

This project checks whether the dependency analysis tool is consistent with how Yarn resolves dependencies with `scopes`.

The number of direct and transitive dependencies varies according to the scope set to run the dependency analysis tool. If you restrict the scope to `production`, you will see:
- 4 direct dependencies.
- 147 transitive dependencies.
- 17 vulnerabilities.

If you use the default scope, you will see:
- 5 direct dependency.
- 147 transitive dependencies.
- 17 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `yarn.lock`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `scopes` project, you should get results that are coherent with the details above if the tool supports Yarn's configurable scopes feature.
