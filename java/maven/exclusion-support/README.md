# Maven Exclusion Support

This project is an example of a project using Maven's Exclusion Support feature.

`pom.xml` at the root directory declares the dependencies:
- sling.engine 2.0.4-incubator by org.apache.sling
- kafka_2.11 0.9.0.1 by org.apache.kafka

The following dependencies are excluded explicitly from `kafka`:
- netty (no transitive dependencies)
- kafka-clients (2 transitive dependencies)

After dependencies exclusion, the final resolved dependencies are:
- 2 direct dependency.
- 27 transitive dependencies.
- 14 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at line 13 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `exclusion-support` project, you should get results that are coherent with the details above if the tool supports Maven's exclusion feature.
