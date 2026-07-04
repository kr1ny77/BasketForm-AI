# ADR-003: JWT Authentication

## Status
Accepted

## Context
BasketForm-AI requires user authentication to protect personal data. Users access via web browser.

Options considered: Server-side sessions, JWT tokens, OAuth2.

## Decision
Use JWT for authentication. On login, server generates JWT with user ID, signed with HMAC-SHA256. Token stored in HttpOnly cookie, validated on every protected request via middleware. Expiry: 24 hours. Passwords: bcrypt hashing.

## Consequences and Tradeoffs

**Positive:** Stateless (no server-side session storage), works across page refreshes, simple implementation, easy to extend.

**Negative:** Cannot revoke tokens before expiry, token theft via XSS (mitigated by HttpOnly), secret key management required, no refresh token mechanism.

## Quality Requirements Addressed
- QR-002 (Security): Passwords hashed with bcrypt, JWT signed, HttpOnly cookies
- QR-003 (Usability): Seamless session across page refreshes
