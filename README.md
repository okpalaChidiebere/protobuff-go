# Protocol Buffers

In this Project, i learn how to use Protocol Buffers and why we might consider using it while building our backend. Protocol Buffers is an alternative to JSON or XML. I learned how to encode and Decode date Data in Protocol Buffers; we encode data in JSON as well and compared the bytes size. Full tutorial Here. in [NodeJS](https://www.youtube.com/watch?v=46O73On0gyI&t=1128s) and Go tutorial [here](https://www.youtube.com/watch?v=qWN69yfRsVs)
In another Project i will learn about gRPC which used protoBuffers

## Cons of Protocol Buffers

- Protocol Buffers are mainly useful when you want so send data between micro services. Whether your your microservices written in various programming languages or not they can share thesame Data structure. With Protocol buffers you MUST define the structure of your payload or schema in advance which can then be encoded or decoded using the protobufers; it is not like JSON where the schema can be flexible

## Pros of Protocol Buffers

-It is very very compact in terms of bytes used in compared to JSON. In this project we learened that the binary of the proto and file size is half in size in compared to JSON binary and size.

- we ran `ls -lh` to check the bytes size of the file `employeesbinary` and `jsondata.json`

## Protocol Buffers Compilers

- We need this compilers to generate a class or interface in any language from our proto so that we can encode or decode data in the backend

## WorkSpace set up to use use Proto

- Make sure that you have the protoc compiler available your system. I confirm that by running `protoc --version` in terminal.
- On my mac, i used homebrew to install the instal the compiler following the instructions [here](https://grpc.io/docs/protoc-installation/)
- If i wanted to set up a CI/CD process using travis CI, so whenever i add a new proto file in a github repo to generate the protos for me on another github repo or on thesame repo i can read [this](https://docs.travis-ci.com/user/installing-dependencies/#installing-packages-on-macos) and [this](https://docs.travis-ci.com/user/installing-dependencies/#installing-projects-from-source)
- **NOTE:** I had to install only the protobuf package and not with the gRPC package because i was only learning about the protobuf. See [here](https://grpc.io/docs/languages/go/quickstart/)

## Additiona links

- [https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable](https://stackoverflow.com/questions/57700860/error-protoc-gen-go-program-not-found-or-is-not-executable)
- [https://stackoverflow.com/questions/61666805/correct-format-of-protoc-go-package](https://stackoverflow.com/questions/61666805/correct-format-of-protoc-go-package)
