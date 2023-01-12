FROM ghcr.io/graalvm/graalvm-ce:ol7-java17-22.3.0-b2 as builder

RUN gu install native-image llvm-toolchain

COPY libs/*.jar *.java .

ENV JAVA_CP=".:pw-swift-core-SRU2021-9.2.11.jar:commons-lang3-3.12.0.jar:gson-2.8.9.jar"

RUN javac MT103Parser.java -cp $JAVA_CP #&& \
	#native-image -H:Name=libmt103parser -cp $JAVA_CP --shared

# Run this in the docker container and it prints the whole things
# RUN java -cp $JAVA_CP MT103Parser

#COPY . .
#
#FROM golang:1.19.4-bullseye
#
#COPY --from=builder /app/* .
#
#CMD go run main.go
