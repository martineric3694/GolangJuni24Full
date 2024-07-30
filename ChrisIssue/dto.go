package main

import "fmt"

func checkPengguna(user Pengguna) bool {
	db := connectDB()

	var result bool = false
	var temp int

	fmt.Println(user)
	db.QueryRow("SELECT count(1) FROM Pengguna WHERE username = ? AND password = ?", user.Username, user.Password).Scan(&temp)

	//cek log
	fmt.Println(temp, "===============")
	if temp == 1 {
		fmt.Println("MASUK TEMPPP")
		result = true
	}
	fmt.Println(result)
	return result
}

func getDataPengguna(user Pengguna) Pengguna {
	db := connectDB()
	defer db.Close()
	var result Pengguna

	db.QueryRow("SELECT username,password,isadmin FROM Pengguna WHERE username=? AND password=? AND isadmin=?", user.Username, user.Password, user.Isadmin)

	fmt.Println(result)
	return result
}
