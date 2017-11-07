# Ruby Project with version ranges in Gemfile

This project is an example of a Ruby project that uses Gemfile to manage its dependencies. The dependencies listed in Gemfile use semver ranges to specify their versions.

`Gemfile` at the root directory declares the dependencies:

- authlogic: ~> 1.4.1
	- This should resolve to 1.4.3.
- sinatra: >= 1.3.0 , < 1.3.4
	- This should resolve to 1.3.3.
- jbuilder: < 2.1
	- This should resolve to 2.0.8.  

Total number of dependencies:
- 3 direct dependencies
- 16 transitive dependencies

Vulnerabilities:
- 6 vulnerabilities

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.rb (and adding the necessary dependencies in `Gemfile`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `with-version-ranges` project, you should get results that are coherent with the details above if the tool supports version ranges in Gemfile to manage Ruby dependencies.
