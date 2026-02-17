# Operational Guide: Fabric State-Machine Orchestrator

🌌 **Domain:** Mesh Operations
🌌 **Status:** ACTIVE
🌌 **Standard:** FABRIC-ORCH-OPS-001

## 1. Identity Establishment

The orchestrator relies on Operating System-managed `gcloud` configurations. Before execution, the operator must initialize both the Personal and Corporate identities in the OS secure store.

### 1.1 Configuration Setup
Execute the following commands in the host terminal:

**Personal Identity Pool (@personal):**
The orchestrator supports a rotation pool of delegate identities to distribute token load. Initialize as many as required:

```bash
gcloud config configurations create personal-1
gcloud config set account <personal_email_1>

gcloud config configurations create personal-2
gcloud config set account <personal_email_2>

# ... etc
```

**Corporate Prime Ego (@corporate):**
The authoritative context for validation and certification.
```bash
gcloud config configurations create corporate-prime
gcloud config set account <corporate_email>
```

### 1.2 Verification
The orchestrator will automatically verify the active configuration using:
```bash
gcloud config configurations list --filter "IS_ACTIVE=true"
```

---

## 2. Expected Breakpoints & Recovery

The State-Machine is designed to handle failures gracefully by entering the **RECOVERY** state.

### 2.1 Audit Failure
If the `internal/parser/audit.go` logic detects **Identity DNA mismatch** or **Security Breaches**:
1. The orchestrator transitions to `StateRecovery`.
2. It triggers `Materializer.Rollback()`.
3. The workspace is reverted to the last known-good **Corporate Baseline**.
4. The mission is paused; manual intervention is required to clean the artifacts.

### 2.2 Rate Limiting (HTTP 429)
The execution loop handles high-traffic errors:
1. Detects `429` status.
2. **Rotates to the next identity in the `@personal` pool** to distribute token load.
3. Implements a 30-second backoff.
4. Attempts to resume from `StateDrafting`.

---

## 3. Proactive Quota Management

To minimize "429 Surprise" (Rate Limiting), the orchestrator implements a **Proactive Pre-flight Check**.

1.  **Stats Hook**: Before entering the `DRAFTING` or `VALIDATION` states, the orchestrator executes the built-in `gemini /stats` command.
2.  **Quota Detection**: It parses the output for indicators of token exhaustion or "low quota" status.
3.  **Preventative Rotation**: If a quota risk is detected, the system rotates the `@personal` identity *before* starting the work, ensuring a high-success probability for the task.

---

## 4. Grammar Invariant Enforcement
*Generated: 2026-02-14*
*Approver: Antigravity*
