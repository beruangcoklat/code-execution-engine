package executecode

type (
	Usecase struct {
	}

	ExecuteCodePayload struct {
		Code        string `json:"code"`
		LanguageID  int    `json:"language_id"`
		Stdin       string `json:"stdin"`
		TimeLimit   int64  `json:"time_limit"`
		MemoryLimit int64  `json:"memory_limit"`
	}
)
