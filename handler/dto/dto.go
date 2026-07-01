package dto

type CreateTaskRequest struct {
	Text string `json:"text"`
}

type UpdateTaskRequest struct {
	Text string `json:"text"`
}
