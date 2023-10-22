package Provider

import (
	"SimpleServer/Internal/App/Models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

type DataBase struct {
	db *sql.DB
}

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
		_, err := dataBase.db.Query("DELETE FROM public.workers;")
		if err != nil {
			return nil, err
		}
		return &dataBase, nil
	}

}

func (base *DataBase) Insert(key string) error {
	// insert in db
	if _, err := base.db.Query("INSERT INTO public.workers(user_name)" +
		fmt.Sprintf(" VALUES ('%s');", key)); err != nil {
		return err
	}
	return nil
}

func (base *DataBase) Delete(keys []string) error {
	// delete from db using where in(names...)
	keysInOneString := strings.Join(keys, ", ")
	queryString := fmt.Sprintf("DELETE FROM public.workers WHERE user_name in (%s);", keysInOneString)
	fmt.Println(queryString)
	if _, err := base.db.Query(queryString); err != nil {
		return err
	}
	return nil

}

func (base *DataBase) GetEmployeeInfo(key string) (*Models.TableData, error) {
	//get information about salary
	row, err := base.db.Query("SELECT user_name, current_status, salary, employment_time FROM public.workers" +
		fmt.Sprintf(" where user_name = '%s';", key))
	if err != nil {
		return nil, err
	}

	var employee Models.TableData = Models.TableData{}
	for row.Next() {
		err = row.Scan(&employee.Name, &employee.Status, &employee.Salary, &employee.EmploymentTime)
		if err != nil {
			return nil, err
		}
	}

	return &employee, nil
}

func (base *DataBase) AskForPromotion(key string) error {
	//asking for promotion
	employee, err := base.GetEmployeeInfo(key)
	if err != nil {
		return err
	}
	if employee.Status == "lead" {
		return fmt.Errorf("error: you're already in highest position")
	}
	_, err = base.db.Query("UPDATE public.workers SET current_status = next_status(current_status), salary = salary * 2" +
		fmt.Sprintf(" WHERE user_name = '%s' and current_status != 'lead';", key))
	if err != nil {
		return err
	}

	return nil
}
