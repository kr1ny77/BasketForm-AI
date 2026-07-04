# ADR-001: Use JSON File Storage

## Status
Accepted

## Context
BasketForm-AI needs to persist user accounts, video metadata, friend relationships, shared results, and analysis results. We need a storage solution that is simple to implement and doesn't require external database setup.

<<<<<<< HEAD
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
=======
Options considered:
- Relational database (PostgreSQL, SQLite)
- NoSQL database (MongoDB)
- JSON files on disk

## Decision
Use JSON files stored on the local filesystem. Each entity type gets its own directory:
- `data/users/{id}.json` — user accounts
- `data/videos/{id}.json` — video metadata
- `data/friends/{id}.json` — friend requests
- `data/shared/{id}.json` — shared results
- `results/{id}.json` — analysis results

All operations use file I/O with a mutex for thread safety.

## Consequences and Tradeoffs

**Positive:**
- Zero dependencies — no database server to install/maintain
- Simple debugging — files are human-readable
- Easy backup — just copy the directory
- Fast development iteration

**Negative:**
- No ACID transactions — concurrent writes can corrupt data
- No indexing — full scan required for all queries
- Not scalable beyond ~10 concurrent users
- No referential integrity enforcement

## Quality Requirements Addressed
- **QR-001 (Performance)**: Sufficient for small user base; file reads are fast for small JSON files
- **QR-003 (Usability)**: Simplifies development, faster feature iteration
- **QR-005 (Concurrency)**: Limited by mutex-based locking; acceptable for MVP
>>>>>>> a0ea7bf844a20294cb3ee0403c4363ceca8235c0
