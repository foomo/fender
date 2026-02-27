# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What is Fender

Fender is a Go validation library that provides a unified, composable way to validate data. It uses generics to build type-safe validation pipelines with three layers: **fender** (orchestrator), **fend** (field-level validators), and **rule** (individual validation rules).

## Commands

All commands go through `make`. Run `make help` for the full list.

- `make check` — full pipeline: tidy, generate, lint, test
- `make lint` / `make lint.fix` — run golangci-lint (config in `.golangci.yml`)
- `make test` — run tests with `-tags=safe`
- `make test.race` — run tests with race detector
- `make test.update` — run tests with `-update` flag
- Single test: `go test -tags=safe -v -run TestName ./path/to/package/...`
- Tool setup: `make .mise` (installs mise dependencies), `make .lefthook` (configures git hooks)

Build tag `-tags=safe` is required for all builds and tests.

## Architecture

Three-layer validation pipeline, each layer wrapping the one below:

**`rule/`** — Atomic validation rules. A `rule.Rule[T]` is `func(ctx context.Context, v T) error`. Returns `*rule.Error` on validation failure or `rule.ErrBreak` to short-circuit. Includes built-in rules: `RequiredString`, `MinString`, `Email`, `Match`, `Range`, `UUID`, etc. Type aliases (`StringRule`, `IntRule`, etc.) provide convenience.

**`fend/`** — Field-level validators. A `fend.Fend` is `func(ctx context.Context, mode Mode) error`. Created via `fend.Field(name, value, rules...)`, `fend.Var(value, rules...)`, or `fend.DynamicField`/`fend.DynamicVar` for runtime-typed rules. `fend.Union` implements OR-logic across rule sets. `fend.Rules[T]` allows prepending shared rules to fields.

**`fender.go` (root)** — Orchestrator with three validation modes:
- `fender.All` — runs all fends, collects all errors from all rules
- `fender.First` — stops at first fend with an error, returns only its first rule error
- `fender.AllFirst` — runs all fends but returns only the first fend's errors

**`config/`** — Shared configuration: a `go-playground/validator` instance and delimiter/regex constants.

**Error hierarchy**: `fender.Error` contains `[]*fend.Error`, each containing `[]*rule.Error`. Use `fender.AsError()`, `fend.AsError()`, `rule.AsError()` to unwrap. String format uses `;` between rule errors and `:` between fend errors.
