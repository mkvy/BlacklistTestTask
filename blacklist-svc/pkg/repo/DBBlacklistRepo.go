package repo

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/utils"
	"log"
)

type DBBlacklistRepo struct {
	db *sqlx.DB
}

func NewDBBlacklistRepo(dbConn *sqlx.DB) *DBBlacklistRepo {
	return &DBBlacklistRepo{db: dbConn}
}

func (d *DBBlacklistRepo) Create(data models.BlacklistData) error {
	log.Println("DBBlacklistRepo: creating data with id ", data.Id)
	query := `INSERT INTO blacklist VALUES(:id,:phone_number,:user_name,:ban_reason,:date_banned,:username_who_banned);`
	_, err := d.db.NamedExec(query, &data)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBBlacklistRepo) Delete(id string) error {
	log.Println("DBBlacklistRepo: deleting data with id ", id)
	res, err := d.db.Exec(`delete from blacklist where id=$1;`, id)
	if err != nil {
		log.Println("DBBlacklistRepo: error occurred while deleting record: ", err)
		return err
	}
	rowsC, _ := res.RowsAffected()
	if rowsC == 0 {
		log.Println("DBBlacklistRepo: record not found id: ", id)
		return utils.ErrNotFound
	}
	return nil
}

func (d *DBBlacklistRepo) GetByPhoneNumber(phone string) ([]models.BlacklistData, error) {
	log.Println("DBBlacklistRepo: GetByPhoneNumber phone: ", phone)
	bData := []models.BlacklistData{}
	err := d.db.Select(&bData, `select * from blacklist where phone_number=$1`, phone)
	if err != nil {
		log.Println("DBBlacklistRepo: error while GetByPhoneNumber: ", err)
		return nil, err
	}
	return bData, nil
}

func (d *DBBlacklistRepo) GetByUsername(username string) ([]models.BlacklistData, error) {
	log.Println("DBBlacklistRepo: GetByUsername username: ", username)
	bData := []models.BlacklistData{}
	err := d.db.Select(&bData, `select * from blacklist where user_name=$1`, username)
	if err != nil {
		log.Println("DBBlacklistRepo: error while GetByUsername: ", err)
		return nil, err
	}
	return bData, nil
}
