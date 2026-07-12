package dto

type CreateTaskRequest struct {
	Text string `json:"text"`
}

type UpdateTaskRequest struct {
	Text string `json:"text"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
