package databaseminer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoMiner struct {
	Host    string
	session *mgo.Session
}

func NewMongoMiner(host string) (*MongoMiner, error) {
	m := MongoMiner{Host: host}
	err := m.Connect()
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *MongoMiner) Connect() error {
	session, err := mgo.Dial(m.Host)
	if err != nil {
		return err
	}
	m.session = session
	return nil
}

func (m *MongoMiner) GetSchema() (*Schema, error) {
	s := new(Schema)

	dbNames, err := m.session.DatabaseNames()
	if err != nil {
		return nil, err
	}

	for _, dbName := range dbNames {
		db := Database{Name: dbName}
		collections, err := m.session.DB(dbName).CollectionNames()
		if err != nil {
			return nil, err
		}

		for _, collection := range collections {
			table := Table{Name: collection}

			var docRaw bson.Raw
			if err := docRaw.Unmarshal(&table); err != nil {
				return nil, err
			}

			var doc bson.RawD
			if err := docRaw.Unmarshal(&doc); err != nil {
				return nil, err
			}

			for _, f := range doc {
				table.Columns = append(table.Columns, f.Name)
			}
			db.Tables = append(db.Tables, table)
		}

		s.Databases = append(s.Databases, db)
	}

	return s, nil
}
