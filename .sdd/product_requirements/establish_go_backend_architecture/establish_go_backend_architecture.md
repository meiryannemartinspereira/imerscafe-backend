---
status: New Requirement
date: 2026-08-24
title: Establish Go Backend Architecture
requirementType: new_feature
---

## Description

Establish the foundational architecture for the imerscafe backend using Go.

The backend must be the authoritative source of truth for all deterministic game rules, including session state, game progression, recipe validation, ingredient comparison, scoring, and official round results.

The architecture must strictly separate deterministic game logic from generative AI functionality.

Generative AI may only be used as an auxiliary component for soft-skill and natural-language evaluation. It must never determine game state, recipe correctness, ingredient correctness, deterministic scoring, or official round outcomes.

The initial implementation must use Go, Go Modules, net/http, JSON, and the standard Go testing package.

All source code identifiers, domain models, services, handlers, and comments must use English.

## Acceptance Criteria

1. The project uses Go Modules.
2. The backend follows a domain-oriented architecture.
3. Deterministic game rules are isolated from AI integration.
4. The backend remains the authoritative source of truth for game state.
5. Recipe and ingredient validation are deterministic.
6. Scoring is calculated exclusively by the backend.
7. AI integration is isolated behind a dedicated interface.
8. AI cannot determine game state, recipe correctness, ingredient correctness, or official scoring.
9. Source code identifiers, domain models, services, handlers, and comments use English.
10. The project can be tested using Go's standard testing package.
11. The initial architecture avoids unnecessary external dependencies.

## Notes

This project is a Go reimplementation of an existing Spring Boot backend.

The existing system contains deterministic game logic for recipes, ingredients, customer types, sessions, scoring, and round progression, as well as an AI integration for soft-skill evaluation.

The existing implementation should be treated as a behavioral reference, not as an architectural constraint.

The architectural guidelines defined for the new Go backend take precedence over the legacy implementation.

The migration must preserve deterministic game behavior while correcting the architectural boundary between deterministic game rules and generative AI.

The first implementation phase should establish the domain model before implementing HTTP handlers or external AI integrations.