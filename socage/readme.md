# Socratic Coding Mentor Agent for Zone01

This workspace contains reusable instructions and memory files for a Socratic coding mentor agent.

The goal of the agent is to help with Zone01 coding projects by guiding reasoning first, giving hints before code, and adapting help based on the learner's level and failed attempts.

---

## File roles

### `soc-agent-sharp.md`

Main cheap/sharp behavior instruction file.

Use this as the active VS Code agent instruction.

This file should stay short and practical so it does not waste tokens.

---

### `learner_state.json`

Current project-specific learner state.

This file changes from project to project.

Example:

```json
{
  "current_project": "groupie-tracker",
  "current_topic": "filters",
  "current_problem": "range filter for creation date",
  "failed_attempts": 0,
  "last_help_mode": null,
  "project_topic_levels": {
    "html_forms": 4,
    "javascript_dom": 3,
    "go_handlers": 4,
    "filter_logic": 3
  }
}
```

### `user_global_level.json`
Portable long-term learner profile.

Copy this file manually into the next project so the agent can continue from the previous overall learner level.

Example:

```json
{
    "user_id": "Merllin",
    "global_level": 4.8,
    "topic_levels": {
    "go_basics": 5,
    "go_pointers": 4,
    "go_slices": 5,
    "go_maps": 4,
    "go_structs": 5,
    "go_errors": 4,
    "go_files": 4,
    "go_http": 4,
    "go_json": 4,
    "go_concurrency": 3,
    "git": 5,
    "docker": 4,
    "algorithms": 4,
    "html": 4,
    "css": 3,
    "javascript": 3
    },
        "known_tools": [],
        "weak_tools": [],
        "repeated_mistakes": [],
        "preferences": {
            "style": "balanced_socratic",
            "code_policy": "no_full_code_until_requested_or_blocked",
            "preferred_language": "english_with_possible_greek"
        }
    }
   ```

### `convDB.md`
Current project conversation summaries and history.
This file can stay project-local.

Use it for:

* session summaries
* repeated mistakes
* solved concepts
* next steps
* important design decisions

Do not paste every full conversation unless needed. Prefer short summaries.

### `docs/soc-agent-detailed.md`
Full theory/design reference.

This file is not loaded every time. Use it only when improving or redesigning the agent behavior.

---
### How to use this in every project


At the start of a new project:

1. Copy the Socratic agent files into the project.
2. Copy the latest user_global_level.json.
3. Reset or update learner_state.json for the new project.
4. Use soc-agent-sharp.md as the VS Code agent instruction.
5. Ask the agent to update convDB.md, learner_state.json, and user_global_level.json after important sessions.
---
### Optional note for soc-agent-sharp.md

Add this line inside soc-agent-sharp.md:

    For deeper design reference, see docs/soc-agent-detailed.md, but do not load it unless the user asks to refine the agent behavior.
  
This lets the active agent know the detailed file exists without wasting tokens during normal project help.

---
### Prompt in VS Code agent
    Act as my Socratic mentor using soc-agent-sharp.md.
    You may update learner_state.json, convDB.md, and user_global_level.json automatically when progress changes.
    At the end of each session, update all memory files.