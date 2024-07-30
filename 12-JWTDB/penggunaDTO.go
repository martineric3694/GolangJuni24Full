package main

type Pengguna struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  string `json:"admin"`
}

func checkPengguna(user Pengguna) bool {
	db := connectDB()

	var result bool = false
	var temp int

	db.QueryRow("SELECT count(1) FROM Pengguna WHERE username=? AND password =?", user.Username, user.Password).Scan(&temp)

	defer db.Close()

	if temp == 1 {
		result = true
	}
	return result

}

func getDataPengguna(user Pengguna) Pengguna {
	db := connectDB()
	defer db.Close()
	var result Pengguna

	db.QueryRow("SELECT username, isAdmin FROM Penguna WHERE username=? AND password =?", user.Username, user.Password).Scan(&result.Username, &result.IsAdmin)

	return result
}
