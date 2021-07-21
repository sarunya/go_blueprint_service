package schema

type Schema interface {
}

type TaskSchema struct {
	UserId       string
	Id           string
	IsActive     bool
	Task         string
	Status       string
	CreatedDate  string
	ModifiedDate string
}
