package integration_test

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"docker-example/internal/database"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func (s *integrationSuite) TestGetMessage() {
	tests := []struct {
		name    string
		message database.Message
		want    string
	}{
		{
			name: "happy path",
			message: database.Message{
				ID:      "163f7cb7-896b-4e9f-8818-feea162d915d",
				Message: "Successful test",
			},
			want: "A message with the id 163f7cb7-896b-4e9f-8818-feea162d915d: Successful test\n",
		},
		{
			name: "other case",
			message: database.Message{
				ID:      "163f7cb7-896b-4e9f-8818-feea162d915d",
				Message: "A different message",
			},
			want: "A message with the id 163f7cb7-896b-4e9f-8818-feea162d915d: A different message\n",
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.setupDB(tt.message)
			defer s.teardownDB()

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
			defer cancel()

			// Trigger the endpoint
			url := fmt.Sprintf("%v/message/%v", s.serverAddr, tt.message.ID)

			resp, err := http.Get(url)
			require.NoError(s.T(), err)

			// Verify that the message from the server is correct
			respBytes, err := ioutil.ReadAll(resp.Body)
			require.Equal(s.T(), tt.want, string(respBytes))

			// Verify that the record on the db is updated (seen = true)
			db, err := sql.Open("postgres", s.dbURL)
			require.NoError(s.T(), err)

			var got bool
			err = db.QueryRow(`SELECT seen FROM message WHERE id = $1`, tt.message.ID).Scan(&got)
			require.NoError(s.T(), err)

			require.Equal(s.T(), true, got)
		})
	}
}
