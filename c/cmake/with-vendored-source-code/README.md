# C / C++ CMake with Dependencies as Source

This project is an example of a C / C++ project built with CMake that has
dependencies vendored in-tree as source code.

The 1 open-source library used in this project is:

1. SQLite3 version="3.22.0"

Total number of dependencies:

* 1 direct dependencies.
* 0 transitive dependencies.
* 1 vulnerability.

main.cpp has a method call at line 14 - 21 that leads to a vulnerability.

# Build

Run the following commands to build the project.

```shell
mkdir build/
cd build
cmake ..
make
```
