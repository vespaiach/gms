package mysql

import "github.com/vespaiach/auth/pkg/adding"

var sqlCreateUser = `INSERT INTO users(username, email, hash) VALUES(?, ?, ?);`

func (st *Storage) AddUser(u adding.User) (int64, error) {
	stmt, err := st.DbClient.Prepare(sqlCreateUser)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(u.Username, u.Email, u.Hash)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

var sqlCheckUsername = `SELECT count(id) FROM users WHERE username = ?;`

func (st *Storage) IsDuplicatedUsername(username string) (bool, error) {
	rows, err := st.DbClient.Queryx(sqlCheckUsername, username)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	var count int
	if err := rows.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

var sqlCheckEmail = `SELECT count(id) FROM users WHERE email = ?;`

func (st *Storage) IsDuplicatedEmail(email string) (bool, error) {
	rows, err := st.DbClient.Queryx(sqlCheckEmail, email)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	var count int
	if err := rows.Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}