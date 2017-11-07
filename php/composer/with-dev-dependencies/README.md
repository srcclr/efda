# PHP Project with composer.json (production & development dependencies)

This project is an example of a PHP project that uses composer.json to manage its dependencies. The dependencies listed in  composer.json are also divided into production and development dependencies.

`composer.json` at the root directory declares the dependencies:
- require:
 - moodle/moodle 2.6.2
 - appserver-io/http 1.1.6
 - drupal/core 8.3.6
- require-dev:
 - phpunit/phpunit 5.0.10

Total number of dependencies:
- 4 direct dependencies.
- 69 transitive dependencies.
- 97 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.php (and adding the necessary dependencies in `composer.json`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `edfa-composer-with-dev-dependenciess` project, you should get results that are coherent with the details above if the tool supports using composer.json to manage PHP dependencies.
