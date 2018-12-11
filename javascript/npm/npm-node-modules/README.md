# JavaScript Project with only node_modules

This project is an example of a JavaScript project that only has `node_modules` and without `package.json` to manage its dependencies or lock files. It was created based on npm-version-ranges project.

`node_modules` includes the following dependency:
- hawk: 2.3.0

Total number of dependencies:
- 1 direct dependency.
- 4 transitive dependencies.
- 1 vulnerability.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.js (and adding the necessary dependencies in `node_modules`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `npm-node-modules` project, you should get results that are coherent with the details above if the tool supports scanning a project that only has `node_modules`.
