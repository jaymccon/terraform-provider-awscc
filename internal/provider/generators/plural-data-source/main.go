// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/terraform-provider-awscc/internal/naming"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
	"github.com/mitchellh/cli"
)

var (
	cfType           = flag.String("cftype", "", "CloudFormation resource type; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfDataSourceType = flag.String("data-source", "", "Terraform data source type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -data-source <TF-data-source-type> -cftype <CF-type> <generated-schema-file> <generated-acctests-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 || *tfDataSourceType == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	schemaFilename := args[0]
	acctestsFilename := args[1]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := NewPluralDataSourceGenerator(ui, acceptanceTestsTemplateBody, dataSourceSchemaTemplateBody, *cfType, *tfDataSourceType)

	if err := generator.Generate(destinationPackage, schemaFilename, acctestsFilename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform %s data source: %s", *tfDataSourceType, err))
		os.Exit(1)
	}
}

type PluralDataSourceGenerator struct {
	cfType           string
	tfDataSourceType string
	shared.Generator
}

func NewPluralDataSourceGenerator(ui cli.Ui, acceptanceTestsTemplateBody, schemaTemplateBody, cfType, tfDataSourceType string) *PluralDataSourceGenerator {
	return &PluralDataSourceGenerator{
		cfType:           cfType,
		tfDataSourceType: tfDataSourceType,
		Generator: shared.Generator{
			AcceptanceTestsTemplateBody: acceptanceTestsTemplateBody,
			SchemaTemplateBody:          schemaTemplateBody,
			UI:                          ui,
		},
	}
}

// Generate generates the plural data source type's factory into the specified file.
func (p *PluralDataSourceGenerator) Generate(packageName, schemaFilename, acctestsFilename string) error {
	p.Infof("generating Terraform data source code for %[1]q into %[2]q and %[3]q", p.tfDataSourceType, schemaFilename, acctestsFilename)

	// Create target directories.
	for _, filename := range []string{schemaFilename, acctestsFilename} {
		dirname := path.Dir(filename)
		err := os.MkdirAll(dirname, shared.DirPerm)

		if err != nil {
			return fmt.Errorf("creating target directory %s: %w", dirname, err)
		}
	}

	org, svc, res, err := naming.ParseCloudFormationTypeName(p.cfType)

	if err != nil {
		return fmt.Errorf("incorrect format for CloudFormation Resource Provider Schema type name: %s", p.cfType)
	}

	ds := naming.PluralizeWithCustomNameSuffix(res, "Plural")

	factoryFunctionName := string(bytes.ToLower([]byte(ds[:1]))) + ds[1:] + shared.DataSourceType

	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, ds)

	schemaDescription := fmt.Sprintf("Plural Data Source schema for %s", p.cfType)

	templateData := shared.TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		CloudFormationTypeName:       p.cfType,
		FactoryFunctionName:          factoryFunctionName,
		PackageName:                  packageName,
		SchemaDescription:            schemaDescription,
		SchemaVersion:                1,
		TerraformTypeName:            p.tfDataSourceType,
	}

	err = p.ApplyAndWriteTemplate(schemaFilename, p.SchemaTemplateBody, &templateData)

	if err != nil {
		return err
	}

	err = p.ApplyAndWriteTemplate(acctestsFilename, p.AcceptanceTestsTemplateBody, &templateData)

	if err != nil {
		return err
	}

	return nil
}

// Terraform resource schema definition.
var dataSourceSchemaTemplateBody = `
// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("{{ .TerraformTypeName }}", {{ .FactoryFunctionName }})
}

// {{ .FactoryFunctionName }} returns the Terraform {{ .TerraformTypeName }} data source type.
// This Terraform data source type corresponds to the CloudFormation {{ .CloudFormationTypeName }} resource type.
func {{ .FactoryFunctionName }}(ctx context.Context) (tfsdk.DataSourceType, error) {
	// Required for acceptance testing.
	attributes := map[string]tfsdk.Attribute {
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},		
		"ids": {
		  Description: "Set of Resource Identifiers.",
		  Type:        providertypes.SetType{ElemType:types.StringType},
		  Computed:    true,
    	},
	}

	schema := tfsdk.Schema{
		Description: "{{ .SchemaDescription }}",
		Version:     {{ .SchemaVersion }},
		Attributes:  attributes,
	}

    var opts DataSourceTypeOptions

	opts = opts.FromCloudFormationAndTerraform("{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", schema)
	
    pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "{{ .TerraformTypeName }}", "schema", hclog.Fmt("%v", schema))

	return pluralDataSourceType, nil
}
`

// Terraform acceptance tests.
var acceptanceTestsTemplateBody = `
// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package {{ .PackageName }}_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func {{ .AcceptanceTestFunctionPrefix }}DataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "{{ .CloudFormationTypeName }}", "{{ .TerraformTypeName }}", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"), 
			),
		},
	})
}
`
