# YumRun

YumRun starts as a modular monolith a single application split into independent modules, each with clear boundaries. It has HTTP handlers, SQL queries, repositories, application layer, component tests, and inter-module communication. Each decision will be documented and evolve as the project grows.

- **Modular monolith**: two modules to start with, `orders` and `delivery`, wired into one binary by `cmd/main.go`. Cross-module calls are Go function calls, not network calls. As a monolith first even if we get the boundaries wrong, moving them is much easier than untangling microservices. And if we ever need to deploy a module separately, that's a trivial step from here.

- **Per-module Postgres schema**: each module owns its own schema in a single Postgres database. Migrations live in the module that owns them.

- **The common/ package**: general-purpose helpers shared across modules, like UUID generation, error types, HTTP plumbing, and logging. No application logic lives there. Anything specific to `orders` or `delivery` stays inside its own module.

- **Code generation for HTTP and SQL**: HTTP handlers are generated from an OpenAPI spec, and database access code is generated from SQL queries. It's less boilerplate to maintain by hand, and the codebase stays consistent across modules.

- **Clean Architecture**: each module splits into three layers. `api/` holds the entry points, `adapters/` holds external dependencies, and `app/` holds business logic and orchestration. Dependencies flow through interfaces defined close to where they're used.

- **Repositories**: they sit in the `adapters/` layer and abstract database access, so the `app/` layer never talks to SQL directly. YumGo uses a closure-based update pattern for safely loading and modifying an entity in one transaction.

Component tests: most of the test pyramid is component tests. We call one HTTP endpoint, then verify the result via another endpoint. They use Postgres and no external services. Unit tests cover business logic in app/, and integration tests cover adapters. See more: component tests.

Read models: when the data model we read doesn't match any existing model, we write a specialized SQL query and map the result straight to the HTTP response. We don't use the app layer because reads are just data retrieval, not business logic. See more: the first read model for the basics, and ordering and filtering on a read model for more complex use cases.

Inter-module communication: cross-module calls are Go function calls, with no network or serialization. Each module exports a small "contract" interface that defines what other modules can call. The rest is used only within the same module. See more: calling another module.
