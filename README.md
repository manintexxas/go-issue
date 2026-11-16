# Go Core Mastery

A structured learning repository documenting the study and practical exploration of Go, from language fundamentals to backend engineering concepts.

This repository serves as a progressive learning journal. Each topic is explored through conceptual understanding, practical implementation, experimentation, and code review.

The primary objective is not merely to complete exercises, but to understand how Go works, why certain design decisions matter, and how those concepts translate into maintainable software.

## Learning Objectives

- Understand Go language fundamentals and its underlying mechanisms.
- Develop a solid understanding of packages, modules, and project organization.
- Practice idiomatic Go through small, focused implementations.
- Explore software design principles and package boundaries.
- Develop testing, debugging, and documentation practices.
- Gradually apply Go concepts to backend engineering scenarios.

## Learning Approach

Each learning topic follows a progressive approach:

| Stage | Focus | Description |
| --- | --- | --- |
| Concept | Understanding | Study the underlying concepts, terminology, and mechanisms. |
| Exploration | Investigation | Examine behavior, constraints, edge cases, and design alternatives. |
| Implementation | Application | Translate concepts into practical Go code. |
| Testing | Verification | Verify expected behavior and identify incorrect assumptions. |
| Review | Evaluation | Examine correctness, design quality, maintainability, and trade-offs. |
| Documentation | Knowledge retention | Record important findings, decisions, and lessons learned. |

Not every topic requires the same depth of implementation. Exercises are designed according to the scope and practical relevance of each subject.

## Learning Roadmap

The roadmap is organized into progressive learning areas. Completion indicates that a topic has been studied and practiced, not necessarily mastered in every advanced scenario.

### Phase 1: Go Fundamentals

| Day | Topic | Learning Scope | Status |
| ---: | --- | --- | --- |
| 01 | Project Structure, Packages & Modules | Project organization, package boundaries, module management, `go.mod`, `go.sum`, imports, and toolchain fundamentals. | Completed |
| 02 | Type System & Zero Values | Built-in types, defined types, zero values, conversions, `nil`, and integer overflow. | Planned |
| 03 | Arrays & Slices | Array semantics, slice internals, capacity, allocation, and aliasing. | Planned |
| 04 | Maps | Map behavior, key constraints, lookup, mutation, and iteration. | Planned |
| 05 | Strings, Bytes & Runes | String representation, UTF-8, byte indexing, rune iteration, and conversion. | Planned |
| 06 | Pointers & Aliasing | Memory addresses, pointer semantics, mutation, and shared data. | Planned |
| 07 | Structs & Methods | Struct composition, fields, methods, and receiver semantics. | Planned |
| 08 | Interfaces | Interface satisfaction, implicit implementation, type assertions, and type switches. | Planned |
| 09 | Error Handling | Error values, wrapping, sentinel errors, custom errors, and error inspection. | Planned |
| 10 | Defer, Panic & Recover | Deferred execution, panic propagation, recovery, and initialization. | Planned |
| 11 | Generics | Type parameters, constraints, type sets, and generic functions. | Planned |
| 12 | Standard Library & Resource Lifecycle | Standard library exploration, resource ownership, cleanup, and lifecycle management. | Planned |

### Phase 2: Runtime & Memory

| Day | Topic | Learning Scope | Status |
| ---: | --- | --- | --- |
| 13 | Memory Fundamentals | Stack, heap, allocation, and escape analysis. | Planned |
| 14 | Garbage Collection | GC fundamentals, object lifetime, and memory retention. | Planned |
| 15 | Goroutines & Scheduler | Goroutines, OS threads, and Go scheduler fundamentals. | Planned |

### Phase 3: Concurrency

| Day | Topic | Learning Scope | Status |
| ---: | --- | --- | --- |
| 16 | Goroutine Lifecycle | Goroutine creation, execution, synchronization, and termination. | Planned |
| 17 | Channels & Select | Channel communication, buffering, closing, and select behavior. | Planned |
| 18 | Synchronization Primitives | Mutex, RWMutex, WaitGroup, and Once. | Planned |
| 19 | Atomic Operations & Memory Model | Atomic operations, happens-before relationships, and data races. | Planned |
| 20 | Context | Cancellation, deadlines, timeouts, and request propagation. | Planned |
| 21 | Concurrency Patterns | Worker pools, backpressure, deadlocks, and goroutine leaks. | Planned |

### Phase 4: Testing & Engineering Practices

| Day | Topic | Learning Scope | Status |
| ---: | --- | --- | --- |
| 22 | Testing Fundamentals | Unit testing, table-driven tests, subtests, and test organization. | Planned |
| 23 | Test Doubles & HTTP Testing | Interfaces for testing, mocks, stubs, and HTTP testing. | Planned |
| 24 | Debugging | Failure analysis, debugging techniques, and root cause investigation. | Planned |
| 25 | Benchmarking & Profiling | Benchmarks, CPU profiling, memory profiling, and pprof. | Planned |

### Phase 5: Backend Engineering

| Day | Topic | Learning Scope | Status |
| ---: | --- | --- | --- |
| 26 | HTTP Servers & Clients | HTTP fundamentals, request handling, JSON, and validation. | Planned |
| 27 | Middleware & REST | Middleware composition, routing, REST conventions, and API design. | Planned |
| 28 | Database & SQL | SQL integration, transactions, connection pooling, and database lifecycle. | Planned |
| 29 | Backend Security | Input validation, secure handling, and common application vulnerabilities. | Planned |
| 30 | Observability & Graceful Shutdown | Logging, metrics, tracing, and graceful application termination. | Planned |

## Implementation Log

Practical exercises are maintained as part of the learning process. They are intended to connect individual concepts with actual implementation decisions.

| Exercise | Subject | Concepts Applied | Status |
| --- | --- | --- | --- |
| Receipt Generator | Package & Module Fundamentals | Defined types, constants, package visibility, closures, anonymous functions, validation, error handling, rendering, and table-driven testing. | Completed |

Additional exercises will be recorded as the learning roadmap progresses.

## Design & Engineering Principles

The repository emphasizes several recurring engineering considerations:

- **Correctness:** Verify behavior instead of relying on assumptions.
- **Simplicity:** Prefer the smallest design that correctly solves the problem.
- **Separation of Concerns:** Keep responsibilities and package boundaries clear.
- **Explicit Contracts:** Make function behavior, input constraints, and error conditions understandable.
- **Testability:** Design implementations whose behavior can be verified.
- **Maintainability:** Consider readability, consistency, and future changes.
- **Trade-off Awareness:** Understand the benefits and limitations of implementation decisions.

These principles are learning guidelines rather than rigid rules. Their application depends on the problem and its context.

## Learning Notes

Each topic may contain additional notes covering:

| Category | Purpose |
| --- | --- |
| Core Concepts | Important definitions and underlying mechanisms. |
| Technical Findings | Discoveries, unexpected behavior, and language-specific details. |
| Implementation Decisions | Design choices and their reasoning. |
| Testing & Verification | Test cases, edge cases, and observed results. |
| Mistakes & Corrections | Incorrect assumptions and how they were resolved. |
| Further Exploration | Related concepts that deserve deeper investigation. |

## Progress Philosophy

Progress is measured through understanding and practical application, not simply by the number of completed topics.

A completed topic represents meaningful exposure and implementation. Further exploration remains possible as new requirements, projects, and engineering challenges arise.

The repository is intended to evolve alongside the learning process.
