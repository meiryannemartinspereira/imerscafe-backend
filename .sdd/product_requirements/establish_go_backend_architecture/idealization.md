---
status: Idealization In Review
date: 2026-08-24
---

# Establish Go Backend Architecture — Technical Idealization

## Summary

Establish the foundational architecture of the imerscafe backend using Go, preserving the deterministic game behavior of the legacy Spring Boot implementation while enforcing a strict architectural boundary between deterministic game rules and generative AI.

The Go backend is the authoritative source of truth for session state, game progression, recipe validation, ingredient comparison, scoring, and official round results.

Generative AI is an auxiliary integration used exclusively for natural-language and soft-skill evaluation. AI output must never determine recipe correctness, ingredient correctness, deterministic scoring, game state, or official round outcomes.

The first implementation phase focuses on the domain model and deterministic game rules before HTTP handlers or external AI integrations are introduced.

---

## Scope

### In Scope

- Establish the foundational Go project architecture.
- Define the deterministic domain model for the game.
- Represent ingredients and recipes as domain types.
- Preserve the existing recipe behavior from the legacy Spring Boot implementation.
- Implement deterministic recipe validation.
- Implement deterministic hard-skill scoring.
- Represent customer types and their associated deterministic scoring rules.
- Define session and round state concepts required by the current game behavior.
- Define clear boundaries between domain logic, application services, HTTP handlers, repositories, and AI integration.
- Define an AI integration interface that cannot become the authority for deterministic game decisions.
- Use Go Modules, `net/http`, JSON, and the standard Go testing package.
- Keep the initial implementation free from unnecessary external dependencies.
- Keep all source code identifiers, domain models, services, handlers, and comments in English.
- Preserve deterministic behavior from the legacy implementation where it is compatible with the new architectural rules.

### Out of Scope

- Implementing an external generative AI provider.
- Implementing authentication or authorization.
- Implementing a production database.
- Introducing a specific persistence technology.
- Implementing a frontend or modifying frontend behavior.
- Introducing multiplayer functionality unless explicitly required by a later requirement.
- Defining a winner or tournament mechanism.
- Defining tie-breaking rules that do not exist in the legacy behavior.
- Making AI responsible for official game scoring.
- Making AI responsible for recipe or ingredient validation.
- Introducing asynchronous AI processing as a required architectural behavior.
- Adding external frameworks or libraries that are not required by the current scope.
- Reproducing the legacy Spring Boot architecture directly.

---

## Technical Architecture

### Architectural Principles

The Go backend must follow these principles:

1. The backend is the authoritative source of truth for deterministic game state.
2. Deterministic rules must execute independently of generative AI.
3. Recipe validation must be deterministic.
4. Ingredient comparison must be deterministic.
5. Hard-skill scoring must be calculated exclusively by backend logic.
6. AI may evaluate only textual input related to soft skills and communication.
7. AI output must not modify official deterministic game state.
8. HTTP handlers must not contain game rules.
9. Domain logic must not depend on HTTP, persistence, or AI implementations.
10. The initial implementation must prefer the Go standard library.

### Affected System Areas

The initial backend architecture is organized into the following logical areas:

- `internal/domain/` — deterministic game concepts and rules.
- `internal/service/` — application services and orchestration.
- `internal/repository/` — persistence abstractions and initial in-memory implementations when required.
- `internal/handler/` — HTTP request and response handling.
- `internal/config/` — application configuration.
- `cmd/server/` — application bootstrap and dependency wiring.

The exact physical organization may evolve, but the logical separation of responsibilities must remain.

---

## Data Model Changes

The new Go backend must establish domain types corresponding to the deterministic behavior present in the legacy application.

### Ingredient

`Ingredient` represents a valid ingredient that can be selected when preparing a beverage.

The initial ingredient set is:

- `CAFE`
- `LEITE`
- `CHOCOLATE`
- `ESPUMA`
- `CHANTILLY`
- `CARAMELO`
- `CANELA`
- `ACUCAR`
- `BAUNILHA`
- `CREME`
- `AGUA`
- `GELO`

The Go implementation must use English identifiers while preserving the existing external ingredient values where required by the API contract.

For example:

- `IngredientCoffee = "CAFE"`
- `IngredientMilk = "LEITE"`
- `IngredientChocolate = "CHOCOLATE"`
- `IngredientFoam = "ESPUMA"`

No ingredient validation may depend on AI interpretation.

### Recipe

A `Recipe` represents the expected preparation for a beverage.

It must contain at least:

- beverage name;
- expected ingredients;
- base score.

Initial recipes derived from the legacy implementation are:

