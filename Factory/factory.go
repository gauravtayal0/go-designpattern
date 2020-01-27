package Factory

type DatabaseType int

const (
	Mongo DatabaseType = iota
	Sql
)

type (
	mongoDB struct {
		database map[string]string
	}

	sql struct {
		database map[string]string
	}
)

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}
	return mdb.database[query]
}

func (mdb mongoDB) PutData(query, data string) {
	mdb.database[query] = data + " mongoDb"
}

func (sql sql) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}
	return sql.database[query]
}

func (sql sql) PutData(query, data string) {
	sql.database[query] = data + " sql"
}

func DatabaseFactory(database DatabaseType) Database {
	switch database {
	case Mongo:
		return mongoDB{
			database: make(map[string]string, 0),
		}
	case Sql:
		return sql{
			database: make(map[string]string, 0),
		}
	default:
		return nil

	}
}
