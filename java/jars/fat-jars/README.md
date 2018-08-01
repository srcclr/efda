# Fat Jar Example

This project is an example of a project that does not use any build tool to manage its dependencies. This project contains 1 fat jar in the `lib` directory which has all the dependencies it needs.

###### To compile the project:
`mkdir target`

`javac -d "target" -cp ".:lib/io/netty/netty-all/4.1.0.Beta1/netty-all-4.1.0.Beta1.jar" src/main/java/org/continuous_security/Main.java`

###### To run the program:
`java -cp "./target:lib/io/netty/netty-all/4.1.0.Beta1/netty-all-4.1.0.Beta1.jar" org.continuous_security.Main`

Program should exit without errors.

###### To clean the build:
Delete the `target` directory.

The jar used in this project is:
- netty-all-4.1.0.Beta1.jar by io-netty

Total number of dependencies:
- 1 direct dependency.
- 0 transitive dependencies (everything is in the fat jar).
- 4 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at line 38 that leads to a vulnerability. The line calls `MyDecoder.decode()`, which will call `DnsResponseDecoder.decode()`, which eventually calls a private vulnerable method.

#### What's next?
By running the security scanner of your choice against `fat-jars` project, you should get results that are coherent with the details above if the tool is able to locate and identify fat jars correctly.