| Beverage | Expected Ingredients | Base Score |
|---|---|---:|
| Cappuccino | CAFE, LEITE, ESPUMA | 10 |
| Mocha | CAFE, CHOCOLATE, LEITE | 12 |
| Latte | CAFE, LEITE, BAUNILHA | 8 |
| Café com Chantilly | CAFE, CHANTILLY, CARAMELO | 15 |

These recipes are behavioral reference data from the legacy implementation and may later be moved to another source without changing the domain validation rules.

### Customer Type

The legacy implementation defines the following customer types:

| Customer Type | Bonus | Description |
|---|---:|---|
| CALMO | 5 | Customer who wants to feel welcomed. |
| APRESSADO | 10 | Customer who values speed. |
| ESTETICO | 15 | Customer who values presentation and experience. |
| EXECUTIVO | 20 | Customer who values efficiency. |
| ESPECIALISTA | 25 | Experienced customer who expects to be surprised. |
| INDECISO | 30 | Customer who needs clarity. |

Customer type and its deterministic bonus must remain backend-controlled.

### Session and Round State

The backend must maintain the official state required to:

- identify the current game session;
- identify the current scenario or recipe;
- maintain the current total score;
- advance the game to the next scenario;
- determine the official result of a preparation.

The state model must be deterministic and must not depend on AI output.

---

## Component Design

### Domain Layer

**Responsibility:** Represent deterministic business concepts and execute deterministic game rules.

Key domain concepts include:

- `Ingredient`
- `Recipe`
- `CustomerType`
- session state
- round state
- hard-skill validation
- deterministic scoring

The domain layer must not import AI clients, HTTP packages, database implementations, or framework-specific packages.

### Service Layer

**Responsibility:** Orchestrate application use cases while delegating deterministic decisions to the domain.

Key services include:

#### `RoundService`

Responsible for:

- creating or initializing a game session;
- obtaining the current round state;
- processing a player preparation action;
- coordinating recipe validation;
- applying deterministic scoring;
- advancing the game state;
- returning the official round result.

`RoundService` must never delegate deterministic decisions to AI.

#### `RecipeService`

Responsible for:

- loading recipes;
- exposing recipes required by application use cases;
- validating recipe structure when necessary.

Recipe correctness must remain deterministic.

#### `ScoreService`

Responsible for:

- calculating hard-skill points;
- calculating deterministic total score;
- applying only explicitly defined deterministic scoring rules.

It must not use AI output to determine hard-skill points.

#### `AICoordinationService`

Responsible only for the boundary between the backend and future generative AI functionality.

Its responsibility is limited to requesting soft-skill or natural-language evaluation.

The service may receive textual information such as:

- `responseStudent`
- `baristaName`

AI results are advisory data only.

The AI integration must never:

- validate ingredients;
- validate recipes;
- determine whether a preparation is correct;
- calculate hard-skill points;
- modify session state;
- determine the official round result.

The initial implementation may provide a placeholder implementation returning no soft-skill contribution.

### Repository Layer

**Responsibility:** Abstract persistence without containing business rules.

Initial repository abstractions may include:

- `SessionRepository`
- `RecipeRepository`

A `ScoreRepository` must only be introduced if persistence of scores is explicitly required by the current use case.

Repositories must not:

- validate recipes;
- calculate scores;
- decide round outcomes;
- call AI services.

An in-memory implementation is sufficient for the initial architecture where persistence is required.

### Handler Layer

**Responsibility:** Handle HTTP requests and responses.

Handlers are responsible for:

- receiving HTTP requests;
- validating request structure;
- decoding JSON;
- invoking application services;
- encoding JSON responses;
- mapping application errors to HTTP responses.

Handlers must not contain:

- recipe validation;
- ingredient comparison;
- scoring calculations;
- session state transitions;
- AI decision logic.

### Configuration Layer

**Responsibility:** Centralize application configuration.

The initial configuration may include:

- server port;
- environment;
- logging configuration;
- future AI service configuration.

Configuration must not contain business rules.

---

## AI Isolation Boundary

The architectural boundary between deterministic logic and generative AI is mandatory.

The official decision flow is:

