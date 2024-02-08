# Go Proto: Protocol Buffers Implementation in Go

The **Go Proto** project aims to demonstrate how to work with protocol buffers (protobufs) in the Go programming language. By following this example, you'll learn how to:

1. Define message formats in a `.proto` file.
2. Use the **protoc** (protocol buffer compiler) tool.
3. Utilize the Go protocol buffer API to read and write messages.

## Problem Domain

Imagine you're building an application that manages an **address book**. Each entry in the address book contains a person's name, ID, email address, and contact phone number. How do you efficiently serialize and retrieve structured data like this?

### Why Protocol Buffers?

- **Flexible and Efficient**: Protocol buffers provide a flexible and efficient solution. You define the data structure in a `.proto` file, and the protocol buffer compiler generates code that handles encoding and parsing with an efficient binary format.
- **Type-Safe**: The generated code includes getters and setters for the fields, ensuring type safety.
- **Forward-Compatible**: Protocol buffers support extending the format over time, allowing backward compatibility with older data.

## Getting Started

1. Create a message type in `api/v1/activity.proto` (similar to our address book example):

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

## Example Code

Our example includes command-line applications for managing an address book data file encoded using protocol buffers. For more detailed reference information, explore the [Protocol Buffers Documentation](https://protobuf.dev/getting-started/gotutorial/) and the [Go Generated Code Guide](https://protobuf.dev/reference/go/go-generated/).

Happy coding! 🚀
