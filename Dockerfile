# Use the Maven base image for the final stage
FROM maven:3.9.7

# Set the working directory inside the container
#WORKDIR /app

# Copy the built Go executable from the builder stage to the final stage
COPY drone-maven-plugin /bin/

# Set the entrypoint to the Go executable
ENTRYPOINT ["/bin/drone-maven-plugin"]
