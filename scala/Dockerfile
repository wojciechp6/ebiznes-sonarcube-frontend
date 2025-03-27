FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    openjdk-11-jdk \
    curl \
    wget \
    git \
    unzip \
    zip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*


RUN curl -s "https://get.sdkman.io" | bash
RUN bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && sdk install java 8.0.362-amzn"
RUN bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && sdk install sbt 1.10.9"
RUN bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && sdk install scala 3.3.5"

ENV PATH=$PATH:/root/.sdkman/candidates/sbt/1.10.9/bin
ENV PATH=$PATH:/root/.sdkman/candidates/java/8.0.362-amzn/bin/
ENV PATH=$PATH:/root/.sdkman/candidates/scala/3.3.5/bin


WORKDIR /app
COPY /ebiznes-scala /app

RUN sbt compile

EXPOSE 9000

CMD ["sbt", "run"]

