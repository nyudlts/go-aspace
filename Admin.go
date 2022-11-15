package aspace

func Reindex(a *ASClient) (*int, error) {
	response, err := a.post("/plugins/reindex", true, "")
	if err != nil {
		return nil, err
	}
	return &response.StatusCode, nil
}
