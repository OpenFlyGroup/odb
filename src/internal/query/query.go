package query

type Plan struct{}

type Engine interface {
	Plan(stmt Statement) (*Plan, error)
	Execute(*Plan) (Result, error)
}

type Statement interface {
	stmtMarker()
}

type Result interface {
	Next() bool
	Object() ([]byte, error)
	Close() error
}
