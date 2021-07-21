package queries

type TaskQuery string

const (
	Select TaskQuery = "SELECT * from task"
	Delete TaskQuery = "DELETE from task where id=$1"
	Insert TaskQuery = `INSERT into task (id, user_id, task, status, is_active, created_date, modified_date)
	VALUES ($1, $2, $3, $4, "true")`
)
