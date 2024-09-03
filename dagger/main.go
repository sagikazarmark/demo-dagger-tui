package main

import (
	"dagger/demo-dagger-tui/internal/dagger"
)

type DemoDaggerTui struct {
	// Project source directory
	//
	// +private
	Source *dagger.Directory
}

func New(
	// Project source directory.
	//
	// +defaultPath=/
	// +ignore=[".devenv", ".direnv", ".github", ".pre-commit-config.yaml"]
	source *dagger.Directory,
) (*DemoDaggerTui, error) {
	return &DemoDaggerTui{
		Source: source,
	}, nil
}

func (m *DemoDaggerTui) Build() *dagger.File {
	return dag.Go().
		WithSource(m.Source).
		Build()
}

func (m *DemoDaggerTui) Terminal() *dagger.Container {
	return dag.Container().
		WithFile("/app", m.Build()).
		Terminal(dagger.ContainerTerminalOpts{
			Cmd: []string{"/app"},
		})
}
