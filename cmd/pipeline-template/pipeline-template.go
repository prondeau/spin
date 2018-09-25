package pipelinetemplate

import (
	"io"

	"github.com/spf13/cobra"
)

type pipelineTemplateOptions struct{}

var (
	pipelineTemplateShort   = ""
	pipelineTemplateLong    = ""
	pipelineTemplateExample = ""
)

// NewPipelineTemplateCmd command
func NewPipelineTemplateCmd(out io.Writer) *cobra.Command {
	options := pipelineTemplateOptions{}
	cmd := &cobra.Command{
		Use:     "pipeline-template",
		Aliases: []string{"pipeline-template"},
		Short:   pipelineTemplateShort,
		Long:    pipelineTemplateLong,
		Example: pipelineTemplateExample,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// create subcommands
	cmd.AddCommand(NewPublishCmd(options))
	return cmd
}
