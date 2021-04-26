package dbschema

const (
	TableProject = "project_test"
)

const (
	PJColProjectId  = "project_id"
	PJColName       = "name"
	PJColUserId     = "user_id"
	PJColCreateTime = "create_time"
	PJColUpdateTime = "update_time"
)

// columns that can be search through sql '=' operator
var IndexColumnsProject = []string{
	PJColProjectId,
}

// columns that can be search through sql 'like' operator with search_word
var SearchColumnsProject = []string{
	PJColName,
}

var UpdateColumnsProject = []string{
	PJColName,
}

var InsertColumnsProject = []string{
	PJColProjectId, PJColName, PJColUserId, PJColCreateTime, PJColUpdateTime,
}

var TimeGteColumnsProject = []string{
	PJColCreateTime,
}

var TimeLteColumnsProject = []string{
	PJColUpdateTime,
}
