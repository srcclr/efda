# Ruby Project with Gemfile (production & development dependencies)

This project is an example of a Ruby project that uses Gemfile to manage its dependencies. The dependencies listed in Gemfile are also divided into production and development dependencies.

`Gemfile` at the root directory declares the dependencies:

- dependencies:
	- authlogic: 1.4.3
	- sinatra: 1.3.3
	- jbuilder: 2.0.8
- development dependencies:
	- ember-source: ~> 1.8.0

Total number of dependencies:

- If all dependencies are included:
	- 4 direct dependencies
	- 17 transitive dependencies
	- 8 vulnerabilities
- If scope is limited to "development":
	- 1 direct dependencies
	- 1 transitive dependencies
	- 2 vulnerabilities

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.rb (and adding the necessary dependencies in `Gemfile`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `with-dev-dependencies` project, you should get results that are coherent with the details above if the tool supports scopes in Gemfile to manage Ruby dependencies.
