name := "multi-projects"

version := "1.0"

scalaVersion := "2.11.8"

lazy val utils = project

lazy val app = project.
  dependsOn(utils)