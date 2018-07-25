package module

import (
	"fmt"
	"github.com/lawrencegripper/ion/internal/app/handler/constants"
	"github.com/lawrencegripper/ion/internal/app/handler/helpers"
)

// cSpell:ignore bson

// Environment represents the directory structure in
// which the module operates
type Environment struct {
	InputBlobDirPath  string
	InputMetaFilePath string

	OutputBlobDirPath   string
	OutputMetaFilePath  string
	OutputEventsDirPath string
}

// GetModuleEnvironment returns a struct that represents
// the require directory structure for the module
// environment.
func GetModuleEnvironment(baseDir string) *Environment {
	return &Environment{
		InputBlobDirPath:  helpers.GetPath(baseDir, constants.InputBlobDir),
		InputMetaFilePath: helpers.GetPath(baseDir, constants.InputEventMetaFile),

		OutputBlobDirPath:   helpers.GetPath(baseDir, constants.OutputBlobDir),
		OutputMetaFilePath:  helpers.GetPath(baseDir, constants.OutputInsightsFile),
		OutputEventsDirPath: helpers.GetPath(baseDir, constants.OutputEventsDir),
	}
}

// Build creates a clean directory structure for the module
func (m *Environment) Build() error {
	if err := helpers.CreateDirClean(m.InputBlobDirPath); err != nil {
		return fmt.Errorf("could not create input blob directory, %+v", err)
	}
	if err := helpers.CreateDirClean(m.OutputBlobDirPath); err != nil {
		return fmt.Errorf("could not create output blob directory, %+v", err)
	}
	if err := helpers.CreateFileClean(m.OutputMetaFilePath); err != nil {
		return fmt.Errorf("could not create output meta file, %+v", err)
	}
	if err := helpers.CreateDirClean(m.OutputEventsDirPath); err != nil {
		return fmt.Errorf("could not create output events directory, %+v", err)
	}
	return nil
}

// Clear will clean down the module's directory structure
func (m *Environment) Clear() error {
	if err := helpers.ClearDir(m.InputBlobDirPath); err != nil {
		return fmt.Errorf("could not create input blob directory, %+v", err)
	}
	if err := helpers.ClearDir(m.OutputBlobDirPath); err != nil {
		return fmt.Errorf("could not create output blob directory, %+v", err)
	}
	if err := helpers.RemoveFile(m.OutputMetaFilePath); err != nil {
		return fmt.Errorf("could not create output meta file, %+v", err)
	}
	if err := helpers.ClearDir(m.OutputEventsDirPath); err != nil {
		return fmt.Errorf("could not create output events directory, %+v", err)
	}
	return nil
}
