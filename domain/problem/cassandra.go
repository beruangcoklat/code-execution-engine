package problem

func (r *Resource) GetProblemsCass(limit, offset int) ([]Problem, error) {
	resultSet := r.Cassandra.QueryIter(QUERY_GET_PROBLEMS)
	problems := []Problem{}
	m := map[string]interface{}{}
	for resultSet.MapScan(m) {
		problems = append(problems, Problem{
			ID:    m["id"].(string),
			Title: m["title"].(string),
		})
		m = map[string]interface{}{}
	}

	return problems, nil
}
