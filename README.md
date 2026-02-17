# OlympusFabric

**The Universal Grammar Manufacturing Plant**

OlympusFabric is the successor to the original DSLEngine, evolved for the **Plural Fabric** ecosystem. It provides a Go-native, high-performance pipeline for transforming user intent into deterministic system artifacts.

## 🚀 Key Features

*   **Grammar-First Architecture**: Every workspace is defined by a formal grammar.
*   **jeBNF Optimization**: Hard-coded support for ordered choices (`/`), eliminating ambiguity.
*   **Go-Native Materializer**: One-click workspace bootstrapping from eBNF programs.
*   **Multi-Target Generation**:
    *   **PEG**: High-speed Go parsers via Pigeon.
    *   **Data**: YAML/JSON IR exports.
    *   **Shell**: PowerShell/Bash materialization scripts.

## 🛠️ Getting Started

### 1. Build the Tool
```bash
cd 20000-MCP-Servers/OlympusFabric
go build -o fabric.exe ./cmd/fabric
```

### 2. Validate a Grammar
```bash
./fabric validate path/to/grammar.eBNF
```

### 3. Initialize a Workspace
```bash
./fabric init path/to/workspace_definition.ebnf
```

### 4. Compile to PEG
```bash
./fabric compile my_dsl.jebnf peg
```

## 🏗️ Workspace Structure
OlympusFabric follows the `.wraith` workspace standard, ensuring strict parity between AI reasoning and physical file structures.

## ⚖️ License
Proprietary - VelociKey LLC.
