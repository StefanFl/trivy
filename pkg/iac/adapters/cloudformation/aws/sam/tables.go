package sam

import (
	"github.com/aquasecurity/defsec/pkg/providers/aws/sam"
	defsecTypes "github.com/aquasecurity/defsec/pkg/types"
	parser2 "github.com/aquasecurity/trivy/pkg/iac/scanners/cloudformation/parser"
)

func getSimpleTables(cfFile parser2.FileContext) (tables []sam.SimpleTable) {

	tableResources := cfFile.GetResourcesByType("AWS::Serverless::SimpleTable")
	for _, r := range tableResources {
		table := sam.SimpleTable{
			Metadata:         r.Metadata(),
			TableName:        r.GetStringProperty("TableName"),
			SSESpecification: getSSESpecification(r),
		}

		tables = append(tables, table)
	}

	return tables
}

func getSSESpecification(r *parser2.Resource) sam.SSESpecification {

	spec := sam.SSESpecification{
		Metadata:       r.Metadata(),
		Enabled:        defsecTypes.BoolDefault(false, r.Metadata()),
		KMSMasterKeyID: defsecTypes.StringDefault("", r.Metadata()),
	}

	if sse := r.GetProperty("SSESpecification"); sse.IsNotNil() {
		spec = sam.SSESpecification{
			Metadata:       sse.Metadata(),
			Enabled:        sse.GetBoolProperty("SSEEnabled"),
			KMSMasterKeyID: sse.GetStringProperty("KMSMasterKeyID"),
		}
	}

	return spec
}
