# Maven Project Aggregation

This project is an example of a project using Maven's Project Aggregation feature.

`pom.xml` at the root directory declares the submodules to build. Each submodule declares its own dependencies. The 2 submodules are:
- submodule-1
- submodule-2

submodule-1 dependencies:
- jbcrypt 0.3m by org.mindrot

submodule-2 dependencies:
- apache.sling.engine 2.0.4-incubator by org.apache.sling

The setup of this project results in the following final configuration:
- submodule-1
 - 1 direct dependency.
 - 0 transitive dependencies.
- submodule-2
 - 1 direct dependency.
 - 14 transitive dependencies.
Total vulnerabilities:
 - 10

#### Vulnerable Call Chain
For `submodule-1`: Main.java has a method call at lines 9-10 that leads to a vulnerability.

For `submodule-2`: This module is not vulnerable as it only declares dependencies on vulnerable libraries and imports them in `Main.java`. No code was called from the libraries.

#### What's next?
By running the security scanner of your choice against `project-aggregation` project, you should get results that are coherent with the details above.
