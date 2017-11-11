package acceptance

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/imega-teleport/auth/api"
	"github.com/stretchr/testify/assert"
)

func Test_HandleAuthHeader_Returns200(t *testing.T) {
	id := uuid.New().String()

	request := auth.AuthRequest{
		Login: id,
		Pass:  id,
	}
	requestBuf := bytes.Buffer{}
	marshaler := jsonpb.Marshaler{}
	err := marshaler.Marshal(&requestBuf, &request)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "http://app:8080/api/v1/auth/basic", &requestBuf)
	req.SetBasicAuth(id, id)
	req.Header.Add("Content-Type", "application/json")
	assert.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode)
}
