package handlers

type Employee struct {
	EmpID         string `yaml:"emp_id" json:"emp_id"`
	FirstName     string `yaml:"first_name" json:"first_name"`
	SecondName    string `yaml:"last_name" json:"last_name"`
	DefaultSalary string `yaml:"default_salary" json:"default_salary"`
	Experience    string `yaml:"experience" json:"experience"`
	Type          string `yaml:"types" json:"types"`
	Salary        string `json:"salary"`
}

