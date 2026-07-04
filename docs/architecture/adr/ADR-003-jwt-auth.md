# ADR-003: JWT Authentication

## Status
Accepted

## Context
BasketForm-AI requires user authentication to protect personal data (videos, results, friends). Users access the application via web browser.

Options considered:
- Server-side sessions (cookies + session store)
- JWT tokens (stateless)
- OAuth2 / third-party auth

## Decision
Use JWT (JSON Web Tokens) for authentication. On login, the server generates a JWT containing the user ID, signed with HMAC-SHA256. The token is stored in an HttpOnly cookie and validated on every protected request via middleware.

Token properties:
- Expiry: 24 hours
- Signing: HS256 with server-side secret
- Claims: `user_id`, `exp`

Password storage: bcrypt hashing with default cost factor.

## Consequences and Tradeoffs

**Positive:**
- Stateless — no server-side session storage needed
- Works across page refreshes (cookie persists)
- Simple implementation with `golang-jwt/jwt` library
- Easy to extend with additional claims

**Negative:**
- Cannot revoke tokens before expiry (no server-side session store)
- Token theft via XSS gives full access (mitigated by HttpOnly cookie)
- Secret key management required (stored in env var)
- No refresh token mechanism (user must re-login after 24h)

## Quality Requirements Addressed
- **QR-002 (Security)**: Passwords hashed with bcrypt; JWT signed with secret; HttpOnly cookies prevent XSS token theft
- **QR-003 (Usability)**: Seamless session across page refreshes; no re-login required during session
