# Contributing to BasketForm-AI

Thank you for your interest in contributing to BasketForm-AI. This document describes the development workflow, code review process, and quality standards expected for contributions.

## Development Workflow

### Branching Strategy

The team uses trunk-based development with short-lived feature branches:

1. All changes go through Pull Requests (PRs) to the protected `main` branch
2. Direct pushes to `main` are disabled
3. Feature branches are named `<type>/<short-description>` (e.g., `feature/multi-upload`, `fix/avatar-upload`)

### Issue-Linked Development

Every non-automated change starts from a GitHub issue:

1. Find or create an issue describing the work
2. Create a branch from the issue
3. Implement changes in the branch
4. Open a PR linking to the issue
5. Get at least one review approval
6. Merge to `main` after CI passes

### Pull Request Process

When opening a PR:

1. Provide a clear summary of changes
2. Link the related issue
3. Describe testing performed
4. Verify acceptance criteria are met
5. Update CHANGELOG.md if user-visible changes are included
6. Request review from a different team member

### Required Reviews

- At least one approval from a different team member is required
- Authors cannot approve their own PRs
- All CI checks must pass before merge

## Code Quality Standards

### Linting

Run the linter before committing:

```bash
golangci-lint run
```

### Testing

Run tests before committing:

```bash
# Unit and integration tests
go test -race -coverprofile=coverage.out ./...

# Quality requirement tests
go test -tags=qrt -v ./internal/qrt/...
```

### Coverage

Critical modules (handlers, services) must maintain at least 30% line coverage.

### Security

- Never commit secrets, API keys, or credentials
- Use environment variables for configuration
- `.env` files are gitignored; use `.env.example` as a template

## Commit Messages

Write clear, concise commit messages that describe what changed and why.

## Documentation

When making changes that affect:

- **API endpoints**: Update relevant handler documentation
- **User-facing features**: Update user stories in `docs/user-stories.md`
- **Architecture**: Update architecture docs in `docs/architecture/`
- **Quality requirements**: Update `docs/quality-requirements.md` and `docs/quality-requirement-tests.md`
- **Development process**: Update `docs/development-process.md`

## Getting Help

- Check existing issues and documentation first
- Ask questions in the issue or PR comments
- Review the [Development Process](docs/development-process.md) for workflow details
