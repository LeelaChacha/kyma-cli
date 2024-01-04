package scaffold

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/errors"
)

// Options specifies the flags for the scaffold command
type Options struct {
	Overwrite bool
	Directory string

	ModuleConfigFile   string
	ManifestFile       string
	SecurityConfigFile string
	DefaultCRFile      string

	ModuleName    string
	ModuleVersion string
	ModuleChannel string
}

var (
	errInvalidDirectory             = errors.New("provided directory does not exist")
	errModuleConfigExists           = errors.New("scaffold module config already exists. use --overwrite flag to overwrite the file")
	errModuleNameEmpty              = errors.New("--module-name flag must not be empty")
	errModuleVersionEmpty           = errors.New("--module-version flag must not be empty")
	errModuleChannelEmpty           = errors.New("--module-channel flag must not be empty")
	errManifestFileEmpty            = errors.New("--gen-manifest flag must not be empty")
	errModuleConfigEmpty            = errors.New("--module-config flag must not be empty")
	errInvalidManifestOptions       = errors.New("flag --gen-manifest cannot be set when argument --module-manifest-path provided")
	errManifestCreationFailed       = errors.New("could not generate manifest")
	errObjectsCreationFailed        = errors.New("could not generate webhook, rbac, and crd objects")
	errSecurityConfigCreationFailed = errors.New("could not generate security config")
	errDefaultCRCreationFailed      = errors.New("could not generate default CR")
	errModuleConfigCreationFailed   = errors.New("could not generate module config")
)

func (o *Options) validateDirectory() error {
	_, err := os.Stat(o.Directory)
	if errors.Is(err, os.ErrNotExist) {
		return errInvalidDirectory
	}
	absolutePath, err := filepath.Abs(o.Directory)
	if err != nil {
		return fmt.Errorf("error getting absolute file path to module directory: %w", err)
	}
	o.Directory = "/" + absolutePath
	return nil
}

func (o *Options) getCompleteFilePath(fileName string) string {
	return path.Join(o.Directory, fileName)
}
