# Apache Ant Example

This project is an example of a project that uses Apache Ant to manage the build and without any other tools (such as Ivy) to manage dependencies.

There are 2 subfolders in this project: `lib` and `project`. `lib` contains all the third-party jars and `project` contains the source code and Ant build file `build.xml`.

`build.xml` declares several Ant targets:
- `build` -> compiles the project.
- `run` -> runs the project.
- `clean` -> cleans the project build.

###### To compile and run the project:
`ant run`: This will call `build` to build the project and run the compiled program.

Total number of dependencies:
- Excluding javadocs and sources jar files, we have 18 jars in the `lib` folder. All jars are considered direct dependencies.
- 14 vulnerabilities

#### Vulnerable Call Chain
Main.java has a method call at lines 9-10 that leads to a vulnerability.

#### What's next?
By running the security scanner of your choice against `without-ivy` project, you should get results that are coherent with the details above if the tool is able to support Apache Ant without Ivy.
