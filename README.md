# Go Network Programming & Network Automation

A comprehensive collection of educational examples, experiments, and mini-projects focused on network programming and network automation using the Go programming language.

---

## Overview

This repository serves as both a learning resource and a reference implementation for network programming concepts. The primary objectives are:

- Demonstrate practical networking implementations in Go
- Build a strong conceptual foundation in computer networking
- Bridge the gap between networking theory and real-world implementation
- Provide production-inspired code patterns for network automation

The content combines networking theory (TCP/IP, OSI model, RFC specifications) with hands-on Go implementations, making it suitable for:

- Network engineers (CCNA, CCNP level and beyond)
- Backend and systems developers
- Computer science students studying networks
- Engineers working on network automation and protocol-level programming

---

## Scope and Topics

### Core Networking Protocols

This repository covers fundamental networking protocols and their implementations:

- **TCP/UDP Client-Server Models**: Complete implementations of connection-oriented and connectionless communication patterns
- **Echo Servers**: Single and multi-client echo server implementations demonstrating basic socket programming
- **Daytime Protocol (RFC 867)**: RFC-compliant time server implementation
- **FTP-Style Communication**: File transfer protocol client and server examples
- **DNS Lookup Mechanisms**: Domain name resolution and service discovery patterns
- **ICMP Diagnostics**: Network diagnostic tools and ping implementations

### Internet and Link Layer

Low-level networking concepts and implementations:

- **IP Addressing**: IPv4 and IPv6 address handling, parsing, and manipulation
- **CIDR and Subnetting**: Network prefix calculations and subnet operations
- **ARP and Neighbor Discovery**: Address resolution protocol concepts
- **Ethernet Fundamentals**: Link-layer frame handling and MAC address operations
- **Raw Packet Processing**: Low-level packet construction and parsing

### Modern Go Networking

Leveraging Go's networking ecosystem:

- **net Package**: Core networking functionality including TCP, UDP, and Unix sockets
- **net/netip Package**: Modern IP address handling with improved performance
- **Binary Representations**: Efficient binary encoding of network data
- **Protocol Serialization**: JSON, Protocol Buffers, and custom text-based protocols
- **gRPC**: Unary and streaming RPC implementations

### Concurrency and Performance

High-performance networking patterns:

- **Goroutine-Based Servers**: Concurrent connection handling
- **Channel Communication**: Inter-goroutine messaging patterns
- **Multithreaded Echo Servers**: Scalable server architectures
- **Timeout Management**: Connection and operation timeout handling
- **Keep-Alive Mechanisms**: Long-lived connection management
- **Pipeline Patterns**: Staged concurrent processing

### Security and Cryptography

Introductory security concepts:

- **Symmetric Hashing**: Hash functions and their applications
- **Merkle Trees**: Cryptographic data structure implementations
- **X.509 Certificates**: Certificate parsing and validation
- **Asymmetric Cryptography**: Public key cryptography basics
- **Secure Communication**: TLS and encrypted channel fundamentals

### Network Automation

Programmable network management:

- **Closed-Loop Automation**: Automated network state management
- **Container-Based Labs**: Network simulation using Containerlab
- **Protocol-Aware Automation**: Intelligent network configuration tools
- **SDN Concepts**: Software-defined networking patterns

---

## Repository Structure

All code examples are organized within the `Go Codes/` directory. Each subdirectory is self-contained and focuses on a specific protocol, concept, or pattern.

### Transport and Application Layer

| Directory | Description |
|-----------|-------------|
| `Day Time Server/` | RFC 867 Daytime Protocol server implementation |
| `FTP client and server/` | File Transfer Protocol client and server examples |
| `Client and Server Directory Protocol/` | Directory listing service implementation |
| `Get Head Info/` | HTTP HEAD request handling and response parsing |
| `HTTP_APP/` | HTTP application server examples |
| `MultiProtocolServer/` | Server supporting multiple protocols simultaneously |

### gRPC and Modern RPC

| Directory | Description |
|-----------|-------------|
| `GRPC_Client/` | gRPC client implementation with various call patterns |
| `GRPC_Server/` | gRPC server with service definitions and handlers |
| `GRPC_STREAM/` | Server-side streaming RPC examples |
| `GRPC_STREAM_CLIENT/` | Client-side and bidirectional streaming examples |

### Concurrency and Performance

| Directory | Description |
|-----------|-------------|
| `Multithreadedecho/` | Concurrent echo server with goroutine pool |
| `Goroutine/` | Goroutine lifecycle and management patterns |
| `Channels/` | Channel-based communication examples |
| `Message passing/` | Inter-process and distributed messaging patterns |
| `Timeoutandkeepalive/` | Connection timeout and keep-alive implementations |

### Internet and Network Layer

| Directory | Description |
|-----------|-------------|
| `InternetLayer/` | IP addressing, CIDR calculations, and routing logic |
| `IPv4 router/` | IPv4 packet routing and forwarding simulation |
| `IP mask/` | Subnet mask operations and network calculations |
| `ICMP/` | ICMP message handling and ping implementation |
| `Lookup port/` | Service port resolution and discovery |

### Link Layer

| Directory | Description |
|-----------|-------------|
| `Linklayer/` | Ethernet frame handling, ARP, and MAC operations |

### Serialization and Data Encoding

