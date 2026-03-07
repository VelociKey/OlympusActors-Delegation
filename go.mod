module olympus.fleet/00SDLC/OlympusActors-Delegation

go 1.26.0

// Local Resolution

require (
	connectrpc.com/connect v1.19.1
	golang.org/x/net v0.51.0
	google.golang.org/protobuf v1.36.11
)

require golang.org/x/text v0.34.0 // indirect

// --- Local Sovereign Resolution ---

replace olympus.fleet/00SDLC/Olympus2/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc => ../Olympus2/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc

replace olympus.fleet/00SDLC/Olympus2/50000-Intelligence-Framework/50200-Logic-Libraries => ../Olympus2/50000-Intelligence-Framework/50200-Logic-Libraries

replace olympus.fleet/00SDLC/Olympus2/70000-Environmental-Harness/dagger => ../Olympus2/70000-Environmental-Harness/70700-Harness-Drivers/dagger-70000

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries

replace olympus.fleet/00SDLC/OlympusFabric/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc => ../OlympusFabric/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc

replace olympus.fleet/00SDLC/OlympusGrammar/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGrammar/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Data/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Data/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Firebase/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Firebase/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Storage/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Storage/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Vault/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Vault/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Events/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Events/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Observability/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Observability/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-FinOps/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-FinOps/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Intelligence/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Intelligence/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusGCP-Compute/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen => ../OlympusGCP-Compute/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen

replace olympus.fleet/00SDLC/OlympusMCP/P0300-Mesh/100-Ground-Substrate => ../OlympusMCP/P0300-Mesh/100-Ground-Substrate

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/110-Auth => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/110-Auth

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/120-Econotel => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/120-Econotel

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/130-ForgeContext => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/130-ForgeContext

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/140-MCPBridge => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/140-MCPBridge

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/150-Mesh => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/150-Mesh

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/170-Policy => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/170-Policy

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/190-Search => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/190-Search

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/200-Substrate => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/200-Substrate

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/210-Vault => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/210-Vault

replace olympus.fleet/00SDLC/Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/220-Whisper => ../Olympus2/90000-Enablement-Labs/90200-Logic-Libraries/220-Whisper

replace olympus.fleet/00SDLC/OlympusGrammar/parser => ../OlympusGrammar/parser

replace olympus.fleet/10GNDT/GND-Clearinghouse/01000-Identity-Foundations/020-Hardware => ../10GNDT/GND-Clearinghouse/01000-Identity-Foundations/020-Hardware

replace olympus.fleet/10GNDT/GND-Clearinghouse/P0000-pkg/clearing => ../10GNDT/GND-Clearinghouse/P0000-pkg/clearing

replace olympus.fleet/10GNDT/GND-Clearinghouse/P0000-pkg/revenue => ../10GNDT/GND-Clearinghouse/P0000-pkg/revenue

replace olympus.fleet/10GNDT/GND-Customs/01000-Identity-Foundations/020-Hardware => ../10GNDT/GND-Customs/01000-Identity-Foundations/020-Hardware

replace olympus.fleet/10GNDT/GND-Customs/P0000-pkg/compliance => ../10GNDT/GND-Customs/P0000-pkg/compliance

replace olympus.fleet/10GNDT/GND-Freight/P0000-pkg/bale => ../10GNDT/GND-Freight/P0000-pkg/bale

replace olympus.fleet/10GNDT/GND-Freight/P0000-pkg/logistics => ../10GNDT/GND-Freight/P0000-pkg/logistics

replace olympus.fleet/10GNDT/GND-Freight/P0000-pkg/satchel => ../10GNDT/GND-Freight/P0000-pkg/satchel

replace olympus.fleet/10GNDT/GND-Registry/P0000-pkg/indexing => ../10GNDT/GND-Registry/P0000-pkg/indexing

replace olympus.fleet/10GNDT/GND-Substrate/01000-Identity-Foundations/020-Hardware => ../10GNDT/GND-Substrate/01000-Identity-Foundations/020-Hardware

replace olympus.fleet/30INFR/Pinnacle/01000-Identity-Foundations/P0000-pkg/generate => ../30INFR/Pinnacle/01000-Identity-Foundations/P0000-pkg/generate

replace olympus.fleet/30INFR/Pinnacle/01000-Identity-Foundations/P0000-pkg/parse => ../30INFR/Pinnacle/01000-Identity-Foundations/P0000-pkg/parse
