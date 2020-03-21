package handlers

type Employee struct {
	EmpID         string `yaml:"emp_id" json:"emp_id"`
	FirstName     string `yaml:"first_name" json:"firstname"`
	SecondName    string `yaml:"last_name" json:"lastname"`
	DefaultSalary string `yaml:"default_salary" json:"default_salary"`
	Experience    string `yaml:"experience" json:"experiance"`
	Type          string `yaml:"types" json:"position"`
	Salary        string `json:"salary"`
}