```text
Client
  |
  v
HTTP Handler
  |
  v
RoundService
  |
  +----------------------+
  |                      |
  v                      v
Domain Rules         AI Coordination
  |                      |
  |                      v
  |                Soft-Skill Result
  |                      |
  +----------+-----------+
             |
             v
     Official Backend Result

The important distinction is that AI does not decide the deterministic result.

For example:

Selected Ingredients
        |
        v
Recipe Validation
        |
        v
Hard-Skill Score
        |
        +--------------------+
        |                    |
        v                    v
Official Game State     Optional AI Evaluation
                             |
                             v
                       Advisory Feedback

AI output must remain subordinate to the backend's authoritative state.

Data Flow
Primary Use Case: Player Prepares a Beverage
The client sends the selected ingredients and optional textual response to the backend.
The HTTP handler validates the request structure and decodes the JSON payload.
The handler delegates the operation to RoundService.
RoundService retrieves the current session state.
RoundService retrieves the current recipe.
The domain compares the selected ingredients with the expected recipe ingredients.
The domain determines whether the preparation is correct.
ScoreService calculates the deterministic hard-skill score.
The backend updates the official session score and round state.
If textual input is present and soft-skill evaluation is enabled, AICoordinationService may request an AI evaluation.
The AI result is treated only as advisory soft-skill information.
The backend constructs the official response.
The handler serializes the result as JSON and returns it to the client.

The AI evaluation must never change steps 6 through 9.

Deterministic Recipe Validation

Recipe validation must compare the selected ingredients with the expected ingredients.

The validation must detect at least:

correct ingredient set;
missing ingredients;
extra ingredients;
incorrect ingredients.

Ingredient ordering must not affect correctness if the legacy behavior considers the sets equivalent.

The same input must always produce the same validation result.

Implementation Plan
Phase 1: Domain Foundation
Define the complete Ingredient type.
Correct the existing IncredientCream naming error if present.
Define Recipe.
Define CustomerType.
Define the required session and round state types.
Define deterministic recipe validation.
Define deterministic hard-skill scoring.
Define the initial recipes from the legacy implementation.
Define deterministic customer bonus rules.
Add table-driven unit tests for domain rules.
Ensure domain code uses English identifiers and comments.
Keep the domain package independent from AI, HTTP, and persistence.
Phase 2: Application Services
Define RoundService.
Define RecipeService.
Define ScoreService.
Define the AICoordinationService interface.
Implement the initial AI coordination placeholder.
Implement session and round orchestration.
Ensure deterministic decisions are performed by domain logic.
Add service-level tests using in-memory dependencies where necessary.
Phase 3: Repository Abstractions
Define SessionRepository.
Define RecipeRepository.
Implement in-memory repositories where required by the current use cases.
Keep repository implementations free of business rules.
Add repository tests.
Phase 4: HTTP API
Define JSON request and response types.
Implement the health endpoint.
Implement endpoints required by the current round use cases.
Register handlers in cmd/server/main.go.
Map application errors to appropriate HTTP responses.
Ensure handlers contain no game rules.
Add HTTP tests using net/http/httptest.
Phase 5: Configuration, Integration and Validation
Implement environment-based configuration.
Wire dependencies in cmd/server/main.go.
Validate application startup configuration.
Run formatting and static analysis.
Run the complete test suite.
Verify deterministic behavior against the legacy implementation.
Document the architecture and AI isolation boundary.
Acceptance Criteria
Automated
 go mod verify succeeds.
 go test ./... passes.
 go test -race ./... passes.
 go build ./cmd/server succeeds.
 go fmt ./... produces no source changes.
 go vet ./... completes without warnings.
 Domain tests cover ingredient validation, recipe validation, and deterministic scoring.
 The domain package has no dependency on AI, HTTP, or persistence implementations.
 Recipe validation produces deterministic results for identical inputs.
 Hard-skill scoring is calculated exclusively by backend code.
 AI integration is hidden behind a dedicated interface.
 The AI coordination component cannot directly modify official game state.
 HTTP handlers delegate business decisions to application services.
 Repository implementations contain no game rules.
 The application starts successfully using go run ./cmd/server.
 The health endpoint returns HTTP 200.
Manual
 A technical reviewer confirms that deterministic game rules are independent of AI.
 A reviewer confirms that recipe correctness cannot be determined by AI.
 A reviewer confirms that ingredient correctness cannot be determined by AI.
 A reviewer confirms that official scoring is calculated exclusively by backend logic.
 A reviewer confirms that AI output is advisory and cannot overwrite official game state.
 A reviewer confirms that handlers contain no deterministic game rules.
 A reviewer confirms that repository implementations contain no business logic.
 A reviewer confirms that source code identifiers, domain models, services, handlers, and comments use English.
 The deterministic behavior is reviewed against the legacy Spring Boot implementation.
Technical Risks
Risk	Impact	Likelihood	Mitigation
Legacy behavior is incorrectly translated to Go	High	Medium	Use the existing Spring Boot implementation as a behavioral reference and create explicit domain tests for each deterministic rule.
AI logic leaks into deterministic game rules	High	Medium	Keep AI behind a dedicated interface and prohibit AI dependencies in the domain package.
AI output modifies official game state	High	Medium	Treat AI results as advisory data and allow state mutation only through deterministic application services.
Recipe validation becomes non-deterministic	High	Low	Implement exact ingredient comparison using deterministic domain functions and table-driven tests.
HTTP handlers accumulate business logic	Medium	Medium	Keep handlers limited to transport concerns and delegate all decisions to services/domain logic.
Repository layer contains business rules	Medium	Low	Define repositories strictly as persistence abstractions and test implementations independently.
Architecture introduces unnecessary dependencies	Medium	Medium	Use the Go standard library for the initial implementation.
Legacy concepts are incorrectly expanded into new requirements	Medium	Medium	Treat the legacy implementation as behavioral reference only and avoid introducing multiplayer, persistence, or ranking requirements without explicit product decisions.
English-only source code requirement is violated	Low	Medium	Include language compliance in code review and keep all identifiers and comments in English.
Recommendations
Domain-first implementation: Build and test deterministic game rules before implementing HTTP or AI integrations. This establishes the authoritative core of the system first.
Explicit AI boundary: Keep generative AI behind an application-level interface such as AICoordinationService. The domain must not know that AI exists.
Behavioral compatibility: Use the Spring Boot implementation as a behavioral reference for existing recipes, ingredients, customer types, scoring, and round progression, but do not reproduce its architectural structure.
Standard library first: Use net/http, encoding/json, testing, sync, and other standard-library packages before considering external dependencies.
Table-driven tests: Use Go table-driven tests for recipe validation and scoring because these rules have multiple deterministic input combinations.
Explicit domain errors: Use typed or sentinel errors for invalid recipes, invalid sessions, invalid ingredients, and invalid round states instead of relying exclusively on string comparisons.
State transitions in domain/application logic: Round and session state transitions must be explicit and validated outside HTTP handlers.
JSON contract separation: Keep HTTP DTOs separate from domain types so API representation does not dictate domain design.
No premature persistence: Use in-memory state where appropriate for the initial implementation. Do not select a production database until persistence requirements are explicitly defined.
No premature AI provider: Keep the AI interface independent of a specific provider until the external AI contract is defined.
Dependencies and Prerequisites
Required
Go 1.22 or later.
Go Modules.
Go standard library.
Existing go.mod.
Existing Go project structure.
Existing Reference

The legacy Spring Boot implementation is the behavioral reference for:

ingredients;
recipes;
customer types;
session behavior;
deterministic scoring;
round progression;
AI-related textual evaluation boundaries.

The legacy implementation must not dictate the architecture of the new Go backend.

Architectural Reference

The project's architectural guidelines define the authoritative boundaries between:

frontend;
deterministic backend;
generative AI integration.

These guidelines take precedence over legacy implementation details.

Files and Areas to Preserve
go.mod must not be modified unless a dependency is explicitly required.
.gitignore must preserve its existing rules.
Existing domain behavior must not be removed without an explicit requirement.
The initial implementation must not introduce an external AI SDK without an approved integration requirement.
Open Questions

The following decisions are intentionally left outside the current architectural foundation because they are not defined by the current requirement or legacy implementation.

 AI Evaluation Timing — Should soft-skill evaluation occur synchronously during the request or asynchronously after the official round result? Decision owner: Product Owner / System Architect.
 Soft-Skill Score Authority — Should AI-generated soft-skill evaluation contribute to the official total score, remain advisory, or be presented separately? Decision owner: Product Owner / Game Designer.
 Persistence Strategy — Should sessions and scores eventually be persisted in a database, and if so, which technology should be used? Decision owner: System Architect / DevOps.
 Session Lifecycle — What are the exact lifecycle rules for creating, expiring, restarting, and completing a game session? Decision owner: Product Owner.
 Recipe Selection — Should recipes remain sequential as in the current implementation, become random, or be selected by another mechanism? Decision owner: Product Owner / Game Designer.
 Customer Type Selection — Should customer types continue to be randomly selected, or should selection follow a deterministic scenario sequence? Decision owner: Product Owner / Game Designer.
 API Contract — What exact request and response JSON contracts must be maintained for compatibility with the frontend? Decision owner: API Designer / Frontend Lead.
 Authentication — Will the backend require authentication or authorization in a future phase? Decision owner: Product Owner / System Architect.
 AI Data Privacy — What textual data may be sent to an external AI provider, and what retention restrictions apply? Decision owner: Product Owner / Security Lead.
Architectural Invariants

The following invariants are mandatory for all future implementations:

The backend is the authoritative source of truth for deterministic game state.
Recipe validation is deterministic.
Ingredient validation is deterministic.
Hard-skill scoring is deterministic.
Official round results are determined by backend logic.
AI cannot determine recipe correctness.
AI cannot determine ingredient correctness.
AI cannot determine hard-skill scoring.
AI cannot directly modify official game state.
The domain layer must remain independent of AI providers.
HTTP handlers must not contain game rules.
Repository implementations must not contain business rules.
All source code identifiers, domain models, services, handlers, and comments must use English.
The legacy Spring Boot implementation is a behavioral reference, not an architectural constraint.
New architectural decisions must not contradict the project's architectural guidelines.