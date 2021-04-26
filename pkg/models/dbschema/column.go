package dbschema

var InsertColumns = map[string][]string{
	TableProject:     InsertColumnsProject,
	TableProjectItem: InsertColumnsProjectItem,
}

// columns that can be search through sql 'like' operator with search_word
var SearchColumns = map[string][]string{
	TableProject:     SearchColumnsProject,
	TableProjectItem: SearchColumnsProjectItem,
}

// columns that need to be updated
var UpdateColumns = map[string][]string{
	TableProject:     UpdateColumnsProject,
	TableProjectItem: UpdateColumnsProjectItem,
}

var IndexColumns = map[string][]string{
	TableProject:     IndexColumnsProject,
	TableProjectItem: IndexColumnsProjectItem,
}

// columns that type is time and can be search through sql '>=' operator
var TimeGteColumns = map[string][]string{
	TableProject:     TimeGteColumnsProject,
	TableProjectItem: TimeGteColumnsProjectItem,
}

// columns that type is time and can be search through sql '<=' operator
var TimeLteColumns = map[string][]string{
	TableProject:     TimeLteColumnsProject,
	TableProjectItem: TimeLteColumnsProjectItem,
}
