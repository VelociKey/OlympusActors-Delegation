package main

import "context"
import "dagger/olympusactors-delegation/internal/dagger"

type OlympusActorsDelegation struct{}

func (m *OlympusActorsDelegation) HelloWorld(ctx context.Context) string { return "Hello from OlympusActors-Delegation!" }

func main() { dagger.Serve() }
