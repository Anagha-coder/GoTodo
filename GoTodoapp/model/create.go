package model

import "fmt"

func CreateTodo(name, todo string) error {

	insertQ, err := con.Query("INSERT INTO TODO VALUES (?,?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM TODO WHERE NAME=?", name)
	defer deleteQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
