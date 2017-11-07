# ObjectiveC Project with only Podfile in the project folder

This project is an example of an ObjectiveC project with only Podfile but not Podfile.lock to lock the versions.

`Podfile` at the root directory declares the dependencies:
- 'ReactiveCocoa', '5.0.0-rc.1'
- 'Alamofire', '~> 4.0'
- 'HockeyKit'
- 'OpenSSL-Static', '1.0.2.c1'

Total number of dependencies:
- 4 direct dependencies.
- 2 transitive dependencies.
- 33 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.m` and send us a PR.

#### What's next?
By running the security scanner of your choice against `only-podfile` project, you should get results that are coherent with the details above if the tool supports scanning ObjectiveC projects that only have Podfile present in their folder.
