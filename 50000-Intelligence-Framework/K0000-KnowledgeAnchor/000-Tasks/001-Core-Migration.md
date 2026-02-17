# Task: OlympusFabric Core Migration

**Status:** Pending
**Owner:** Antigravity
**Dependencies:** ADR-000

## Action Items
1.  [ ] **Repo Sync**: Clone `VelociKey/DSLEngine` as `xxDSLEngine` (✅ Complete).
2.  [ ] **Module Init**: Initialize `github.com/VelociKey/OlympusFabric` Go module.
3.  [x] **IR Port**: Migrate `OlympusGrammar/pkg/ir` to `OlympusFabric/pkg/ir` (✅ Complete).
4.  [x] **Lexer Port**: Port optimized Lexer from `internal/parser/lexer.go` (✅ Complete).
5.  [x] **Parser Fusion**: Merge legacy `DSLEngine/src/parser` logic with the new `internal/parser` (✅ Complete).
6.  [x] **CLI Wrapper**: Implement the `fabric` CLI tool (✅ Complete).
7.  [x] **Generator Foundation**: Implement the Generator Registry and PEG/Data targets (✅ Complete).

## Verification
- [ ] Successfully parse `test.jebnf` from legacy codebase.
- [ ] Generate a valid JSON IR from a complex calculator grammar.
