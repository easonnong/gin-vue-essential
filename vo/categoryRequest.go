package vo

type CategoryRequest struct {
	//Name字段必须不为空
	Name string `json:"name" binding:"required"`
}