| Directory | Description |
|-----------|-------------|
| `JSON/` | JSON encoding and decoding for network messages |
| `JSON Serialization/` | Advanced JSON marshaling techniques |
| `Protocolbuffers/` | Protocol Buffers schema definition and usage |
| `Gob/` | Go's native binary encoding format |
| `Gob Echo Client and Server/` | Gob-based network communication |
| `ASN.1/` | Abstract Syntax Notation One encoding |
| `ASN.1 Daylight Client and Server/` | ASN.1 protocol implementation |
| `Textproto/` | Text-based protocol parsing utilities |
| `Manuel Serialization/` | Custom binary serialization implementations |
| `Modern Serialization/` | Contemporary serialization approaches |

### Security and Cryptography

| Directory | Description |
|-----------|-------------|
| `Asymetrickey/` | RSA and elliptic curve cryptography examples |
| `SymmetricHashingMerkle/` | Hash functions and Merkle tree implementations |
| `X.509/` | X.509 certificate parsing and chain validation |

### Network Automation

| Directory | Description |
|-----------|-------------|
| `Closed-Loop Network Automation/` | Automated network state reconciliation |
| `Containerlab/` | Container-based network topology definitions |

### Additional Topics

| Directory | Description |
|-----------|-------------|
| `DCE File System/` | Distributed Computing Environment file operations |
| `Karakter Felsefesi/` | Character encoding and text representation |

---

## Requirements

### Software Requirements

- **Go**: Version 1.20 or newer (latest stable release recommended)
- **Protocol Buffers Compiler**: Required for gRPC examples (`protoc`)
- **Git**: For cloning and version control

### System Requirements

- Linux or UNIX-like operating system recommended for low-level networking examples
- Windows Subsystem for Linux (WSL2) supported for Windows users
- Root/Administrator privileges may be required for raw socket and ICMP examples

### Knowledge Prerequisites

- Basic understanding of TCP/IP networking concepts
- Familiarity with IP addressing and subnetting
- Understanding of client-server architecture
- Basic Go programming experience

---

## Getting Started

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/YasinEnginExpert/go-network-programming.git
cd go-network-programming
```

### Running Examples

Each directory contains standalone examples. Navigate to the desired directory and run:

```bash
# Internet Layer examples
cd "Go Codes/InternetLayer"
go run .

# gRPC Server
cd "Go Codes/GRPC_Server"
go run .

# Multithreaded Echo Server
cd "Go Codes/Multithreadedecho"
go run .
```

### Project Structure Navigation

```
go-network-programming/
├── Go Codes/
│   ├── InternetLayer/          # Start here for IP fundamentals
│   ├── Day Time Server/        # Simple protocol implementation
│   ├── Multithreadedecho/      # Concurrent server patterns
│   ├── GRPC_Server/            # Modern RPC examples
│   └── ...                     # Additional modules
└── README.md
```

---

## References and Further Reading

### Books

- **Network Programming with Go** by Jan Newmarch and Ronald Petty  
  Springer, 2021 | ISBN: 978-1-4842-6874-8  
  A comprehensive guide to Go's networking capabilities with practical examples.

### Official Documentation

- [Go net Package](https://pkg.go.dev/net) — Standard library networking documentation
- [Go net/netip Package](https://pkg.go.dev/net/netip) — Modern IP address handling
- [Effective Go](https://go.dev/doc/effective_go) — Go programming best practices
- [gRPC-Go Documentation](https://grpc.io/docs/languages/go/) — gRPC implementation guide

### RFC Specifications

- [RFC 791](https://www.rfc-editor.org/rfc/rfc791) — Internet Protocol (IP)
- [RFC 793](https://www.rfc-editor.org/rfc/rfc793) — Transmission Control Protocol (TCP)
- [RFC 768](https://www.rfc-editor.org/rfc/rfc768) — User Datagram Protocol (UDP)
- [RFC 867](https://www.rfc-editor.org/rfc/rfc867) — Daytime Protocol
- [RFC 826](https://www.rfc-editor.org/rfc/rfc826) — Address Resolution Protocol (ARP)
- [RFC 792](https://www.rfc-editor.org/rfc/rfc792) — Internet Control Message Protocol (ICMP)

### Additional Resources

- [Go Blog: Pipelines and Cancellation](https://blog.golang.org/pipelines) — Concurrency patterns
- [Go Blog: Context](https://blog.golang.org/context) — Request-scoped values and cancellation
- [Protocol Buffers Documentation](https://protobuf.dev/) — Serialization format guide

---

## Design Philosophy

This repository follows an **educational-first approach** rather than a framework-driven design. The implementation choices prioritize:

### Transparency Over Abstraction

Instead of hiding networking complexity behind high-level abstractions, examples expose the underlying mechanisms. This approach helps developers understand:

- How operating systems handle network operations
- The actual bytes transmitted over the wire
- Protocol state machines and handshakes

### Theory-Implementation Bridge

Each example connects theoretical networking concepts to practical code:

- RFC specifications are referenced and implemented
- OSI/TCP-IP layer concepts are demonstrated concretely
- Real-world use cases inform example design

### Production-Ready Patterns

While educational, the code follows production best practices:

- Proper error handling and propagation
- Graceful shutdown and resource cleanup
- Concurrent-safe implementations
- Configurable timeouts and retries

### Progression Path

Examples are organized to support progressive learning:

1. Start with basic socket operations (`InternetLayer/`)
2. Progress to protocol implementations (`Day Time Server/`)
3. Explore concurrent patterns (`Multithreadedecho/`)
4. Advance to modern RPC (`GRPC_Server/`)
5. Apply to automation (`Closed-Loop Network Automation/`)

---

## Contributing

Contributions are welcome. Please ensure that new examples:

- Follow the existing code style and organization
- Include appropriate comments and documentation
- Are self-contained within their directory
- Reference relevant RFCs or specifications where applicable

---

## License

This project is available for educational purposes. See the repository for specific license terms.
