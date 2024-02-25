package http

type postRequest struct {
	Name string `json:"name" binding:"required,gte=1,max=15"`
}
