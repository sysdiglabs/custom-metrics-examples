FROM openjdk:8-jdk-slim
WORKDIR /app
COPY target/java-prom-example-1.0-SNAPSHOT-jar-with-dependencies.jar /app/app.jar
CMD ["java", "-cp", "/app/app.jar" , "Main"]