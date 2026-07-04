# ADR-001: Use JSON File Storage

## Status
Accepted

## Context
BasketForm-AI needs to persist user accounts, video metadata, friend relationships, shared results, and analysis results. We need a storage solution that is simple to implement and doesn't require external database setup.

Options considered: Relational database (PostgreSQL, SQLite), NoSQL database (MongoDB), JSON files on disk.

## Decision
Use JSON files stored on the local filesystem. Each entity type gets its own directory: data/users/, data/videos/, data/friends/, data/shared/, results/.

## Consequences and Tradeoffs

**Positive:** Zero dependencies, simple debugging (human-readable files), easy backup, fast development iteration.

**Negative:** No ACID transactions, no indexing (full scan required), not scalable beyond ~10 concurrent users, no referential integrity enforcement.

## Quality Requirements Addressed
- QR-001 (Performance): Sufficient for small user base
- QR-003 (Usability): Simplifies development, faster feature iteration
- QR-005 (Concurrency): Limited by mutex-based locking; acceptable for MVP
