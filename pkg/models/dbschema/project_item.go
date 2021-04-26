package dbschema

const (
	TableProjectItem = "project_item_test"
)

const (
	PJIColProjectItemId = "project_item_id"
	PJIColResourceId    = "resource_id"
	PJIColResourceName  = "resource_name"
	PJIColResourceType  = "resource_type"
	PJIColProjectId     = "project_id"
	PJIColCreateTime    = "create_time"
)

// columns that can be search through sql '=' operator
var IndexColumnsProjectItem = []string{
	PJIColProjectItemId, PJIColResourceId, PJIColResourceName, PJIColResourceType,
}

// columns that can be search through sql 'like' operator with search_word
var SearchColumnsProjectItem = []string{
	PJIColResourceName,
}

var UpdateColumnsProjectItem = []string{
	PJIColResourceName,
}

var InsertColumnsProjectItem = []string{
	PJIColProjectItemId, PJIColResourceId, PJIColResourceName, PJIColResourceType, PJIColProjectId, PJIColCreateTime,
}

var TimeGteColumnsProjectItem = []string{
	PJIColCreateTime,
}

var TimeLteColumnsProjectItem = []string{
	PJIColCreateTime,
}
