package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	execute()
}

var (
	anchor     string
	autoIndent bool
	insertions []string
)

func execute() {
	cmd := &cobra.Command{
		Use:   "LineInsertionTransformer",
		Short: "Kustomize exec plugin that inserts lines into base manifests",
		Long: `Kustomize exec plugin that inserts lines into base manifests. This plugin inserts specified strings above
lines those include the anchor text. By default, the same indentation as lines those includes the 
anchor text will be automatically added to string to be inserted.`,
		RunE: run,
	}

	cmd.Flags().StringVar(&anchor, "anchor", "", "Insertions will be added above the lines those contains text specified by this flag")
	cmd.MarkFlagRequired("anchor")
	cmd.Flags().BoolVar(&autoIndent, "auto-indent", true, "Automatically add the same indentation as the anchor lines")
	cmd.Flags().StringArrayVar(&insertions, "insertion", nil, "Lines you want to insert")
	cmd.MarkFlagRequired("insertion")

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) error {
	if terminal.IsTerminal(0) {
		// do nothing
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, anchor) {
			var indent string
			if autoIndent {
				re := regexp.MustCompile(`^\s*`)
				indent = re.FindString(text)
			}
			for _, ins := range insertions {
				fmt.Fprintln(os.Stdout, indent+ins)
			}
		}
		fmt.Fprintln(os.Stdout, text)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
