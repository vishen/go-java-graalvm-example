FROM ghcr.io/graalvm/graalvm-ce:ol7-java17-22.3.0-b2 as builder

RUN gu install native-image llvm-toolchain

COPY . .

RUN javac LibEnvMap.java && \
	native-image -H:Name=libenvmap --shared && \
	/opt/graalvm-ce-java17-22.3.0/lib/llvm/bin/clang -I ./ -L ./ -l envmap -Wl,-rpath ./ -o main main.c

FROM golang:1.19.4-bullseye

COPY --from=builder /app/* .

CMD go run main.go
