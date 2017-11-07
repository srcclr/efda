# SBT project with version ranges in build.sbt

This project is an example of a SBT project that uses version ranges to manage its dependencies.

`build.sbt` at the root directory declares the dependencies:
- org.keycloak:keycloak-saml-core:[2.0.0,)
- net.liftweb:lift-webkit:latest.integration

Total number of dependencies:
- 3 direct dependencies (Scala standard library is considered as a direct dependency.).
- 26 transitive dependencies.
- 4 vulnerabilities.
The number of dependencies and vulnerabilities depend on the actual version of the dependencies resolved. As of this writing (2/11/2017), org.keycloak:keycloak-saml-core:[2.0.0,) resolves to 3.3.0.Final, and net.liftweb:lift-webkit:latest.integration resolves to 3.2.0-M2.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `Main.scala` and send us a PR.

#### What's next?
By running the security scanner of your choice against `version-ranges` project, you should get results that are coherent with the details above if the tool supports SBT projects with version ranges in their `build.sbt` file.
