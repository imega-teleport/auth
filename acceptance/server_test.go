package acceptance

import (
	"bytes"
	"database/sql"
	"net/http"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/imega-teleport/auth/api"
	"github.com/imega/dbunit"
	"github.com/stretchr/testify/assert"
)

func Test_Auth(t *testing.T) {
	id := uuid.New().String()
	t.Run("Valid account", func(t *testing.T) {
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
		_, teardown := dbunit.NewUnitDB(t, dbunit.WithDSN(getDSN()), setup, fixtures)
		defer teardown()

		request := auth.AuthRequest{
			Login: id,
			Pass:  id,
		}
		requestBuf := bytes.Buffer{}
		marshaler := jsonpb.Marshaler{}
		err := marshaler.Marshal(&requestBuf, &request)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, getAPIEntryPoint("basic"), &requestBuf)
		req.Header.Add("Content-Type", "application/json")
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("Account expired", func(t *testing.T) {
		setup := dbunit.WithSetup(func(tx *sql.Tx) {
			tx.Exec("TRUNCATE users")
		})
		fixtures := dbunit.WithFixtures([]func(tx *sql.Tx){
			func(tx *sql.Tx) {
				cdt := time.Now().Format("2006-01-02 15:04:05")
				edt := time.Now().AddDate(0, 0, -1).Format("2006-01-02 15:04:05")
				tx.Exec("INSERT users (login, pass, created_at, expired_at, active) VALUES (?, ?, ?, ?, 1)", id, id, cdt, edt)
			},
		})
		_, teardown := dbunit.NewUnitDB(t, dbunit.WithDSN(getDSN()), setup, fixtures)
		defer teardown()

		request := auth.AuthRequest{
			Login: id,
			Pass:  id,
		}
		requestBuf := bytes.Buffer{}
		marshaler := jsonpb.Marshaler{}
		err := marshaler.Marshal(&requestBuf, &request)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, getAPIEntryPoint("basic"), &requestBuf)
		req.Header.Add("Content-Type", "application/json")
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusForbidden, res.StatusCode)
	})

	t.Run("Wrong password", func(t *testing.T) {
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
		_, teardown := dbunit.NewUnitDB(t, dbunit.WithDSN(getDSN()), setup, fixtures)
		defer teardown()

		request := auth.AuthRequest{
			Login: id,
			Pass:  "wrong-pass",
		}
		requestBuf := bytes.Buffer{}
		marshaler := jsonpb.Marshaler{}
		err := marshaler.Marshal(&requestBuf, &request)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, getAPIEntryPoint("basic"), &requestBuf)
		req.Header.Add("Content-Type", "application/json")
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusForbidden, res.StatusCode)
	})
}

func Test_CreateUser(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, getAPIEntryPoint("user/create"), nil)
	req.Header.Add("Content-Type", "application/json")
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
