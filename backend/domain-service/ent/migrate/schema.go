// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DelegationsColumns holds the columns for the "delegations" table.
	DelegationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "user", Type: field.TypeString},
		{Name: "domain_delegations", Type: field.TypeInt},
	}
	// DelegationsTable holds the schema information for the "delegations" table.
	DelegationsTable = &schema.Table{
		Name:       "delegations",
		Columns:    DelegationsColumns,
		PrimaryKey: []*schema.Column{DelegationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "delegations_domains_delegations",
				Columns:    []*schema.Column{DelegationsColumns[4]},
				RefColumns: []*schema.Column{DomainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "delegation_user_domain_delegations",
				Unique:  true,
				Columns: []*schema.Column{DelegationsColumns[3], DelegationsColumns[4]},
			},
		},
	}
	// DomainsColumns holds the columns for the "domains" table.
	DomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "fqdn", Type: field.TypeString, Unique: true},
		{Name: "owner", Type: field.TypeString},
		{Name: "approved", Type: field.TypeBool, Default: false},
	}
	// DomainsTable holds the schema information for the "domains" table.
	DomainsTable = &schema.Table{
		Name:       "domains",
		Columns:    DomainsColumns,
		PrimaryKey: []*schema.Column{DomainsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DelegationsTable,
		DomainsTable,
	}
)

func init() {
	DelegationsTable.ForeignKeys[0].RefTable = DomainsTable
}
