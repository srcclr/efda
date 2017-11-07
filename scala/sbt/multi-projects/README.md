# SBT project with a multi-project setup

This project is an example of a SBT project with a multi-project setup.

`build.sbt` at the root directory declares that the `app` project depends on `utils`.
`build.sbt` at the `app` directory declares the dependencies:
- org.mindrot:jbcrypt:0.3m
- net.liftweb:lift-webkit:3.0-M4

Total number of dependencies:
- 3 direct dependencies (Scala standard library is considered as a direct dependency.).
- 21 transitive dependencies.
- 6 vulnerabilities.

#### Vulnerable Call Chain
This project does not have a vulnerable call chain. If you would like to test this feature, feel free to create a vulnerable call chain by modifying `Main.scala` and send us a PR.

#### What's next?
By running the security scanner of your choice against `multi-project` project, you should get results that are coherent with the details above if the tool supports SBT projects with multi-project setup.
