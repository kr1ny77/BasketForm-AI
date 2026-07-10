# Week 5 Reflection — BasketForm-AI

## Learning points

- **Documenting architecture** (static, dynamic, and deployment views as PlantUML diagrams-as-code) gave us a shared, versioned picture of the system instead of tribal knowledge. Writing the deployment view in particular is what pushed us to explicitly document the VM deployment and CI/CD steps, which later fed directly into resolving the customer's "deployment complexity" feedback (PBI-028).
- **Recording ADRs** made the biggest technical decision of the sprint — replacing the locally hosted feedback model with an API-based LLM — traceable. The old approach was overloading the server and slowing startup; writing the decision down as an ADR (not just fixing it silently) means we can point to *why* the swap happened and confirm the core scoring algorithm itself didn't change, which matters for anyone questioning result consistency later.
- **Refining the workflow** paid off concretely this sprint: the retrospective's action item to require a UAT dry-run before calling something "done" is why the Friends system UAT with the customer went smoothly — both "send request" and "accept request" flows passed on the first attempt in the live review, with no last-minute blockers.
- **Managing configuration and deployment documentation** turned a recurring customer complaint (Python dependency / deployment complexity) into a closed item (PBI-028, marked Done). Writing it down once means new team members and the customer no longer have to ask us directly how to run the product.
- **Delivering MVP v2** meant shipping a wide slice at once — LLM feedback, friends, report sharing, progress tracking, pro comparison, RU translation fixes, architecture docs, ADRs, and storage hardening (deadlock fix, batch-delete cleanup) — while explicitly deferring batch upload (PBI-023) rather than half-implementing it. Saying no to a feature on purpose, with a documented reason, kept the release coherent.
- **Reviewing the increment with the customer** surfaced a useful negative result: the customer's idea of a direct video-to-video comparison feature ("today's throw vs. yesterday's throw") was dropped on the spot once the full-stack developer pointed out the existing points-based progress system (throw, foot position, arm position criteria) already covers that need. Not every customer idea should become a backlog item — sometimes the right response is explaining what already solves it.

## Validated assumptions

- Keeping architecture docs and ADRs versioned in the repo (rather than a separate wiki) made them directly reusable as evidence when responding to customer feedback about deployment complexity — the docs weren't just process theater.
- The "no direct API model change to the scoring algorithm" assumption held: swapping the ML backend for an API-based model improved speed and stability without customers noticing any change in how scores are computed.
- Our assumption that stricter Definition-of-Done (UAT dry-run pre-review) would prevent live-demo failures was confirmed — both UAT scenarios passed cleanly in front of the customer.
- The customer confirmed that all originally scoped requirements have been delivered or explicitly approved in modified form, validating that our backlog re-negotiations throughout the project were the right call rather than scope creep.

## Friction and gaps

- The team is entering the second-to-last sprint before final defense with real deadline pressure. The customer explicitly flagged the risk of ending up with a "broken product" if the team overcommits on the newly approved multi-throw video segmentation feature.
- The segmentation feature's technical approach (using ball-release as a trigger to queue and process preceding frames) is only a proposed design from the customer's brainstorming — it hasn't been validated against our actual pipeline yet, so there's a real chance it turns out harder to implement than discussed.
- We don't yet have a fallback plan formalized beyond "don't do it if it's too hard" — that's a good instinct but not yet a concrete decision point (e.g., how many days before we cut losses).
- Batch upload (PBI-023) remains deferred for a second sprint in a row; we haven't yet confirmed whether it's realistically in scope for Sprint 4 or should be cut from the project entirely given the remaining time.

## Planned response

- Timebox the multi-throw video segmentation feature to 3 days; if it isn't stable and integrated by then, pivot to a simpler manual-trimming fallback to protect release stability before final defense, per the retrospective's action point.
- Add a mandatory mid-sprint regression test for the core single-throw analysis pipeline specifically, so the new multi-video ingestion logic can't silently break the existing working product.
- Continue writing ADRs at the point of decision (e.g., before/while implementing the segmentation trigger design) rather than after, so the recorded rationale reflects real-time reasoning rather than a reconstruction.
- Explicitly decide, early in Sprint 4 planning, whether batch upload (PBI-023) is still in scope given the two sprints remaining, rather than letting it roll over by default a third time.
- Keep using the customer review as a filter against scope creep — when a requested feature is already served by an existing mechanism (as with video comparison vs. the progress system), document that decision and move on instead of adding redundant work.
