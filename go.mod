module OlympusActors-Delegation

go 1.25.7

// Local Resolution
replace Olympus2 => ../Olympus2
replace Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1
replace Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal => ../Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal
replace Olympus2/00000-Identity-Foundations/P0000-pkg/pretty => ../Olympus2/00000-Identity-Foundations/P0000-pkg/pretty
replace Olympus2/00000-Identity-Foundations/P0000-pkg/text => ../Olympus2/00000-Identity-Foundations/P0000-pkg/text
replace OlympusActors-Cognition => ../OlympusActors-Cognition
replace OlympusAscent => ../OlympusAscent
replace OlympusAssurance => ../OlympusAssurance
replace OlympusAtelier => ../OlympusAtelier
replace OlympusFabric => ../OlympusFabric
replace OlympusForge => ../OlympusForge
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
replace github.com/mark3labs/mcp-go => ../OlympusForge/ZC0400-Sovereign-Source/mcp-go
replace text => ../Olympus2/00000-Identity-Foundations/P0000-pkg/text
replace pretty => ../Olympus2/00000-Identity-Foundations/P0000-pkg/pretty
replace go-internal => ../Olympus2/00000-Identity-Foundations/P0000-pkg/go-internal
replace check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1
replace gopkg.in/check.v1 => ../Olympus2/00000-Identity-Foundations/P0000-pkg/check.v1

require (
	Olympus2 v0.0.0-00010101000000-000000000000
	connectrpc.com/connect v1.19.1
	golang.org/x/net v0.50.0
	google.golang.org/protobuf v1.36.11
)

require golang.org/x/text v0.34.0 // indirect
