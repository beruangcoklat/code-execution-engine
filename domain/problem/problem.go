package problem

import "github.com/beruangcoklat/code-execution-engine/pkg/lib/errorlib"

func New(rsc ResourceItf) DomainItf {
	return &Domain{
		Resource: rsc,
	}
}

func (d *Domain) GetProblems(limit, offset int) ([]Problem, error) {
	problems, err := d.Resource.GetProblemsCass(limit, offset)

	if err != nil {
		return []Problem{}, errorlib.AddTrace(err)
	}

	return problems, nil
}
