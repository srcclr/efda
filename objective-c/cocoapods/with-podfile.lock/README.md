# ObjectiveC Project with only Podfile.lock in the project folder

This project is an example of an ObjectiveC project with only Podfile.lock to lockdown the version of its dependencies.

`Podfile.lock` at the root directory declares the dependencies:
- Alamofire (4.3.0)
- HockeyKit (2.0.7)
- OpenSSL-Static (1.0.2.c1)
- ReactiveCocoa (5.0.0-rc.1):
- ReactiveSwift (1.0.0):
- Result (3.1.0)

Total number of dependencies:
- 4 direct dependencies.
- 2 transitive dependencies.
- 33 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `main.m` and send us a PR.

#### What's next?
By running the security scanner of your choice against `with-podfile.lock` project, you should get results that are coherent with the details above if the tool supports using Podfile.lock instead of Podfile to resolve dependencies and versions lockdown.
