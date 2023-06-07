package responsemodel

type Employee struct {
	ID           int         `json:"_id"`
	Name         string      `json:"name"`
	SupervisorID int         `json:"supervisor_id"`
	Supervisor   interface{} `json:"supervisor"`
	Subordinates []Employee  `json:"subordinates"`
}
