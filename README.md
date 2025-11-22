# Distributed Key-Value Store

A scalable, high-performance distributed key-value storage system engineered for modern applications. Built on BoltDB's robust foundation, this implementation showcases enterprise-grade distributed systems architecture with intelligent data partitioning, seamless replication, and efficient request routing.

## Key Features

- **Intelligent Sharding**: Advanced consistent hashing algorithm for optimal data distribution across multiple nodes
- **RESTful API**: Clean, intuitive HTTP endpoints for seamless key-value operations
- **Persistent Storage**: BoltDB-powered embedded storage ensuring data durability and ACID compliance
- **Dynamic Shard Management**: Flexible configuration system supporting dynamic cluster topology
- **Fault Tolerance**: Built-in redundancy and failover mechanisms for enterprise reliability
- **Horizontal Scalability**: Easy cluster expansion with automatic load balancing

## System Architecture

Our distributed architecture implements a sophisticated sharding strategy:

- **Consistent Hashing**: Keys are intelligently distributed across shards using a proven hashing algorithm
- **Independent Shards**: Each shard operates as a self-contained unit with dedicated BoltDB storage
- **Smart Routing**: HTTP requests are automatically directed to the appropriate shard based on key hash
- **Configuration-Driven**: TOML-based configuration enables flexible deployment scenarios

## API Reference

### Retrieve Operations
Fetch stored values using simple HTTP GET requests:
```bash
curl -Lkv 'http://localhost:8081/retrieve?key=mykey'
```

### Store Operations
Store key-value pairs with atomic write operations:
```bash
curl -Lkv 'http://localhost:8081/store?key=mykey&value=myvalue'
```

## Configuration Management

The system leverages TOML configuration files for flexible cluster management:

```toml
[[shards]]
name = "shard-1"
idx = 0
address = "localhost:8081"

[[shards]]
name = "shard-2"
idx = 1
address = "localhost:8082"
```

## Quick Start Guide

### Prerequisites
- Go 1.19 or higher
- Git

### Installation & Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/AllaDheerajKumar/KeyValueStore.git
   cd KeyValueStore
   ```

2. **Build the application**:
   ```bash
   go build -o kvstore
   ```

3. **Configure your shards**:
   - Edit `shard-config.toml` with your desired cluster topology
   - Ensure each shard has a unique address and index

4. **Launch the cluster**:
   ```bash
   ./start.sh
   ```

5. **Verify deployment**:
   ```bash
   curl http://localhost:8081/retrieve?key=test
   ```

## Performance Characteristics

- **Throughput**: Optimized for high-concurrency read/write operations
- **Latency**: Sub-millisecond response times for in-memory operations
- **Durability**: ACID-compliant transactions with persistent storage
- **Scalability**: Linear performance scaling with additional shards

## Development & Contributing

**Lead Developer**: Dheeraj Kumar Alla

We welcome contributions! Please feel free to submit issues, feature requests, or pull requests to help improve this distributed storage solution.

### Development Setup
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request


