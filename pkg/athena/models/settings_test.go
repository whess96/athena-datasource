package models

import (
	"testing"

	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
)

func TestConnection_getRegionKey(t *testing.T) {
	tests := []struct {
		description string
		settings    *AthenaDataSourceSettings
		region      string
		catalog     string
		database    string
		expected    string
	}{
		{
			description: "undefined region",
			settings:    &AthenaDataSourceSettings{AWSDatasourceSettings: awsds.AWSDatasourceSettings{}},
			expected:    "default-default-default",
		},
		{
			description: "default region",
			settings:    &AthenaDataSourceSettings{AWSDatasourceSettings: awsds.AWSDatasourceSettings{}},
			region:      "default",
			expected:    "default-default-default",
		},
		{
			description: "same region",
			settings: &AthenaDataSourceSettings{
				AWSDatasourceSettings: awsds.AWSDatasourceSettings{
					DefaultRegion: "foo",
				},
			},
			region:   "foo",
			expected: "default-default-default",
		},
		{
			description: "different region",
			settings: &AthenaDataSourceSettings{
				AWSDatasourceSettings: awsds.AWSDatasourceSettings{
					Region: "foo",
				},
			},
			region:   "foo",
			expected: "foo-default-default",
		},
		{
			description: "different catalog",
			settings: &AthenaDataSourceSettings{
				AWSDatasourceSettings: awsds.AWSDatasourceSettings{
					Region: "foo",
				},
			},
			catalog:  "foo",
			expected: "default-foo-default",
		},
		{
			description: "different database",
			settings: &AthenaDataSourceSettings{
				AWSDatasourceSettings: awsds.AWSDatasourceSettings{
					Region: "foo",
				},
			},
			database: "foo",
			expected: "default-default-foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			if res := tt.settings.GetConnectionKey(tt.region, tt.catalog, tt.database); res != tt.expected {
				t.Errorf("unexpected result %v expecting %v", res, tt.expected)
			}
		})
	}
}