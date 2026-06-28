# Definition of Done

This document defines the team's shared minimum completion standard for all Product Backlog Items (PBIs).

## General Requirements

- [ ] All acceptance criteria are met
- [ ] Code is reviewed and approved by at least one team member
- [ ] No regressions in existing functionality
- [ ] All automated tests pass
- [ ] Code follows project conventions and style guidelines

## Code Quality

- [ ] Code is clean, readable, and well-structured
- [ ] No hardcoded values or magic numbers
- [ ] Error handling is implemented where appropriate
- [ ] Security best practices are followed
- [ ] No sensitive data or credentials are committed

## Testing

- [ ] Unit tests are written and passing for critical modules
- [ ] Integration tests are written and passing
- [ ] Each critical module achieves at least 30% line coverage
- [ ] Automated quality requirement tests (QRTs) pass
- [ ] Manual testing is performed for user-facing features
- [ ] Edge cases are considered and tested

## Quality Requirements

- [ ] Relevant quality requirements are satisfied or documented as not applicable
- [ ] Quality requirement tests (QRTs) pass in CI
- [ ] `docs/quality-requirements.md` is current
- [ ] `docs/quality-requirement-tests.md` is current

## CI and Automation

- [ ] CI lint step passes (golangci-lint)
- [ ] CI unit test step passes
- [ ] CI integration test step passes
- [ ] CI coverage check passes (30% for critical modules)
- [ ] CI QRT step passes
- [ ] CI additional QA check passes (govulncheck)
- [ ] CI link check passes (Lychee)

## Documentation

- [ ] Code is documented with comments where necessary
- [ ] API documentation is updated (if applicable)
- [ ] `docs/testing.md` reflects current test status
- [ ] `docs/user-acceptance-tests.md` is current
- [ ] CHANGELOG.md is updated with user-visible changes

## Deployment

- [ ] Feature is deployable to staging/production
- [ ] Environment variables are configured
- [ ] Feature works in the deployed environment

## User Story Specific

- [ ] User can complete the described workflow
- [ ] Feedback is clear and actionable
- [ ] UI/UX follows design guidelines
- [ ] Accessibility requirements are met

## PBI Types

### User Story
- [ ] All acceptance criteria from the issue are satisfied
- [ ] User workflow is complete and functional
- [ ] Feedback loop is implemented

### Bug Fix
- [ ] Bug no longer occurs when following reproduction steps
- [ ] No regression in related functionality
- [ ] Root cause is identified and addressed

### Technical / Infrastructure
- [ ] Technical requirements are met
- [ ] Performance requirements are met
- [ ] Scalability considerations are addressed

### Documentation
- [ ] Content is accurate and up-to-date
- [ ] Links are valid and accessible
- [ ] Format follows project standards

## Sprint Completion

- [ ] PBI is marked as Done in the issue tracker
- [ ] PBI is moved to the Done column in the Sprint Backlog
- [ ] CHANGELOG.md is updated for user-visible changes
- [ ] Release notes are prepared (if applicable)

## Release Requirements

- [ ] All MVP PBIs are completed
- [ ] SemVer release is created with proper tag
- [ ] Release notes are comprehensive
- [ ] Deployment is verified in production
- [ ] Video demonstration is recorded and uploaded
