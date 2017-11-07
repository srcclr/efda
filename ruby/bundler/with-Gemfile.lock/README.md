# Ruby Project with Gemfile.lock

This project is an example of a Ruby project with only a Gemfile.lock to lockdown the version of its dependencies.

`Gemfile.lock` at the root directory declares the dependencies:

- activesupport (4.2.7.1)
- allison (2.0.3)
- authlogic (1.4.3)
- echoe (4.6.6)
- i18n (0.7.0)
- jbuilder (2.0.8)
- json (1.8.3)
- json_pure (2.0.2)
- minitest (5.9.1)
- multi_json (1.12.1)
- rack (1.6.4)
- rack-protection (1.5.3)
- rake (11.3.0)
- rdoc (4.2.2)
- rubyforge (2.0.4)
- sinatra (1.3.3)
- thread_safe (0.3.5)
- tilt (1.4.1)
- tzinfo (1.2.2)

Total number of dependencies:

- 3 direct dependencies
- 16 transitive dependencies

Vulnerabilities:

- 7 vulnerabilities

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying index.rb (and adding the necessary dependencies in `Gemfile`) and send us a PR.

#### What's next?
By running the security scanner of your choice against `with-Gemfile.lock` project, you should get results that are coherent with the details above if the tool supports using Gemfile.lock instead of Gemfile to resolve dependencies and versions lockdown.
