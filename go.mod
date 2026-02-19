module OlympusActors-Delegation

go 1.25.7

// Local Resolution
replace Olympus2/00000-Identity-Foundations/010-Vision/900-AegisGuardian => ../Olympus2/00000-Identity-Foundations/010-Vision/900-AegisGuardian

replace Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1

replace Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal => ../Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal

replace Olympus2/00000-Identity-Foundations/P0000-pkg/pretty => ../Olympus2/00000-Identity-Foundations/P0000-pkg/pretty

replace Olympus2/00000-Identity-Foundations/P0000-pkg/text => ../Olympus2/00000-Identity-Foundations/P0000-pkg/text

replace Olympus2/30000-Federated-Services/310-Core-Registry/900-OlympusRegistry => ../Olympus2/30000-Federated-Services/310-Core-Registry/900-OlympusRegistry

replace Olympus2/30000-Federated-Services/320-Bridges/900-FederatedBridge => ../Olympus2/30000-Federated-Services/320-Bridges/900-FederatedBridge

replace Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1 => ../Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1

replace Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1/olympusv1connect => ../Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1/olympusv1connect

replace Olympus2/50000-Intelligence-Framework/520-Prompt-Synthetics => ../Olympus2/50000-Intelligence-Framework/520-Prompt-Synthetics

replace Olympus2/50000-Intelligence-Framework/530-Semantic-Cartography => ../Olympus2/50000-Intelligence-Framework/530-Semantic-Cartography

replace Olympus2/60000-Information-Storage/610-Memory-Anchors/900-MemoryAnchor => ../Olympus2/60000-Information-Storage/610-Memory-Anchors/900-MemoryAnchor

replace Olympus2/70000-Environmental-Harness/900-EnvironmentHarness => ../Olympus2/70000-Environmental-Harness/900-EnvironmentHarness

replace Olympus2/80000-System-Governance/820-Sovereign-Audit/900-SovereignAudit => ../Olympus2/80000-System-Governance/820-Sovereign-Audit/900-SovereignAudit

replace Olympus2/90000-Enablement-Labs => ../Olympus2/90000-Enablement-Labs

replace Olympus2/90000-Enablement-Labs/000-Tools/sovereign => ../Olympus2/90000-Enablement-Labs/000-Tools/sovereign

replace Olympus2/90000-Enablement-Labs/900-Forge => ../Olympus2/90000-Enablement-Labs/900-Forge

replace Olympus2/90000-Enablement-Labs/910-provisioner => ../Olympus2/90000-Enablement-Labs/910-provisioner

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-actor => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-actor

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-forge-context => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-forge-context

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-mcp-bridge => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-mcp-bridge

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-mesh => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-mesh

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-resonance => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-resonance

replace Olympus2/90000-Enablement-Labs/P0000-pkg/000-whisper => ../Olympus2/90000-Enablement-Labs/P0000-pkg/000-whisper

replace OlympusActors-Cognition => ../OlympusActors-Cognition

replace OlympusAscent => ../OlympusAscent

replace OlympusAssurance => ../OlympusAssurance

replace OlympusAtelier => ../OlympusAtelier

replace OlympusFabric => ../OlympusFabric

replace OlympusForge => ../OlympusForge

replace OlympusForge/00000-Identity-Foundations/P0000-pkg/000-fleet => ../OlympusForge/00000-Identity-Foundations/P0000-pkg/000-fleet

replace OlympusGCP-Compute => ../OlympusGCP-Compute

replace OlympusGCP-Data => ../OlympusGCP-Data

replace OlympusGCP-Events => ../OlympusGCP-Events

replace OlympusGCP-FinOps => ../OlympusGCP-FinOps

replace OlympusGCP-Firebase => ../OlympusGCP-Firebase

replace OlympusGCP-Intelligence => ../OlympusGCP-Intelligence

replace OlympusGCP-Messaging => ../OlympusGCP-Messaging

replace OlympusGCP-Observability => ../OlympusGCP-Observability

replace OlympusGCP-Storage => ../OlympusGCP-Storage

replace OlympusGCP-Vault => ../OlympusGCP-Vault

replace OlympusGrammar => ../OlympusGrammar

replace OlympusInfrastructure => ../OlympusInfrastructure

replace OlympusVision => ../OlympusVision

replace text => ../Olympus2/00000-Identity-Foundations/P0000-pkg/text

replace pretty => ../Olympus2/00000-Identity-Foundations/P0000-pkg/pretty

replace go-internal => ../Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal

replace check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1

replace gopkg.in/check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1

require (
	Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1 v0.0.0-00010101000000-000000000000
	Olympus2/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/000-olympus/000-v1/olympusv1connect v0.0.0-00010101000000-000000000000
	Olympus2/90000-Enablement-Labs/P0000-pkg/000-mesh v0.0.0-00010101000000-000000000000
	Olympus2/90000-Enablement-Labs/P0000-pkg/000-whisper v0.0.0-00010101000000-000000000000
	connectrpc.com/connect v1.19.1
	golang.org/x/net v0.50.0
	google.golang.org/protobuf v1.36.11
)

require (
	Olympus2/90000-Enablement-Labs/P0000-pkg/000-actor v0.0.0-00010101000000-000000000000 // indirect
	Olympus2/90000-Enablement-Labs/P0000-pkg/000-resonance v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/text v0.34.0 // indirect
)
