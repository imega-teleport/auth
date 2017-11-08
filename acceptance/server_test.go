package acceptance

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_HandleAuthHeader_Returns200(t *testing.T) {
	id := uuid.New().String()

	req, err := http.NewRequest(http.MethodGet, "http://app:8080/", nil)
	req.SetBasicAuth(id, id)
	req.Header.Add("Content-Type", "application/json")
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode)
}
