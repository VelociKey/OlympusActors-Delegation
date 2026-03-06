# Multi-Flight Governance Protocol
**Protocol v3.2: Multi-Dev / Multi-Machine Orchestration**

This document is normative: it describes the intended governance of the Multi-Flight Federation layer, not its current implementation state. It sits at the Federation governance layer, above the Poly-Workspace Manifesto in the fleet's documentation hierarchy. The implementation status of these patterns is tracked in Fleet Review documents and Conductor cycle records.

## 1. The Distributed Challenge: Cognitive Divergence
In a poly-workspace environment, the primary risk is not file conflicts, but **Cognitive Divergence**. If multiple human-AI teams ("Flights") work on different machines without a shared semantic state, the system will eventually fragment.

## 2. The Distributed Flight Protocol

### A. Composite Platform Identity (CPI)
A Composite Platform Identity (CPI) is a cryptographically signed attestation token that binds a Flight to its execution context. The token must satisfy the host Sovereign Policy (jeBNF) at the time of execution. The mechanism by which the CPI is produced is determined by the deployment profile.

| Deployment Context | CPI Mechanism | Notes |
|:---|:---|:---|
| Bare metal / owned hardware | TPM attestation + hardware fingerprint | Highest assurance; preferred for sovereign nodes |
| Cloud VM / ephemeral compute | Workload identity (e.g., SPIFFE, cloud IAM) | Token scoped to VM lifecycle |
| Container / CI environment | Short-lived signed attestation token | Issued at job start; expires at job end |
| Developer workstation | Hardware key (YubiKey, Passkey) + OS identity | Human-in-Command binding |

**The "Refusal to Run" Mandate:** A Sovereign Seal artifact is prohibited from executing on any platform that cannot produce a valid CPI token satisfying the current Sovereign Policy (jeBNF), is present on the Global Prohibition Registry, or violates the active Temporal Window.

### B. The Immutable Flight Recorder (Black Box)
To ensure the absolute integrity of the system's evolution, every workspace MUST maintain an **Immutable Flight Recorder**:
- **Non-Optional Logging**: All SAAC Adjudication decisions, composite identity fingerprints (MFI), and **Attested Chain-of-Custody (ACoC)** logs are streamed to the `C0700/FLUX_LEDGER.jebnf`.
- **The "No-Off-Switch" Mandate**: The Flight Recorder is a foundational requirement. It serves as the "Black Box" of the Federation, providing an audited trail of every machine, user, and agent that has interacted with a Sovereign Work-Unit.

### C. The Sovereign Mesh (Live Coordination)
Legacy Git processes (Commit-Push-Pull) are too reactive for high-velocity AI coordination. We transition to a **"Resonate-and-Align"** model:
- **The Resonance Bus**: When on-network, flights broadcast a **"Semantic Pulse"** (jeBNF) via a mesh-coordination bus. This pulse shares Turn-Level intent and Goal updates in real-time, providing immediate "Air Traffic" visibility to all flights.
- **The Pulse Packet**: Contains `[Flight-ID, Track-ID, Intent-Hash, Current-Goal]`.
- **Off-Network Satchel**: When off-network, all reasoning traces and turn-level telemetry are buffered locally. Upon re-entry, a **"Semantic Delta-Sync"** is performed before any code is pulled, ensuring the flight's intent is aligned with the federation's latest state.

### D. Semantic State Synchronization (LanceDB)
To avoid redundant token usage and cognitive re-indexing:
- **Distributed Substrate**: High-dimensional embeddings (LanceDB) are treated as a shared peer-to-peer asset. Flights may peer-sync their `C0500` vector indexes directly over the mesh to avoid redundant indexing of the same thematic blocks.
- **Hash Verification**: When a new flight enters a workspace, it first verifies its local **LanceDB** index against the mesh-wide hashes. If it matches, the flight is "instantly grounded" in the latest system context.

### E. The "Shadow-Sync" Handshake
AI tools (Gemini-CLI, etc.) must coordinate "hand-in-glove" across machines.
- **Request for Summary**: If Flight A (Machine X) is working on a library used by Flight B (Machine Y), Flight B's AI will request a "Semantic Handshake" from the `conductor/` to understand the *intent* of the latest changes before pulling the code.

## 3. The Flight Cycle: Land and Sync
A flight is not complete until its results and **Human Stories** (Chronicles) are "Landed" (Committed/Pushed) and the **Global Registry** is updated. This ensures that the next flight—on any machine—starts with a 100% accurate context of the Federation.

---
*Distributed Intelligence for Global Solving.*
