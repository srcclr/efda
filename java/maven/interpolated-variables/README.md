# Maven Interpolated Variables

This project is an example of a project using interpolated variables.

`pom.xml` at the root directory declares the properties for the coordinates and version of the dependencies to use.

There is only one child module. `pom.xml` in `child-module/` inherits from the root's `pom.xml` and uses the properties defined in the parent `pom` to define dependencies:
- jbcrypt 0.3m by org.mindrot
- apache.sling.engine 2.0.4-incubator by org.apache.sling

The setup of this project results in the following final configuration:
- 2 direct dependencies.
- 14 transitive dependencies.
- 10 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `interpolated-variables` project, you should get results that are coherent with the details above.
