# Maven Scopes

This project is an example of a project using Maven's Dependencies Scopes feature.

`pom.xml` at the root directory declares the dependencies:
- sling.engine 2.0.4-incubator by org.apache.sling
- kafka_2.11 0.9.0.1 by org.apache.kafka
- junit 4.12 by junit

This project checks whether the dependency analysis tool is consistent with how Maven resolves dependencies with `scopes`. For details on Maven's scopes resolution, see https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html#Dependency_Scope.

The setup of this project results in the following final configuration:
- `sling.engine` with the default scope, which resolves to `compile`.
- `kafka` with `<scope>compile</scope>`.
- `junit` with `<scope>test</scope>`.

The number of direct and transitive dependencies varies according to the scope set to run the dependency analysis tool. If you restrict the scope to `test`, you will see:
- 3 direct dependencies.
- 31 transitive dependencies.
- 15 vulnerabilities.

If you restrict the scope to `compile`, you will see:
- 2 direct dependency.
- 30 transitive dependencies.
- 22 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at line 13 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `scopes` project, you should get results that are coherent with the details above if the tool supports Maven's configurable scopes feature.
