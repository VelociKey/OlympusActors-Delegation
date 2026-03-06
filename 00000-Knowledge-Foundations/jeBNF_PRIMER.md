# jeBNF Primer
**Format:** Joe's Extended Backus-Naur Form
**Lineage:** eBNF + Wirth Syntax Notation (WSN) clarifications

## 1. Lineage and Purpose
jeBNF is eBNF extended with two specific Wirth Syntax Notation (WSN) clarifications to improve precision in machine-readable sovereign policy and coordination state files. It serves as the formal grammar for all metadata, configuration, and state within the Sovereign Fleet, ensuring absolute semantic clarity for both humans and AI agents.

## 2. Key Divergences from Standard eBNF

| Construct | Standard eBNF | jeBNF |
|---|---|---|
| Alternation (unordered) | `|` | `|` (retained) |
| Alternation (ordered) | not specified | `/` — first matching alternative is selected |
| Terminal designation | convention-dependent | `;` — explicit terminal marker |

## 3. Worked Example
The following is a jeBNF fragment representing a Pulse Packet (from Multi-Flight Governance Section 2C):

```jebnf
pulse_packet = flight_id track_id intent_hash current_goal ;
flight_id    = "FLIGHT-" { digit } ;
track_id     = "CYC-" { digit } ;
intent_hash  = 64 * hex_digit ;
current_goal = "\"" { character } "\"" ;

digit        = "0" / "1" / "2" / "3" / "4" / "5" / "6" / "7" / "8" / "9" ;
hex_digit    = digit / "a" / "b" / "c" / "d" / "e" / "f" ;
character    = ? all visible characters ? ;
```

**Concrete Instance:**
```jebnf
FLIGHT-001 CYC-108 5e8dc2f36573... "Establish Sovereign Governance" ;
```

## 4. File Conventions in the Fleet
- **.coord-state.jebnf**: Canonical workspace coordination hash and state tracking.
- **FLUX_LEDGER.jebnf**: Immutable flight recorder capturing all system transitions.
- **Sovereign Policy files**: Executable governance rules for access and execution control.
- **.mission.jebnf**: Thematic mission statements for every 10k block.
