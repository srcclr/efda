# Composer Scopes

This project is an example of a project using Composer's Dependencies Scopes feature.

`composer.json` at the root directory declares the dependencies:
- require:
 - drupal/core 8.3.6
- require-dev:
 - phpunit/phpunit 5.0.10

This project checks whether the dependency analysis tool is consistent with how Composer resolves dependencies with `scopes`.

The number of direct and transitive dependencies varies according to the scope set to run the dependency analysis tool. If you restrict the scope to `--no-dev`, you will see:
- 1 direct dependency.
- 44 transitive dependencies.
- 3 vulnerabilities.

If you leave the scope as default, you will see:
- 2 direct dependencies.
- 66 transitive dependencies.
- 4 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.php (and adding the necessary dependencies in `composer.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `edfa-composer-scopes-example` project, you should get results that are coherent with the details above if the tool supports using composer.json to manage PHP dependencies.
