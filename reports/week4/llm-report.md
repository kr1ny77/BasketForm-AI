# LLM Usage Report — Week 4

## AI/LLM Tools Used

LLM tools (Claude / GitHub Copilot) were used during Sprint 2 for the following tasks:

| Task | Tool | Usage |
|------|------|-------|
| Go backend API design | LLM | Generated initial route handlers and API endpoint structure |
| Python ML script integration | LLM | Assisted with `exec.Command` pattern for Go→Python subprocess call |
| Test scaffolding | LLM | Generated initial unit test structure for new endpoints |
| Documentation drafting | LLM | Assisted with markdown report formatting and structure |
| CI workflow review | LLM | Reviewed `ci.yml` configuration for correctness |

## Impact

- **Time saved:** Estimated 3–5 hours on boilerplate code and test scaffolding
- **Code quality:** LLM suggestions were reviewed and modified before merge; no blind copy-paste
- **Limitations:** LLM could not assist with MediaPipe-specific debugging or video processing edge cases; those required manual testing and domain knowledge

## Policy

All LLM-generated code was reviewed by a team member before merging. No sensitive data or credentials were shared with LLM tools.
