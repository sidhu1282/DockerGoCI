# DockerGoCI
Docker GO CI Pipeline

Problem

Create a simple application which has a single "/healthcheck" endpoint.
Containerise your application as a single deployable artifact, encapsulating all dependencies.
Create a CI pipeline for your application


Solution

I have used Go Lang programming language for this task.

In order to dockerize our environment i have written Dockerfile for the server. I built the image from the Dockerfile and Containerized the build image.

Used Maven as a build tool to create a single deployable artifact which includes all dependencies.

Used Jenkins to create CI Pipeline which includes Junit to impement unit test cases.

Once the application is build and running suucssfully on the server, it will give below output.

"myapplication": [
{
"version": "1.0",
"description" : "pre-interview technical test",
"lastcommitsha": "abc57858585"
}
]

I have used concept of Git Webhook to implement CI pipeline which has to be executed when new code is commit and pushed on Github.


Conclusion
In this repository we demonstrated a /healthcheck endpoint for a server implemented in Go, for a docker container.
