# C# Project with PackageReference

This project is an example of a C# project that uses PackageReference to manage
its dependencies.

The 2 open-source libraries used in this project are:


1. Common.Logging version="3.4.1"
4. YamlDotNet version="4.3.2"

Total number of dependencies:

* 2 direct dependencies.
* 2 transitive dependencies.
* 1 vulnerability.

Program.cs has a method call at line 22 - 23 that leads to a vulnerability.
