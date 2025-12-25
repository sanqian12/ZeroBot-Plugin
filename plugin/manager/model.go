package manager

type welcome struct {
	GrpID int64  `db:"gid"`
	Msg   string `db:"msg"`
}

type member struct {
	QQ int64 `db:"qq"`
	// github username
	Ghun string `db:"ghun"`
}

type blocked struct {
	GrpID int64 `db:"gid"`
	UID   int64 `db:"uid"`
	Time  int64 `db:"time"`
}
