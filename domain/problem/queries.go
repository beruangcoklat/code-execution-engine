package problem

const (
	QUERY_GET_PROBLEMS = `select id, title from problem where id >= ? and id <= ? allow filtering;`
)
