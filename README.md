# ObjectDB

> Ultra-fast object-oriented database engine written in Go  
> No SQL. Native objects, zero-copy reads, and lock-free indexes.

**Status:** Design / pre-implementation. No runnable code exists yet — this document defines the architecture and roadmap.

## Table of Contents

- [Overview](#overview)
- [Core Philosophy](#core-philosophy)
- [Architecture](#architecture)
- [Object Model](#object-model)
- [Storage Engine](#storage-engine)
- [Distributed Cluster](#distributed-cluster)
- [Security](#security)
- [Query Optimization](#query-optimization)
- [Indexes](#indexes)
- [Vector Search](#vector-search)
- [Full Text Search](#full-text-search)
- [Caching](#caching)
- [Observability](#observability)
- [Performance Targets](#performance-targets)
- [Getting Started](#getting-started)
- [License](#license)

## Overview

ObjectDB is a next-generation high-performance object database designed for workloads where relational models become inefficient, overly complex, or too slow.

Instead of tables, joins, and SQL, ObjectDB stores native typed objects directly in optimized binary pages.

**Designed for:**

- ultra-low latency
- distributed systems
- AI infrastructure
- event-driven architectures
- large-scale object graphs

**Optimized for:**

- real-time systems
- game servers
- telemetry
- financial streams
- AI memory systems
- event sourcing
- graph-like object relations
- distributed state engines
- document/object workloads
- large-scale caching
- embedded databases
- microservices
- high-frequency writes
- immutable historical data

**Transport:** binary protocol, [gRPC](https://grpc.io), [HTTP/3](https://httpwg.org/specs/rfc9114.html), [WebSocket](https://datatracker.ietf.org/doc/html/rfc6455), [QUIC](https://quicwg.org)

**Client SDKs:** Go, Rust, C++, Java, Python, JavaScript - and any language via the binary protocol.

## Core Philosophy

### Native Objects Instead of Tables

Traditional SQL databases optimize for:

- tables
- joins
- SQL parsing
- relational normalization

ObjectDB optimizes for:

- object graphs
- nested structures
- direct references
- immutable snapshots
- append-only storage
- typed memory layouts

## Architecture

[![](https://mermaid.ink/img/pako:eNqFlO1uozgUhm_Fopr9s6RNC80HsxqJkO-GJA1sKw3ZHw6YhC3YyJhp06rSXste2lzJnNgpIatKyx9zsM_j9_U55k0LWUQ0S4tT9hzuMBfI768pgscOnDQhVCCvf1f8seFX30YM_fznX7QqCyFfpp4clnuxY_SvNT3moUbjG-oFcyKeGX9C9nKCZnhPuGTc_zlxZNZ2tVQvY99fXhlVek-mO4Fdih3snoRYJIyi35BHwpInYi8xq56tsv2ZEjF99JHPnggtKpIjSf3AYZSSUFKWjKXIxRRvj3I8UhSHCQeHOyJBbpmKJE_JS0K3FaovUYPA55gWWLEGdJtQIinug6PUPNozOXoU58WOiZOYgSQMgwltuCRjfI96ZRwTLiVJyJgJtARhhTpk9lwT9QjGyTGjQg4lchTclwRwNTne_QxIvAD6IXmRiyRLXo_RMsVwHCfKSFLGICySnmu16v3uc6IEzDxXjr1EZDhHcjE5uRtLyCRY2e7HdqpsJ5gdYZj4QWqulpzERIQ7GTgsy7mqRoWdSOw0sIs9DZGPiyfkQXJUph8Kcfi05aykERqmZaFII6cCqlJVvKnk3QWLzd_QEMgTjMOBnx0d2WbQdaoGp2p4no83ac3wnSTNghXJ01OLQqcV0IGlui4rHKt74sG9ig5HewiGOEnZj1oBZhLlBv0E7F0hB3wIqNUc7qXizB9cclTRlyOsFDzZlIKAba8CffmCXEYTMAV7qU_zwCWwNCxAnE9SkkGkLtCSMwh2pFQGobHDD4UwFSepan8JQY3LQ_efRaOzaHIWzWqKloTHjGeYhgThMAQNHIPCQi1YBPV5uzavqjFxlePvhLOGw_K9jB6geGDyFewPXuCvcKyxBJ6pW1TqKj2e2KfV8YQpLoo-iVHIOEFgOrUurslN19joIUsZty7iONbhsOHHYl10TWxsOsew8ZxEYmdd5y9f_8PKwdGR1WyZ5HOWYUZGt_u_rITGHB9hxqbZbpmfwcJmx4zDz2A1HLL1nu7ofX2gj_SxfqfPdFf6ru-JhvpEn-oLaeJsYq60fNV0bcuTSLMEL4muZQRqdwi1t8PqtQYtlZG1ZsFrhPnTWlvTd8jJMf3OWPaRBld2u9OsGKcFRGUeYUH6Cd5yfFpC4B_DHbjbQrO6LYnQrDftRbOuO-3Lltlu3XTaHfO63W3r2l6zGmbzsn3TaZot-HZjtAzjXdde5abNy65hGF3zFnJat7ed1u37L0EMUTw?type=png)](https://mermaid.live/edit#pako:eNqFlO1uozgUhm_Fopr9s6RNC80HsxqJkO-GJA1sKw3ZHw6YhC3YyJhp06rSXste2lzJnNgpIatKyx9zsM_j9_U55k0LWUQ0S4tT9hzuMBfI768pgscOnDQhVCCvf1f8seFX30YM_fznX7QqCyFfpp4clnuxY_SvNT3moUbjG-oFcyKeGX9C9nKCZnhPuGTc_zlxZNZ2tVQvY99fXhlVek-mO4Fdih3snoRYJIyi35BHwpInYi8xq56tsv2ZEjF99JHPnggtKpIjSf3AYZSSUFKWjKXIxRRvj3I8UhSHCQeHOyJBbpmKJE_JS0K3FaovUYPA55gWWLEGdJtQIinug6PUPNozOXoU58WOiZOYgSQMgwltuCRjfI96ZRwTLiVJyJgJtARhhTpk9lwT9QjGyTGjQg4lchTclwRwNTne_QxIvAD6IXmRiyRLXo_RMsVwHCfKSFLGICySnmu16v3uc6IEzDxXjr1EZDhHcjE5uRtLyCRY2e7HdqpsJ5gdYZj4QWqulpzERIQ7GTgsy7mqRoWdSOw0sIs9DZGPiyfkQXJUph8Kcfi05aykERqmZaFII6cCqlJVvKnk3QWLzd_QEMgTjMOBnx0d2WbQdaoGp2p4no83ac3wnSTNghXJ01OLQqcV0IGlui4rHKt74sG9ig5HewiGOEnZj1oBZhLlBv0E7F0hB3wIqNUc7qXizB9cclTRlyOsFDzZlIKAba8CffmCXEYTMAV7qU_zwCWwNCxAnE9SkkGkLtCSMwh2pFQGobHDD4UwFSepan8JQY3LQ_efRaOzaHIWzWqKloTHjGeYhgThMAQNHIPCQi1YBPV5uzavqjFxlePvhLOGw_K9jB6geGDyFewPXuCvcKyxBJ6pW1TqKj2e2KfV8YQpLoo-iVHIOEFgOrUurslN19joIUsZty7iONbhsOHHYl10TWxsOsew8ZxEYmdd5y9f_8PKwdGR1WyZ5HOWYUZGt_u_rITGHB9hxqbZbpmfwcJmx4zDz2A1HLL1nu7ofX2gj_SxfqfPdFf6ru-JhvpEn-oLaeJsYq60fNV0bcuTSLMEL4muZQRqdwi1t8PqtQYtlZG1ZsFrhPnTWlvTd8jJMf3OWPaRBld2u9OsGKcFRGUeYUH6Cd5yfFpC4B_DHbjbQrO6LYnQrDftRbOuO-3Lltlu3XTaHfO63W3r2l6zGmbzsn3TaZot-HZjtAzjXdde5abNy65hGF3zFnJat7ed1u37L0EMUTw)

### Object Storage Layout

Each object is stored as:

| Header | TypeID | Version | Flags | Binary Object Data |

### Binary Encoding

Custom zero-copy binary protocol with:

- fixed-width primitives
- pointer-free storage
- [SIMD](https://en.wikipedia.org/wiki/Single_instruction,_multiple_data)-friendly layout
- endian-safe encoding
- arena allocation
- optional compression

Supports:

- strings
- arrays
- maps
- nested objects
- vectors
- blobs
- references
- enums
- generics

### Memory Model

Arena-based allocation using:

- arenas
- slab allocators
- object pools
- bump allocators

Minimizes:

- GC pressure
- heap fragmentation
- allocations

Critical for Go performance.

### Native Network Protocol

#### ODBP (Object Database Protocol)

Designed for:

- low latency
- streaming objects
- multiplexed requests
- binary efficiency
- zero-copy decoding

Features:

- framed binary packets
- request pipelining
- multiplexing
- compression
- TLS support
- streaming replication
- bidirectional communication

Traditional APIs waste performance on JSON serialization, HTTP overhead, text parsing, and excessive allocations. ODBP avoids this with binary packets, schema-aware serialization, compact object encoding, and direct memory decoding.

#### Transport Protocols

**[QUIC](https://quicwg.org)** - primary transport:

- low latency
- multiplexing
- connection migration
- built-in encryption

## Object Model

Objects are stored in native binary layout:

| Header | TypeID | Version | Flags | Object Data |

Supports:

- nested objects
- arrays
- maps
- vectors
- references
- blobs
- enums

### Schema System

Dynamic + Static Schemas

Supports:

- strongly typed schemas
- schema evolution
- runtime schemas
- generated codecs

### Reflection-Free Serialization

Uses generated codecs instead of reflection.

Benefits:

- faster encoding
- lower allocations
- predictable performance

## Storage Engine

Hybrid LSM + Page-Oriented Architecture combining:

- [LSM-tree](https://en.wikipedia.org/wiki/Log-structured_merge-tree) write optimization
- page-oriented memory locality
- append-only immutable segments
- arena-backed page cache

This gives:

- extremely fast writes
- crash safety
- high compression
- efficient snapshots
- predictable latency

### WAL

[Write Ahead Log](https://en.wikipedia.org/wiki/Write-ahead_logging) (WAL) guarantees durability.

Features:

- checksums
- crash recovery
- fsync batching
- compression

### MVCC

[Multi-Version Concurrency Control](https://en.wikipedia.org/wiki/Multiversion_concurrency_control) (MVCC) provides:

- lock-free reads
- snapshot isolation
- historical queries
- concurrent writers

## Distributed Cluster

Horizontal scalability with:

- sharding
- replication
- failover
- automatic balancing

### Consensus

Uses [Raft](https://raft.github.io) consensus:

- distributed metadata
- leader election

### Replication

Supports:

- async replication
- sync replication
- snapshot replication
- incremental replication

## Security

### Authentication

Supports: [JWT](https://jwt.io), API keys, [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control), [OAuth2](https://oauth.net/2), [mTLS](https://en.wikipedia.org/wiki/Mutual_authentication)

### Encryption

Supports: [TLS 1.3](https://datatracker.ietf.org/doc/html/rfc8446), encrypted WAL, encrypted snapshots, [AES-256](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard) at rest

## Query Optimization

Engine supports:

- index selection
- predicate pushdown
- parallel execution
- vectorized scans

## Indexes

Supports:

- [B+Tree](https://en.wikipedia.org/wiki/B%2B_tree)
- [Adaptive Radix Tree](https://en.wikipedia.org/wiki/Adaptive_radix_tree)
- Hash Index
- [Bitmap Index](https://en.wikipedia.org/wiki/Bitmap_index)
- [HNSW](https://arxiv.org/abs/1603.09320) Vector Index

## Vector Search

Built-in vector search for AI systems.

Supports:

- embeddings
- [cosine similarity](https://en.wikipedia.org/wiki/Cosine_similarity)
- nearest neighbors
- [ANN search](https://en.wikipedia.org/wiki/Nearest_neighbor_search)

## Full Text Search

Integrated search engine:

- [BM25](https://en.wikipedia.org/wiki/Okapi_BM25)
- stemming
- fuzzy search
- tokenization

## Caching

Multi-layer cache:

- page cache
- object cache
- query cache
- index cache

## Observability

Built-in:

- [Prometheus](https://prometheus.io) metrics
- [OpenTelemetry](https://opentelemetry.io) tracing
- slow query logs
- profiling
- cluster diagnostics

## Performance Targets

| Operation         | Target       |
| ----------------- | ------------ |
| Point Read        | < 10µs       |
| Insert            | millions/sec |
| Replication Lag   | milliseconds |
| Snapshot Creation | near instant |
| Object Streaming  | GB/s scale   |

## Getting Started

This project is in early design phase. No code, packages, or binaries are available yet.

Once implementation begins, this section will cover:

- Installing the ObjectDB server
- Starting your first database
- Connecting with a client SDK
- Inserting and querying objects

## License

[MIT](LICENSE)
