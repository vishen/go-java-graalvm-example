This repo provides a basic example of calling Java code from Go using GraalVM to compile
the Java code to a shared library, and using CGO to call into the shared library at runtime.

```
$ docker build -t go-java-example . 
$ docker run -it go-java-example
HOME=/root
Number of entries: 1
```

# Resources

- https://www.graalvm.org/22.2/reference-manual/native-image/guides/build-native-shared-library/#run-a-demo
- https://github.com/lxwagn/using-go-with-c-libraries

