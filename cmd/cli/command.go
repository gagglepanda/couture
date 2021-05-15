package cli

import (
	"github.com/pkg/errors"
	"github.com/riywo/loginshell"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path"
)

func handleDocCommand(cmd *cobra.Command) error {
	format := os.Args[2]
	switch format {
	case "man":
		return doc.GenMan(cmd, &doc.GenManHeader{Title: commandName, Section: "1"}, os.Stdout)
	case "md", "markdown":
		return doc.GenMarkdown(cmd, os.Stdout)
	default:
		return errors.Errorf("invalid documentation format: %s", format)
	}
}

func handleCompleteCommand(cmd *cobra.Command) error {
	const shellNameArgIndex = 2
	var shellName string
	if len(os.Args) > shellNameArgIndex {
		shellName = os.Args[shellNameArgIndex]
	} else {
		shellBinary, err := loginshell.Shell()
		if err != nil {
			return err
		}
		shellName = path.Base(shellBinary)
	}
	switch shellName {
	case "bash":
		return cmd.GenBashCompletion(os.Stdout)
	case "zsh":
		return cmd.GenZshCompletion(os.Stdout)
	case "fish":
		return cmd.GenFishCompletion(os.Stdout, true)
	case "powershell", "powershell.exe":
		return cmd.GenPowerShellCompletionWithDesc(os.Stdout)
	default:
		return errors.Errorf("invalid shell: %s", shellName)
	}
}

func handleLogCommand(cmd *cobra.Command) error {
	return cmd.Execute()
}

// Execute ...
func Execute() error {
	setupFlags(couture.PersistentFlags())

	if (len(os.Args) == 2 || len(os.Args) == 3) && os.Args[1] == ("complete") {
		return handleCompleteCommand(couture)
	}
	if len(os.Args) == 3 && os.Args[1] == ("doc") {
		return handleDocCommand(couture)
	}
	return handleLogCommand(couture)
}