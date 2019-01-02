package environment

import "flag"

// Register environment flag.
func (env *Environment) RegisterFlag() {
	flag.CommandLine.Var(env, "env", `environment: development|testing|pre-production|production (default "development")`)
}
