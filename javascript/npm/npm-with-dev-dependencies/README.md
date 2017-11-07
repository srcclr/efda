# JavaScript Project with package.json (production & development dependencies)

This project is an example of a JavaScript project that uses package.json to manage its dependencies. The dependencies listed in package.json are also divided into production and development dependencies.

`package.json` at the root directory declares the dependencies:
- dependencies:
 - jquery 3.0.0-alpha1
 - express 4.1.1
 - angular 1.3.19
- devDependencies:
 - coffee-script ~1.6.3

Total number of dependencies:
- If all dependencies are included:
 - 4 direct dependencies.
 - 20 transitive dependencies.
 - 18 vulnerabilities.
- If scope is limited to "production":
 - 3 direct dependencies.
 - 20 transitive dependencies.
 - 18 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `package.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `npm-with-dev-dependencies` project, you should get results that are coherent with the details above if the tool supports using package.json to manage JavaScript dependencies.
