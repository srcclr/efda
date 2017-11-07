# Apache Ant With Apache Ivy Example

This project is an example of a project that uses Apache Ant with Ivy to manage the build and dependencies of the project.

`ivy.xml` declares the dependencies of the project. By default, Ivy fetches these dependencies from the Maven repository.

The 2 dependencies are:
- jbcrypt 0.3m by org.mindrot
- kafka_2.11 0.9.0.1 by org.apache.kafka

`build.xml` declares several Ant targets:
- `resolve` -> retrieves dependencies with Ivy.
- `report` -> generates a report of dependencies into the build directory.
- `build` -> compiles the project.
- `run` -> runs the project.
- `clean` -> cleans the project build.
- `clean-lib` -> cleans the directory where the dependencies are stored.
- `clean-all` -> cleans both the project build and dependencies directory.
- `clean-cache` -> cleans the ivy cache in your local machine.

Dependencies are downloaded into `lib` directory and the build directory is `target`.

###### To compile and run the project:
`ant run`: This will call `build` to build the project and run the compiled program.

Total number of dependencies:
- 2 direct dependencies.
- 16 transitive dependencies.
- 14 vulnerabilities.

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `with-ivy` project, you should get results that are coherent with the details above if the tool is able to support Apache Ant with Ivy.
