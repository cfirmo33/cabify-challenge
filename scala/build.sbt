name                         := "cabify-challenge"
organization in ThisBuild    := "com.github.apoloval"
version in ThisBuild         := "0.1.0-SNAPSHOT"
scalaVersion in ThisBuild    := "2.11.8"

libraryDependencies ++= Seq(
  "com.typesafe.akka" %% "akka-actor" % "2.5.2",
  "org.scalatest" %% "scalatest" % "3.0.1" % "test"
)

