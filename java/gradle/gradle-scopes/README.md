# Gradle Scopes

This project is an example of a project using Gradle to manage the build.

There are 2 scopes used in the build file: `compile` and `testCompile`.

`build.gradle` at the root directory declares the dependencies:
- sling.engine 2.0.4-incubator by org.apache.sling (compile)
- kafka_2.11 0.9.0.1 by org.apache.kafka (compile)
- junit 4.12 by junit (testCompile)

#### To build the project:
`gradle build`

#### To run the project:
`gradle run`: this will build the project and run it.

#### The number of direct and transitive dependencies:
- 2 direct dependencies.
- 0 transitive dependencies.
- 3 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at line 13 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `gradle-scopes` project, the tool should be able to detect the correct number of dependencies depending on the scope passed in if it supports Gradle's configurable scopes.
