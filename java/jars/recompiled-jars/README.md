# Recompiled Jars Example

This project is an example of a project that does not use any build tool to manage its dependencies. This project contains 2 jars in the `lib` directory which has all the dependencies it needs. The jars are open-source jars but have been renamed and recompiled with additional manifest data.

###### To compile the project:
`mkdir target`

`javac -d "target" -cp ".:lib/org/hamorg/ham/version/ham-version.jar:lib/org/eggsframework/eggs-web/version.RELEASE/eggs-web-version.RELEASE.jar" src/main/java/org/continuous_security/Main.java`

###### To run the program:
`java -cp "./target:lib/org/hamorg/ham/version/ham-version.jar:lib/org/eggsframework/eggs-web/version.RELEASE/eggs-web-version.RELEASE.jar" org.continuous_security.Main <arg1>`

Program should exit without errors.

###### To clean the build:
Delete the `target` directory.

The 2 open-source libraries (jars) used in this project are:

|Library|Renamed To|Transitive Dependencies|Vulnerabilities|
|-------|----------|-----------------------|---------------|
|jbcrypt-0.3m.jar by org.mindrot|`org/hamorg/ham/version/ham-version.jar`|0|1|
|spring-web-3.1.1.RELEASE.jar by org.springframework |`org/eggsframework/eggs-web/version.RELEASE/eggs-web-version.RELEASE.jar`|8|6|

Total number of dependencies:
- 2 direct dependencies.
- 8 transitive dependencies.
- 7 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `recompiled-jars` project, you should get results that are coherent with the details above if the tool is able to locate and identify recompiled jars correctly.
