package entity

type Conf struct {
	SQLITE Sqlite
}

type Sqlite struct {
	FILENAME string
}
