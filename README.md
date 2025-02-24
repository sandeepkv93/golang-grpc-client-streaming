# 🧮 gRPC Client-Side Streaming Demo in Go! 🚀

This repository demonstrates a simple gRPC client-side streaming example in Go. 💻 It includes two microservices: a server that calculates the average of streamed numbers and a client that sends a stream of numbers to the server. 🔢

## ✨ Functionality ✨

This project showcases the power of gRPC client-side streaming, where the client sends a stream of messages to the server, and the server responds with a single message after the client has finished streaming. 🔄 This is particularly useful for scenarios like uploading large files, sending real-time sensor data, or processing continuous data streams. 📤

- **Server:** 🖥️
    - Implements a gRPC server that listens on port 50051. 🎧
    - Exposes a `ComputeAverage` RPC that uses client-side streaming. 🔢
    - Receives a stream of numbers from the client. 📥
    - Calculates the average of all received numbers. 🧮
    - Sends the calculated average back to the client. 📤
- **Client:** 💻
    - Establishes a gRPC connection to the server. 🤝
    - Sends a stream of numbers (from 1 to 10) to the server every second. 📤
    - Receives the average of the numbers from the server. 📥
    - Prints the result to the console. 🖨️

## 📂 Project Structure 📂

```bash
grpc-client-streaming/
├── proto/
│   └── average.proto
├── server/
│   └── server.go
├── client/
│   └── client.go
└── go.mod
```

- **`proto/`:** Contains the Protocol Buffer definition (`average.proto`) for the service. 📜
- **`server/`:** Contains the Go code for the gRPC server (`server.go`). ⚙️
- **`client/`:** Contains the Go code for the gRPC client (`client.go`). 💻
- **`go.mod`:** Go module definition file. 📦

## 🚀 Running the Example 🚀

1. **Generate gRPC code:**

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/average.proto
```

2. **Run the server:**

```bash
go run server/server.go
```

3. **Run the client (in a separate terminal):**

```bash
go run client/client.go
```

## 📦 Dependencies 📦

- `google.golang.org/grpc`

## 🔑 Key Concepts 🔑

- **gRPC:** A high-performance, open-source universal RPC framework. [cite: https]
- **Protocol Buffers:** A language-neutral mechanism for serializing structured data. [cite: https]
- **Client-Side Streaming:** A gRPC communication model where the client sends a stream of messages to the server, and the server responds with a single message after the client has finished streaming. [cite: https]

## 💡 Use Cases 💡

Client-side streaming is ideal for applications that require:

- Uploading large files or data streams. 📤
- Sending real-time sensor data. 🌡️
- Processing continuous data streams. 〰️

## 📝 Notes 📝

- This example uses an insecure connection for simplicity. In a production environment, you should use secure connections with TLS. 🔒
- The client sends numbers from 1 to 10, but you can modify this to send a different range of numbers or even read data from a file or another source. 🔢
- The server calculates the average of the numbers, but you can modify this to perform any other calculation or processing on the streamed data. 🧮