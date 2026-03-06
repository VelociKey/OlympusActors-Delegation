# Conductor Implementation — AntigravitySpace
**Implements:** Conductor Protocol (Poly-Workspace Manifesto v3.1, Section 4)
**Primary Tools:** Antigravity, Gemini-CLI

## 1. Directory Structure
The `/conductor/` directory at the federation root serves as the authoritative Shared Intent Layer. It is the "Supreme Law of the Federation" for the AntigravitySpace instantiation, housing the Tracks Registry, semantic taxonomy, and shared symbols that coordinate all agentic and human activity across constellations.

## 2. Cycle Track Schema (jeBNF)
All high-level work units follow the `CYC-NNN` numbering convention. A cycle track record, expressed in jeBNF, contains:
- **Track-ID**: The unique `CYC-NNN` identifier.
- **Current State**: The lifecycle status (Active, Complete, On Hold).
- **Active Goal**: The semantic "North Star" for the current cycle.
- **Landing Status**: Verification of committed results and chronicle production.

*See `jeBNF_PRIMER.md` for format reference.*

## 3. Coord-State
The `.coord-state.jebnf` file maintains the canonical workspace state hash. It tracks the integrity of each node in the federation:
- **DIRTY Entries**: Indicates a workspace has uncommitted changes or a hash mismatch (e.g., `scaffold-workspace` and `openclaw-research` in the 2026-03-03 review).
- **GIT_HASH / Coord-Hash**: The relationship between the physical git commit and the semantic coordination hash ensuring structural integrity.

## 4. Tool Integration

### 4.1 Antigravity
Antigravity utilizes the Conductor state to provision and orchestrate fleet operations. It consumes binary execution reports (e.g., `binary-id-run-example.ansi`) to verify task completion. Primary fleet tools such as `fleet-index`, `fleet-substrate`, and `fleet-mission-runner` are the principal consumers of this state.

### 4.2 Gemini-CLI
Gemini-CLI interacts with the Conductor via:
- **Shadow-Sync Handshake**: A "Request for Summary" flow where Gemini-CLI retrieves semantic context from the `/conductor/` before initiating code modifications.
- **Semantic Delta-Sync**: A synchronization process performed upon network re-entry to align local reasoning with the federation's latest state.

## 5. Landing Protocol
A successful "Landing" in the AntigravitySpace federation requires:
- **Committed Results**: All code and configuration changes must be pushed to their respective repositories.
- **Updated Global Registry**: The `conductor/tracks.md` and related `.jebnf` coordination files must be updated.
- **Chronicle Production**: A "Kidder-style" prose narrative (Story Mode) must be produced to summarize the technical transformation.

*See Multi-Flight Governance for the distributed variant of this protocol.*
