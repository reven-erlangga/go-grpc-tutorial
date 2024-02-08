# Golang gRPC Example

This repository demonstrates how to use **gRPC with Golang**. We'll focus on the **protoc tool** and building both a gRPC client and server. If you're curious about gRPC, its benefits, and how to write example code, you're in the right place!

## Table of Contents

- [Introduction](#introduction)
- [Why gRPC?](#why-grpc)
- [Protocol Buffers](#protocol-buffers)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Example](#example)

## Introduction

As an experienced developer, I'm learning Golang by building an **activity tracker**. In a previous step, I added **SQLite persistence**. Now, I'll be porting everything to gRPC.

## Why gRPC?

While JSON-based REST services are great for client-side JavaScript or when dealing with many clients, gRPC offers several advantages:

1. **Binary format**: gRPC uses a binary format, making it quicker to send data.
2. **Efficient serialization**: It's faster to serialize and deserialize compared to JSON.
3. **Better types**: gRPC provides better type support than JSON.
4. **Code generation**: The `protoc` tool generates code from `.proto` files.

## Protocol Buffers

One of the significant advantages of gRPC is **protocol buffers (protobufs)**. These allow you to encode message semantics in a parsable form shared between client and server. Protobufs are platform-neutral, offer fast serialization, and support schema migration.

## Getting Started

1. Create a message type in `api/v1/activity.proto`:

   ```proto
   syntax = "proto3";
   package api.v1;

   import "google/protobuf/timestamp.proto";

   message Activity {
       int32 id = 1;
       google.protobuf.Timestamp time = 2;
       string description = 3;
   }
   ```

2. Install the protobuf compiler:

   ```bash
   brew install protobuf
   ```

3. Generate Go structs for the message type:

   ```bash
   protoc activity-log/api/v1/*.proto \
       --go_out=. \
       --go_opt=paths=source_relative \
       --proto_path=.
   ```

## Usage

Feel free to explore the full code in the [GitHub repository](https://github.com/reven-erlangga/go-grpc-tutorial).

## Contributing

Contributions are welcome! Please follow the [code of conduct](https://github.com/reven-erlangga/go-grpc-tutorial/blob/main/CODE_OF_CONDUCT.md).

## License

This project is licensed under the [Apache-2.0 License](https://github.com/reven-erlangga/go-grpc-tutorial/blob/main/LICENSE).

## Example

So, this project have some example to gRPC implementation, looks like :

1. **Go Proto**: example project to implementation proto using go language.
