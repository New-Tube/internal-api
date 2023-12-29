package endpoints

import "internal-api/db"

func resetDB() error {
	conn, err := db.GetDBConnection()

	if err != nil {
		return err
	}

	result := conn.Exec("DROP TABLE comments, media_sources, reactions, users, videos")

	return result.Error
}
