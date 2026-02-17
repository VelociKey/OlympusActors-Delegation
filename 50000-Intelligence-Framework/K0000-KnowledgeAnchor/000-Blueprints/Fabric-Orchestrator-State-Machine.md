# Blueprint: Fabric State-Machine Orchestrator

🌌 **Domain:** Orchestration / States
🌌 **Reference:** CYC-002-FABRIC-ORCH
🌌 **Status:** MATERIALIZING

## 1. Technical Specification

The `OlympusFabric` Orchestrator is a Go-native finite state machine (FSM) that governs the materialization lifecycle of Sovereign Node artifacts. It enforces identity-swapping at the Operating System layer to ensure strict labor/validation segregation.

### 1.1 State Definitions

| State | Identity Context | Action | Transition (Success) | Transition (Failure) |
| :--- | :--- | :--- | :--- | :--- |
| **DRAFTING** | `@personal` Delegate | Port IR/Lexer logic; Perform high-token labor. | `VALIDATION` | `RECOVERY` |
| **VALIDATION** | `@personal` Delegate | Run `fabric validate` against WSG. Enforce numeric prefixing. | `AUDIT` | `RECOVERY` |
| **AUDIT** | `@corporate` Auditor | Execute Sovereign Audit hooks; SPIFFE attestation. | `CERTIFICATION` | `RECOVERY` |
| **CERTIFICATION** | `@corporate` Auditor | Materialize artifacts into `20000-MCP-Servers`. | **COMPLETE** | `RECOVERY` |
| **RECOVERY** | `System Default` | Revert to last Baseline; Pause for 429 backoff. | `DRAFTING` | **FATAL** |

### 1.2 State Transition Logic (Go-Native)

```go
type State int

const (
	StateDrafting State = iota
	StateValidation
	StateAudit
	StateCertification
	StateRecovery
)

func (s State) String() string {
	return [...]string{"DRAFTING", "VALIDATION", "AUDIT", "CERTIFICATION", "RECOVERY"}[s]
}
```

## 2. Identity Transition Logic

Transitions between Delegate labor and Auditor validation are triggered by environment variable manipulation of the `gcloud` configuration manager.

- **Delegate Block:** `CLOUDSDK_ACTIVE_CONFIG_NAME=personal`
- **Auditor Block:** `CLOUDSDK_ACTIVE_CONFIG_NAME=corporate`

**Safety Protocol:** 
1. The Orchestrator must verify the `gcloud config configurations list` output before executing any identity-gated command.
2. No persistent secrets are stored in code; all tokens are retrieved via the OS-managed secure store.

---
*Blueprint ID: FABRIC-ORCH-001*
*Approver: Antigravity*
