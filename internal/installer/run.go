package installer

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cheat/cheat/internal/config"
)

// Run runs the installer
func Run(configs string, confpath string) error {

	// determine the appropriate paths for config data and (optional) community
	// cheatsheets based on the user's platform
	confdir := path.Dir(confpath)

	// create paths for community and personal cheatsheets
	community := path.Join(confdir, "/cheatsheets/community")
	personal := path.Join(confdir, "/cheatsheets/personal")

	// template the above paths into the default configs
	configs = strings.Replace(configs, "COMMUNITY_PATH", community, -1)
	configs = strings.Replace(configs, "PERSONAL_PATH", personal, -1)

	// prompt the user to download the community cheatsheets
	yes, err := Prompt(
		"Would you like to download the community cheatsheets? [Y/n]",
		true,
	)
	if err != nil {
		return fmt.Errorf("failed to prompt: %v", err)
	}

	// clone the community cheatsheets if so instructed
	if yes {
		// clone the community cheatsheets
		if err := clone(community); err != nil {
			return fmt.Errorf("failed to clone cheatsheets: %v", err)
		}

		// also create a directory for personal cheatsheets
		if err := os.MkdirAll(personal, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// the config file does not exist, so we'll try to create one
	if err = config.Init(confpath, configs); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}

	return nil
}
