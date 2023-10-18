package user

import "fmt"

func Getname() string {

	var name string

	fmt.Println("Digite o seu nome")
	fmt.Scanf("%s", &name)

	return name

}

func Getlastname() string {

	var lname string

	fmt.Println("Digite o seu sobrenome")
	fmt.Scanf("%s", &lname)

	return lname
}

func Getidade() int {

	var idade int

	fmt.Println("Digite a sua idade")
	fmt.Scanf("%d", &idade)

	return idade
}

func Getemail() string {

	var email string

	fmt.Println("Digite o seu email")
	fmt.Scanf("%s", &email)

	return email
}
