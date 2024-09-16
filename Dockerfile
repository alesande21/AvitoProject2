FROM golang:1.23.1-alpine3.20 AS build
COPY . /home/src
WORKDIR /home/src
RUN apk add make && make build

FROM alpine:3.20
EXPOSE 8080
WORKDIR /app
COPY --from=build /home/src/bin/app /app/avito_service

ENTRYPOINT [ "/bin/sh", "-c", "/app/avito_service" ]








#FROM gradle:4.7.0-jdk8-alpine AS build
#COPY --chown=gradle:gradle . /home/gradle/src
#WORKDIR /home/gradle/src
#RUN gradle build --no-daemon
#
#FROM openjdk:8-jre-slim
#
#EXPOSE 8080
#
#RUN mkdir /app
#
#COPY --from=build /home/gradle/src/build/libs/*.jar /app/spring-boot-application.jar
#
#ENTRYPOINT ["java", "-XX:+UnlockExperimentalVMOptions", "-XX:+UseCGroupMemoryLimitForHeap", "-Djava.security.egd=file:/dev/./urandom","-jar","/app/spring-boot-application.jar"]
#
