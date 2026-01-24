# Stage 1: Build Stage
FROM ghcr.io/graalvm/native-image-community:21 AS builder

# Install findutils to provide xargs, which enables the Gradle wrapper to work
RUN microdnf install -y findutils

WORKDIR /app

# Copy gradle files first for caching
COPY build.gradle gradle.properties settings.gradle gradlew ./
COPY gradle gradle

# Copy source code
COPY src src

# Build native executable
RUN ./gradlew nativeCompile --no-daemon

# Stage 2: Run Stage
FROM ubuntu:noble

WORKDIR /app

# Copy only the built artifacts from the previous stage
COPY --from=builder /app/build/native/nativeCompile/default /app/spcv

# Expose the port
EXPOSE 3339

# Create a non-root user
RUN useradd -m -u 1001 spcv && \
    chown -R spcv /app

# Switch to the non-root user
USER spcv

# Healthcheck
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD ["/app/spcv", "--health"] || exit 1 
# Note: Basic check, refine if app has specific health endpoint handling from shell

CMD ["/app/spcv"]
