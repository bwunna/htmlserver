package db

import (
	"SimpleServer/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DataBase struct {
	db *sql.DB
}

// constructor for data base

func NewDB(host, user, password, nameOfDB, driverName string, port int) (*DataBase, error) {
	// constructor for DataBase struct
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, nameOfDB)
	if result, err := sql.Open(driverName, params); err != nil {
		return nil, err
	} else {
		if err = result.Ping(); err != nil {
			return nil, err
		}
		dataBase := DataBase{result}
		_, err := dataBase.db.Query("DELETE FROM cache_employee;")
		if err != nil {
			return nil, err
		}
		return &dataBase, nil
	}

}

// adding employee in db

func (base *DataBase) InsertEmployee(info models.EmployeeInfo) error {
	var queryString string
	if info.Status == "" {
		queryString = fmt.Sprintf(`insert into cache_employee (email, emp_name)
		values('%s', '%s')`, info.Email, info.Name)
	} else {
		queryString = fmt.Sprintf(`insert into cache_employee (email, emp_name, current_status, dep_name, company_name, salary)
		values('%s', '%s', '%s', '%s','%s', %d)`, info.Email, info.Name, info.Status, info.DepartmentName, info.CompanyName, info.Salary)
	}

	_, err := base.db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

// deleting employee by his email

func (base *DataBase) DeleteByEmail(emails []string) error {
	queryString := fmt.Sprintf("DELETE FROM cache_employee WHERE email in (%s);", KeysInString(emails))
	_, err := base.db.Query(queryString)
	return err

}

// slice of keys in one string
// ex: "key1", "key2", "key3" => 'key1', 'key2', 'key3'

func KeysInString(keys []string) string {
	var answer string
	for _, value := range keys {
		answer += fmt.Sprintf("'%s',", value)
	}

	if len(answer) > 1 {
		answer = answer[:len(answer)-1]
	}
	return answer
}

// getting information about salary and company of the employee

func (base *DataBase) GetEmployeeInfoByEmail(email string) (*models.EmployeeInfo, error) {

	queryString := fmt.Sprintf(`select * from cache_employee where email = '%s'`, email)
	rows, err := base.db.Query(queryString)
	if err != nil {
		return nil, err
	}
	info := &models.EmployeeInfo{}

	for rows.Next() {
		err = rows.Scan(&info.Email, &info.Name, &info.Status, &info.DepartmentName, &info.CompanyName, &info.Salary)
		if err != nil {
			if email != "" {
				return info, nil
			}
			return nil, err
		}
	}

	return info, nil
}

// updating employee info

func (base *DataBase) UpdateEmployeeInfo(info *models.EmployeeInfo) error {
	if info.Status == "" {
		return nil
	}
	queryString := fmt.Sprintf(`update cache_employee set current_status = '%s', dep_name = '%s', company_name = '%s',
		salary = %d`, info.Status, info.DepartmentName, info.CompanyName, info.Salary)
	_, err := base.db.Query(queryString)
	return err

}
