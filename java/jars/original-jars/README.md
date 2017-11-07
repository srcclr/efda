# Original Jars-only Example

This project is an example of a project that does not use any build tool to manage its dependencies. This project contains 3 jars in the `lib` directory which has all the dependencies it needs. The jars are original open-source jars but have been renamed.

###### To compile the project:
`mkdir target`

`javac -d "target" -cp ".:lib/org/barorg/bar/version/bar-version.jar:lib/org/spamframework/spam-web/version.RELEASE/spam-web-version.RELEASE.jar:lib/com/foodatabase/foo/version/foo-version.jar" src/main/java/org/continuous_security/Main.java`

###### To run the program:
`java -cp "./target:lib/org/barorg/bar/version/bar-version.jar" org.continuous_security.Main <arg1>`

Program should exit without errors.

###### To clean the build:
Delete the `target` directory.

The 3 open-source libraries (jars) used in this project are:

|Library|Renamed To|Transitive Dependencies|Vulnerabilities|
|-------|----------|-----------------------|---------------|
|h2-1.3.176.jar by com.h2database|`com/foodatabase/foo/version/foo-version.jar`|0|2|
|jbcrypt-0.3m.jar by org.mindrot|`org/barorg/bar/version/bar-version.jar`|0|1|
|spring-web-3.1.1.RELEASE.jar by org.springframework |`org/spamframework/spam-web/version.RELEASE/spam-web-version.RELEASE.jar`|8|6|

Total number of dependencies:
- 3 direct dependencies.
- 8 transitive dependencies.
- 9 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `original-jars` project, you should get results that are coherent with the details above if the tool is able to locate and identify jars correctly.
