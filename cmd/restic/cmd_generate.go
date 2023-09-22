package main

import (
	"time"

	"github.com/restic/restic/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var cmdGenerate = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate manual pages and auto-completion files (bash, fish, zsh, powershell)",
	Long: `
The "generate" command writes automatically generated files (like the man pages
and the auto-completion files for bash, fish and zsh).

EXIT STATUS
===========

Exit status is 0 if the command was successful, and non-zero if there was any error.
`,
	DisableAutoGenTag: true,
	RunE:              runGenerate,
}

type generateOptions struct {
	ManDir                   string
	BashCompletionFile       string
	FishCompletionFile       string
	ZSHCompletionFile        string
	PowerShellCompletionFile string
}

var genOpts generateOptions

func init() {
	cmdRoot.AddCommand(cmdGenerate)
	fs := cmdGenerate.Flags()
	fs.StringVar(&genOpts.ManDir, "man", "", "write man pages to `directory`")
	fs.StringVar(&genOpts.BashCompletionFile, "bash-completion", "", "write bash completion `file`")
	fs.StringVar(&genOpts.FishCompletionFile, "fish-completion", "", "write fish completion `file`")
	fs.StringVar(&genOpts.ZSHCompletionFile, "zsh-completion", "", "write zsh completion `file`")
	fs.StringVar(&genOpts.PowerShellCompletionFile, "powershell-completion", "", "write powershell completion `file`")
}

func writeManpages(dir string) error {
	// use a fixed date for the man pages so that generating them is deterministic
	date, err := time.Parse("Jan 2006", "Jan 2017")
	if err != nil {
		return err
	}

	header := &doc.GenManHeader{
		Title:   "restic backup",
		Section: "1",
		Source:  "generated by `restic generate`",
		Date:    &date,
	}

	Verbosef("writing man pages to directory %v\n", dir)
	return doc.GenManTree(cmdRoot, header, dir)
}

func writeBashCompletion(file string) error {
	if stdoutIsTerminal() {
		Verbosef("writing bash completion file to %v\n", file)
	}
	return cmdRoot.GenBashCompletionFile(file)
}

func writeFishCompletion(file string) error {
	if stdoutIsTerminal() {
		Verbosef("writing fish completion file to %v\n", file)
	}
	return cmdRoot.GenFishCompletionFile(file, true)
}

func writeZSHCompletion(file string) error {
	if stdoutIsTerminal() {
		Verbosef("writing zsh completion file to %v\n", file)
	}
	return cmdRoot.GenZshCompletionFile(file)
}

func writePowerShellCompletion(file string) error {
	if stdoutIsTerminal() {
		Verbosef("writing powershell completion file to %v\n", file)
	}
	return cmdRoot.GenPowerShellCompletionFile(file)
}

func runGenerate(_ *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.Fatal("the generate command expects no arguments, only options - please see `restic help generate` for usage and flags")
	}

	if genOpts.ManDir != "" {
		err := writeManpages(genOpts.ManDir)
		if err != nil {
			return err
		}
	}

	if genOpts.BashCompletionFile != "" {
		err := writeBashCompletion(genOpts.BashCompletionFile)
		if err != nil {
			return err
		}
	}

	if genOpts.FishCompletionFile != "" {
		err := writeFishCompletion(genOpts.FishCompletionFile)
		if err != nil {
			return err
		}
	}

	if genOpts.ZSHCompletionFile != "" {
		err := writeZSHCompletion(genOpts.ZSHCompletionFile)
		if err != nil {
			return err
		}
	}

	if genOpts.PowerShellCompletionFile != "" {
		err := writePowerShellCompletion(genOpts.PowerShellCompletionFile)
		if err != nil {
			return err
		}
	}

	var empty generateOptions
	if genOpts == empty {
		return errors.Fatal("nothing to do, please specify at least one output file/dir")
	}

	return nil
}
