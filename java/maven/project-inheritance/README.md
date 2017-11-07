# Maven Project Inheritance

This project is an example of a project using Maven's Project Inheritance feature.

`pom.xml` at the root directory declares the dependencies:
- jbcrypt 0.3m by org.mindrot
- apache.sling.engine 2.0.4-incubator by org.apache.sling

There is only one child module. `pom.xml` in `child-module/` inherits from the root's `pom.xml`.

The setup of this project results in the following final configuration:
- 2 direct dependencies.
- 14 transitive dependencies.
- 10 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `child-module` project, you should get results that are coherent with the details above.

**Note**: If you intend to compile this project with `mvn`, you should compile the `child-module` directory and not this directory. The POM file in this directory only declares the dependencies for the child module and does not compile anything.
