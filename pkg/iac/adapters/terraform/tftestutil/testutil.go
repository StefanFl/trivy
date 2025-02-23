package tftestutil

import (
	"context"
	"testing"

	"github.com/aquasecurity/defsec/pkg/terraform"
	"github.com/aquasecurity/trivy/internal/testutil"
	parser2 "github.com/aquasecurity/trivy/pkg/iac/scanners/terraform/parser"
)

func CreateModulesFromSource(t *testing.T, source, ext string) terraform.Modules {
	fs := testutil.CreateFS(t, map[string]string{
		"source" + ext: source,
	})
	p := parser2.New(fs, "", parser2.OptionStopOnHCLError(true))
	if err := p.ParseFS(context.TODO(), "."); err != nil {
		t.Fatal(err)
	}
	modules, _, err := p.EvaluateAll(context.TODO())
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}
	return modules
}
