// Copyright (c) 2018, Google, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package pipelinetemplate

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/spinnaker/spin/cmd/gateclient"
	"github.com/spinnaker/spin/util"
)

// PublishOptions comment
type PublishOptions struct {
	*pipelineTemplateOptions
	output       string
	templateFile string
	templateID   string
	source       string
}

var (
	publishPipelineTemplateShort   = "Publish the pipeline template"
	publishPipelineTemplateLong    = "Publish the pipeline template"
	publishPipelineTemplateExample = `
	usage: spin pipeline-template publish [options]

	Publish a pipeline template

		--file: Path to the pipeline template file
`
)

// NewPublishCmd comment
func NewPublishCmd(appOptions pipelineTemplateOptions) *cobra.Command {
	options := PublishOptions{
		pipelineTemplateOptions: &appOptions,
	}
	cmd := &cobra.Command{
		Use:     "publish",
		Aliases: []string{"publish"},
		Short:   publishPipelineTemplateShort,
		Long:    publishPipelineTemplateLong,
		Example: publishPipelineTemplateExample,
		RunE: func(cmd *cobra.Command, args []string) error {
			fflags := cmd.Flags()
			if fflags.Changed("file") == false {
				fmt.Println("No file specified")
				return nil
			}
			return publishPipelineTemplate(cmd, options)
		},
	}

	cmd.PersistentFlags().StringVarP(&options.templateFile, "file", "f", "", "Path to template file")
	cmd.PersistentFlags().StringVarP(&options.templateID, "templateID", "t", "", "Override the template ID")
	cmd.PersistentFlags().StringVarP(&options.source, "source", "", "", "Override or add the source template")
	cmd.PersistentFlags().StringVarP(&options.output, "output", "o", "", "Configure output formatting")

	return cmd
}

func publishPipelineTemplate(cmd *cobra.Command, options PublishOptions) error {
	util.UI.Info(fmt.Sprintf("publish pipeline template %s\n", options.templateFile))
	gateClient, err := gateclient.NewGateClient(cmd.InheritedFlags())
	if err != nil {
		util.UI.Error(fmt.Sprintf("%s\n", err))
		return err
	}

	template, err := util.ReadYamlFile(options.templateFile)
	if err != nil {
		return err
	}
	util.UI.Info(fmt.Sprintf("publish pipeline template %s\n", template))

	successPayload, resp, err := gateClient.PipelineTemplateControllerApi.PublishPipelineTemplateUsingPOST(
		gateClient.Context,
		options.templateID,
		template,
	)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		util.UI.Error(fmt.Sprintf("Encountered an error publishing pipeline template %s, status code: %d\n",
			options.templateFile,
			resp.StatusCode))
		return err
	}

	util.UI.JsonOutput(successPayload, util.UI.OutputFormat)
	return nil
}
