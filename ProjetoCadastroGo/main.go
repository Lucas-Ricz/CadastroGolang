package main

import (
	"database/sql"
	"fmt"
	"projetoCadastro/Internal/config"
	menu "projetoCadastro/Internal/interface"
	user "projetoCadastro/Internal/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var err error
	var n int = 0

	db_con := fmt.Sprintf("%s:%s@(%s:%s)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_DATABASE)

	db, err = sql.Open("mysql", db_con)
	if err != nil {
		panic(err)
	}

	for ok := true; ok; ok = (n < 4) {

		menu.Menu()
		fmt.Scanf("%d", &n)

		if n == 1 {
			user := User{
				Name:     user.Getname(),
				Lastname: user.Getlastname(),
				Idade:    user.Getidade(),
				Email:    user.Getemail(),
			}

			if insertError := insertUser(user); insertError != nil {
				panic(err)
			}

			fmt.Println("Os seus dados foram cadastrados com sucesso!")

		}

		if n == 2 {

			users, err := getAllUsers()
			if err != nil {
				panic(err)
			}

			fmt.Println("\nTabela user")
			fmt.Print("Name LastName Idade Email\n\n")

			for _, user := range users {
				fmt.Println(*user)
			}

		}

		if n == 3 {

			if deleteError := deletetUserWithName(); deleteError != nil {
				panic(err)
			}

		}

	}

}

var (
	db *sql.DB
)

type User struct {
	Name     string
	Lastname string
	Idade    int
	Email    string
}

func getAllUsers() ([]*User, error) {
	res, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	users := []*User{}

	for res.Next() {

		var user User

		if err := res.Scan(&user.Name, &user.Lastname, &user.Idade, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func insertUser(user User) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO user VALUES('%s', '%s', %d, '%s')", user.Name, user.Lastname, user.Idade, user.Email))
	if err != nil {
		return err
	}

	fmt.Println("Usuario inserido com sucesso")
	return nil
}

func deletetUserWithName() error {

	var nome string

	fmt.Println("Digite o nome do usuario a ser deletado")
	fmt.Scanf("%s", &nome)

	_, err := db.Exec(fmt.Sprintf("DELETE FROM user WHERE name = '%s'", nome))
	if err != nil {
		return err
	}

	fmt.Println("Usuario deletado com sucesso")
	return nil
}


