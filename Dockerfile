FROM gradle:jdk21 AS cache
WORKDIR /app
ENV GRADLE_USER_HOME /cache
COPY build.gradle gradle.properties settings.gradle ./
RUN gradle --no-daemon build --stacktrace

FROM gradle:jdk21 AS builder
WORKDIR /app
COPY --from=cache /cache /home/gradle/.gradle
COPY build.gradle gradle.properties settings.gradle gradlew ./
COPY src/ src/
RUN gradle --no-daemon build --stacktrace

# Stage 2: Run Stage
FROM ghcr.io/graalvm/jdk-community:21

WORKDIR /app

# Copy only the built artifacts from the previous stage
COPY --from=builder /app/build/libs/. /app/.

# Expose the port that your Micronaut application is running on
EXPOSE 3339

# Create a non-root user named "trunk"
RUN useradd -m -u 1001 spcv

# Change ownership of the application directory to the non-root user
RUN chown -R spcv /app

# Switch to the non-root user
USER spcv

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "ping localhost:3339 -c 1" ]" ]

# Specify the command to run your application when the container starts
CMD ["java", "-jar", "default-0.2-all.jar"]
