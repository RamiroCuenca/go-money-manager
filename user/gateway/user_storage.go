package gateway

import (
	"github.com/RamiroCuenca/go-moneyManager/common"
	"github.com/RamiroCuenca/go-moneyManager/user/models"
)

// Create an interace for UsersStorageGateway
type UserStorageGateway interface {
	// Create function that will recibe as a parameter a pointer to models.CreateUserCMD
	// and will return an other pointer to the user model or an error
	Create(cmd *models.CreateUserCMD) (*models.User, error)
}

// The db
type UserStorage struct {
	// A pointer to sql.DB()
	*common.MySqlClient
}

// CRUD Create function
// Works like a method of the db (UserStorage witch is an instance of sql.DB)
// As UsersStorageGateway receives as parameter a pointer to CreateUserCMD
// And returns a user or an error
func (s *UserStorage) Add(cmd *models.CreateUserCMD) (*models.User, error) {
	// Use the database/sql.DB.begin() method and starts a transaction to the db
	// Usually it wont return an error
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		common.Logger().Errorf("Couldnt connect to db from user Add method. Reason: %v", err)
	}

	// Execute a SQL query with the transaction
	// We dont include the id, it is generated automatically
	result, err := tx.Exec(`INSERT INTO user (username, age, gender, nationality, balance)
	values (?, ?, ?, ?, ?)`, cmd.Username, cmd.Age, cmd.Gender, cmd.Nationality, cmd.Balance)

	if err != nil {
		common.Logger().Errorf("Couldnt execute SQL query. Reason: %v", err)
		// Abort the transaction
		_ = tx.Rollback()
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		common.Logger().Errorf("Couldnt fetch last id. Reason: %v", err)
		_ = tx.Rollback() // Abort the transaction
		return nil, err
	}

	// If the transaction was a success, commit it!
	_ = tx.Commit()

	// Return the user and a nil err
	return &models.User{
		ID:          id,
		Username:    cmd.Username,
		Age:         cmd.Age,
		Gender:      cmd.Gender,
		Nationality: cmd.Nationality,
		Balance:     cmd.Balance,
	}, nil
}
