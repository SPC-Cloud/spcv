# Suraj Patel CV
- A Java, [Micronaut](https://micronaut.io/) Docker based single page application.
- Currently trunk is being built and deployed to [Google Cloud](https://cloud.google.com/?hl=en) using [Cloudbuild](https://cloud.google.com/build?hl=en) and [Cloud Run](https://cloud.google.com/run?hl=en).
  - All infrastructure is deployed and configured using [Pulumi](https://pulumi.com/) - Go
  - Using [Pulumi Cloud](https://pulumi.com/) and GitHub Actions to preview infrastructure changes on PRs
  - Pulumi Cloud is also setup to deploy the infrastructure stack from their dashboard using [Pulumi Deployments](https://www.pulumi.com/docs/pulumi-cloud/deployments/)
- Using [Canva](https://www.canva.com/) to embed the majority of the front-end logic.
- Branches are currently being built using GitHub Actions.

## Build and run instructions
- Clone the Github Repo locally.
- There are three build options:
  - Run a `docker build . -t spcv`
    - This will build a docker container called spcv locally using a [multi-stage](https://docs.docker.com/build/building/multi-stage/) Dockerfile.
  - Run a `docker compose up`
    - This will build and run the spcv container already built, or will use the aforementioned Dockerfile to build and run.
  - Run a `gradle clean build`
    - This will assemble a shadow jar that can be executed locally (gradle run)
- The app will be accessible on port 3339, e.g. `localhost:3339`



## Micronaut 4.3.3 Documentation

- [User Guide](https://docs.micronaut.io/4.3.3/guide/index.html)
- [API Reference](https://docs.micronaut.io/4.3.3/api/index.html)
- [Configuration Reference](https://docs.micronaut.io/4.3.3/guide/configurationreference.html)
- [Micronaut Guides](https://guides.micronaut.io/index.html)
---

- [Micronaut Gradle Plugin documentation](https://micronaut-projects.github.io/micronaut-gradle-plugin/latest/)
- [GraalVM Gradle Plugin documentation](https://graalvm.github.io/native-build-tools/latest/gradle-plugin.html)
- [Shadow Gradle Plugin](https://plugins.gradle.org/plugin/com.github.johnrengelman.shadow)
## Feature serialization-jackson documentation

- [Micronaut Serialization Jackson Core documentation](https://micronaut-projects.github.io/micronaut-serialization/latest/guide/)


## Feature ksp documentation

- [Micronaut Kotlin Symbol Processing (KSP) documentation](https://docs.micronaut.io/latest/guide/#kotlin)

- [https://kotlinlang.org/docs/ksp-overview.html](https://kotlinlang.org/docs/ksp-overview.html)


## Feature micronaut-aot documentation

- [Micronaut AOT documentation](https://micronaut-projects.github.io/micronaut-aot/latest/guide/)


