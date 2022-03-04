package integration_test

import (
	"database/sql"
	"docker-example/internal/database"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func TestWorkflow(t *testing.T) {
	if !testing.Short() {
		suite.Run(t, new(integrationSuite))
	}
}

type integrationSuite struct {
	suite.Suite
	dbURL      string
	serverAddr string
}

func (s *integrationSuite) SetupSuite() {
	s.dbURL = os.Getenv("DB_URL")
	require.NotEmpty(s.T(), s.dbURL)

	s.serverAddr = os.Getenv("SERVER_ADDR")
	require.NotEmpty(s.T(), s.serverAddr)
}

func (s *integrationSuite) setupDB(message database.Message) {
	db, err := sql.Open("postgres", s.dbURL)
	require.NoError(s.T(), err)

	query := `INSERT INTO message (id, content) VALUES ($1, $2);`

	_, err = db.Exec(query, message.ID, message.Message)
	require.NoError(s.T(), err)
}

func (s *integrationSuite) teardownDB() {
	db, err := sql.Open("postgres", s.dbURL)
	require.NoError(s.T(), err)

	_, err = db.Exec("truncate message;")
	require.NoError(s.T(), err)
}
