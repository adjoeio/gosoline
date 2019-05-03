package crud_test

import (
	"github.com/applike/gosoline/pkg/apiserver/crud"
	"github.com/applike/gosoline/pkg/db-repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListQueryBuilder_Build_IdMissing(t *testing.T) {
	metadata := db_repo.Metadata{}
	inp := &crud.Input{}

	lqb := crud.NewListQueryBuilder(metadata)
	_, err := lqb.Build(inp)

	assert.EqualError(t, err, "no primary key defined")
}

func TestListQueryBuilder_Build_DimensionMissing(t *testing.T) {
	metadata := db_repo.Metadata{}
	inp := &crud.Input{
		Filter: crud.Filter{
			Matches: []crud.FilterMatch{
				{
					Dimension: "bla",
					Operator:  "=",
					Values:    []interface{}{"blub"},
				},
			},
		},
	}

	lqb := crud.NewListQueryBuilder(metadata)
	_, err := lqb.Build(inp)

	assert.EqualError(t, err, "no list mapping found for dimension bla")
}

func TestListQueryBuilder_Build(t *testing.T) {
	metadata := db_repo.Metadata{
		PrimaryKey: "id",
		Mappings: db_repo.FieldMappings{
			"id":     db_repo.NewSimpleFieldMapping("id"),
			"bla":    db_repo.NewSimpleFieldMapping("foo"),
			"fieldA": db_repo.NewSimpleFieldMapping("fieldA"),
			"fieldB": db_repo.NewSimpleFieldMapping("fieldB"),
		},
	}

	inp := &crud.Input{
		Filter: crud.Filter{
			Matches: []crud.FilterMatch{
				{
					Dimension: "bla",
					Operator:  "=",
					Values:    []interface{}{"blub"},
				},
			},
		},
		Order: []crud.Order{
			{
				Field:     "fieldA",
				Direction: "ASC",
			},
			{
				Field:     "fieldB",
				Direction: "DESC",
			},
		},
		Page: &crud.Page{
			Offset: 0,
			Limit:  3,
		},
	}

	lqb := crud.NewListQueryBuilder(metadata)
	qb, err := lqb.Build(inp)

	assert.NoError(t, err)

	expected := db_repo.NewQueryBuilder()
	expected.Where("(((foo = ?)))", "blub")
	expected.GroupBy("id")
	expected.OrderBy("fieldA", "ASC")
	expected.OrderBy("fieldB", "DESC")
	expected.Page(0, 3)

	assert.Equal(t, expected, qb)
}

func TestListQueryBuilder_Build_ComplexFilter(t *testing.T) {
	metadata := db_repo.Metadata{
		PrimaryKey: "id",
		Mappings: db_repo.FieldMappings{
			"id":     db_repo.NewSimpleFieldMapping("id"),
			"bla":    db_repo.NewSimpleFieldMapping("foo"),
			"fieldA": db_repo.NewSimpleFieldMapping("fieldA"),
			"fieldB": db_repo.NewSimpleFieldMapping("fieldB"),
		},
	}

	inp := &crud.Input{
		Filter: crud.Filter{
			Matches: []crud.FilterMatch{
				{
					Dimension: "bla",
					Operator:  "=",
					Values:    []interface{}{"blub", "blubber"},
				},
				{
					Dimension: "fieldA",
					Operator:  "!=",
					Values:    []interface{}{1},
				},
			},
			Groups: []crud.Filter{
				{
					Matches: []crud.FilterMatch{
						{
							Dimension: "fieldB",
							Operator:  "~",
							Values:    []interface{}{"foo"},
						},
						{
							Dimension: "fieldB",
							Operator:  "=",
							Values:    []interface{}{"bar"},
						},
					},
					Bool: "or",
				},
			},
			Bool: "and",
		},
	}

	lqb := crud.NewListQueryBuilder(metadata)
	qb, err := lqb.Build(inp)

	assert.NoError(t, err)

	expected := db_repo.NewQueryBuilder()
	expected.Where("(((foo IN (?))) and ((fieldA != ?)) and (((fieldB LIKE ?)) or ((fieldB = ?))))", []interface{}{"blub", "blubber"}, 1, "%foo%", "bar")
	expected.GroupBy("id")

	assert.Equal(t, expected, qb)
}
