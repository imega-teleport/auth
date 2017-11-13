package acceptance

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/imega-teleport/auth/mysql"
	"github.com/imega/dbunit"
	"github.com/stretchr/testify/assert"
)

func TestRepo_GetUser(t *testing.T) {
	id := uuid.New().String()
	setup := dbunit.WithSetup(func(tx *sql.Tx) {
		tx.Exec("TRUNCATE users")
	})
	fixtures := dbunit.WithFixtures([]func(tx *sql.Tx){
		func(tx *sql.Tx) {
			cdt := time.Now().Format("2006-01-02 15:04:05")
			edt := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
			tx.Exec("INSERT users (login, pass, created_at, expired_at, active) VALUES (?, ?, ?, ?, 1)", id, id, cdt, edt)
		},
	})
	unit, teardown := dbunit.NewUnitDB(t, dbunit.WithDSN(getDSN()), setup, fixtures)
	defer teardown()

	repo := mysql.NewRepository(mysql.WithDB(unit.DB()))
	actual, err := repo.GetUser(context.Background(), id, id)
	assert.NoError(t, err)

	assert.Equal(t, id, actual.GetLogin())
	assert.Equal(t, id, actual.GetPass())
}
