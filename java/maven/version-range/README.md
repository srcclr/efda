# Maven Version Range Example

This project is an example of a project using Maven's Version Range feature.

`pom.xml` at the root directory declares the dependencies:
- keycloak-saml-core by org.keycloak. version is specified at `<version>[2.0.0,)</version>`
- neo4j-jmx by org.neo4j. version is specified at `<version>(,1.1.4],[2.0.0,2.0.4)</version>`
- jbcrypt by org.mindrot. version is specified at `<version>[0.3m]</version>`

The setup for `pom.xml` will result in the following versions being resolved:
- For keycloak-saml-core, it will resolve to the latest available version that is >=2.0.0, which as of this writing (02 Nov 2017) is 3.3.0.Final.
- For neo4j-jmx, it will resolve to the latest available version V such that V <= 1.1.4 OR 2.0.0 <= V < 2.0.4, which as of this writing (02 Nov 2017) is 2.0.3.
- For jbcrypt, it will resolve to exactly version 0.3m.

Number of dependencies:
- 3 direct dependencies.
- 4 transitive dependencies.
- 2 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `version-range` project, you should get results that are coherent with the details above if the tool supports Maven's version range feature.
