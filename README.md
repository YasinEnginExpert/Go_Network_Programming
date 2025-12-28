# Go Network Programming & Network Automation

This repository is a curated collection of **educational examples, experiments, and mini-projects**
focused on **network programming and network automation using Go**.

The goal of this repository is not only to demonstrate how networking works in Go, but also to help
build a **strong conceptual foundation** in computer networking by combining:

- Networking theory (TCP/IP, OSI, RFCs)
- Practical Go implementations
- Real-world use cases inspired by production systems

This repository is especially useful for:
- Network engineers (CCNA / CCNP level and beyond)
- Backend and systems developers
- Students studying computer networks
- Engineers interested in **network automation and protocol-level programming**

---

## 🧭 Scope & Topics Covered

The examples in this repository cover a wide range of networking concepts, including but not limited to:

### 🔹 Core Networking Protocols
- TCP and UDP client–server models
- Echo servers and multi-client systems
- Daytime protocol (RFC 867)
- FTP-style client/server communication
- DNS-style lookup mechanisms
- ICMP and basic diagnostics

### 🔹 Internet & Link Layer Concepts
- IP addressing (IPv4 & IPv6)
- CIDR, subnetting, and prefix calculations
- ARP and Neighbor Discovery concepts
- Ethernet and link-layer fundamentals
- Raw packet and low-level networking examples

### 🔹 Modern Go Networking
- Usage of the `net` and `net/netip` packages
- Binary representation of IP addresses
- Serialization and deserialization
- JSON, Protobuf, and text-based protocols
- gRPC (unary and streaming)

### 🔹 Concurrency & Performance
- Goroutines and channels in network servers
- Multithreaded / concurrent echo servers
- Timeouts, keep-alive mechanisms
- Pipeline and concurrency patterns

### 🔹 Security & Cryptography (Introductory)
- Symmetric hashing and Merkle trees
- X.509 certificates
- Secure communication basics

### 🔹 Network Automation Foundations
- Programmatic interaction with network components
- Protocol-aware automation examples
- Concepts applicable to SDN, cloud networking, and orchestration systems

---

## 📁 Repository Structure

Each directory focuses on a **specific protocol, concept, or pattern**, and is designed to be
**self-contained** and easy to explore.

Examples include:
- `InternetLayer/` – IP addressing, CIDR, and routing-related logic
- `LinkLayer/` – Ethernet, ARP, and low-level networking
- `ICMP/` – Diagnostic protocols
- `GRPC_Client/`, `GRPC_Server/`, `GRPC_STREAM/` – gRPC communication patterns
- `MultithreadedEcho/` – Concurrent server design
- `JSON/`, `ProtocolBuffers/` – Data serialization
- `IPv4_router/` – Routing logic and packet forwarding concepts

---

## ⚙️ Requirements

- **Go 1.20 or newer**
- A basic understanding of networking concepts (TCP/IP, IP addressing)
- Linux or UNIX-like systems recommended for low-level examples

---

## 📚 References & Further Reading

The concepts and examples in this repository are inspired by and aligned with the following resources:

### 📖 Books
- **Jan Newmarch, Ronald Petty** – *Network Programming with Go*  
  Springer, 2021  
  ISBN: 978-1-4842-6874-8  
  A comprehensive and practical guide to Go’s networking capabilities.

### 📘 Official Go Documentation
- Go `net` package  
  https://pkg.go.dev/net
- Effective Go  
  https://go.dev/doc/effective_go

### 📜 RFC Standards
- RFC 791 – Internet Protocol (IP)
- RFC 793 – Transmission Control Protocol (TCP)
- RFC 867 – Daytime Protocol

### 📝 Go Blog
- Concurrency and Networking Patterns  
  https://blog.golang.org/pipelines

---

## 🎯 Philosophy

This repository is intentionally **educational rather than framework-driven**.

Instead of hiding networking details behind abstractions, the examples aim to:
- Show **how protocols actually work**
- Bridge the gap between **theory and implementation**
- Prepare readers for **real-world networking and automation tasks**

If you understand the code here, you will better understand:
- How operating systems handle networking
- How modern distributed systems communicate
- How network automation tools are built internally

---

## 🚀 Getting Started

Clone the repository and explore individual directories:

```bash
git clone https://github.com/your-username/go-network-programming.git
cd go-network-programming
go run ./InternetLayer
