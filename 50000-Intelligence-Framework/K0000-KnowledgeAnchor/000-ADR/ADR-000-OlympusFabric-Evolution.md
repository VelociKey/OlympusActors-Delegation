# ADR-000: WraithFabric Evolution Strategy

**Status:** Proposed
**Date:** 2026-02-13
**Subject:** Upgrading DSLEngine to WraithFabric

## Context
The existing `DSLEngine` repository provides a mature foundation for grammar parsing and transpilation. However, to support the "Strict Mode" requirements of the Plural Fabric, we need a more integrated, AI-native engine that adheres to the `.wraith` workspace standards and uses modern Go patterns.

## Decision
We will establish **WraithFabric** as the successor to `DSLEngine`.

### Key Improvements:
1.  **Identity First**: Integrated SPIFFE/Identity attestation for generated artifacts.
2.  **jeBNF Optimization**: Hard-coded support for jeBNF (PEGs) with automatic eBNF transformer.
3.  **Modern Go**: Standardized IR (Intermediate Representation) using `Go 1.25+`.
4.  **AI Orchestration**: Built-in prompts for Gemma 3 (via Ollama) to assist in DSL authoring.
5.  **Multi-Target Artifacts**: First-class support for gRPC, SQL, and Cloud infrastructure.

## Implementation Phases

### Phase 1: Foundation (Current)
- Initialize `.wraith` structure.
- Port IR and Lexer logic from `OlympusGrammar`.
- Merge robust recursive descent parsing from legacy `DSLEngine`.

### Phase 2: Transpilation (Week 1)
- Implement `Transformer` (eBNF ↔ jeBNF).
- implement `Generator` registry.
- Target: YAML/JSON parity with legacy system.

### Phase 3: Fabric Integration (Week 2)
- Wire into `CapabilityProvider` via MCP.
- Implement `SovereignAudit` hooks to verify safety of generated code.

### Phase 4: AI Mastery (Week 3)
- Integrate Gemma 3 for requirements-to-grammar workflows.
- Implement the "Design Partner" agent within `WraithFabric`.

## Consequences
- Requires a clean-room migration of core logic from `xxDSLEngine`.
- Deprecates the numerical prefixing of legacy `DSLEngine` in favor of full Wraith semantics.
