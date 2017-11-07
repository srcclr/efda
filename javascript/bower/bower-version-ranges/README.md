# JavaScript Project with bower.json and version ranges

This project is an example of a JavaScript project that uses Bower to manage its dependencies. The dependencies listed in `bower.json` use semver ranges to specify their versions.

`bower.json` at the root directory declares the dependencies:
- dojo: 1.3.0
 - This should resolve to 1.3.0.
- validator-js: <=1.0.0
 - This should resolve to 1.0.0.
- mapbox.js: ^1.5.0 <=1.6.4
 - This should resolve to 1.6.4.
- ember: 1.2.0
 - This should resolve to 1.2.0.

Total number of dependencies:
- 4 direct dependencies.
- 1 transitive dependencies.
- 12 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `bower.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `bower-version-ranges` project, you should get results that are coherent with the details above if the tool supports version ranges in `bower.json` to manage JavaScript dependencies.
