# Gradle Multi-Module Project

This project is an example of a project using Gradle to manage its build and dividing the projects into multiple sub-projects.

`build.gradle` at the root directory declares the submodules to build and the dependencies. The 2 submodules are:
- submodule-1
- submodule-2

The dependencies are:
- jbcrypt 0.3m by org.mindrot
- apache.sling.engine 2.0.4-incubator by org.apache.sling

#### To build the projects:
`gradle build`

#### To run the projects:
`gradle run`: this will build the projects and run them.

### Number of dependencies:
- 2 direct dependencies.
- 0 transitive dependencies.
- 2 vulnerabilities.

#### Vulnerable Call Chain
For `submodule-1`: Main.java has a method call at lines 9-10 that leads to a vulnerability.

For `submodule-2`: This module is not vulnerable as it only declares dependencies on vulnerable libraries and imports them in `Main.java`. No code was called from the libraries.

#### What's next?
By running the security scanner of your choice against `multi-modules` project, the tool should be able to process both sub-projects properly if it supports Gradle's multi-module feature.
