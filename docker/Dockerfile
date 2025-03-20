FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    openjdk-8-jdk wget unzip software-properties-common \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
    
RUN add-apt-repository ppa:deadsnakes/ppa && \
    apt-get update && \
    apt-get install -y python3.10 python3.10-venv python3.10-distutils


ARG GRADLE_VERSION=8.0
RUN wget https://services.gradle.org/distributions/gradle-${GRADLE_VERSION}-bin.zip \
    && unzip gradle-${GRADLE_VERSION}-bin.zip -d /opt/gradle \
    && ln -s /opt/gradle/gradle-${GRADLE_VERSION}/bin/gradle /usr/bin/gradle \
    && rm gradle-${GRADLE_VERSION}-bin.zip

ARG SQLITE_JDBC_VERSION=3.36.0.3
RUN wget https://repo1.maven.org/maven2/org/xerial/sqlite-jdbc/${SQLITE_JDBC_VERSION}/sqlite-jdbc-${SQLITE_JDBC_VERSION}.jar -P /usr/lib/sqlite

COPY build.gradle /app/build.gradle
COPY src /app/src

WORKDIR /app

CMD ["gradle", "run"]

