package class

import "github.com/goal-web/contracts"

var (
	Application  = Define(new(contracts.Application))
	Container    = Define(new(contracts.Container))
	MagicalFunc  = Define(new(contracts.MagicalFunc))
	Component    = Define(new(contracts.Component))
	Json         = Define(new(contracts.Json))
	HttpRequest  = Define(new(contracts.HttpRequest))
	HttpResponse = Define(new(contracts.HttpResponse))
	Fields       = Define(new(contracts.FieldsProvider))
	Hash         = Define(new(contracts.Hasher))
	Exception    = Define(new(contracts.Exception))

	Auth  = Define(new(contracts.Auth))
	Guard = Define(new(contracts.Guard))

	ValidatableForm = Define(new(contracts.ValidatableForm))

	Redis      = Define(new(contracts.RedisConnection))
	Cache      = Define(new(contracts.CacheStore))
	FileSystem = Define(new(contracts.FileSystem))

	Event    = Define(new(contracts.Event))
	Listener = Define(new(contracts.EventListener))

	DB           = Define(new(contracts.DBConnection))
	SqlExecutor  = Define(new(contracts.SqlExecutor))
	QueryBuilder = Define(new(contracts.QueryBuilder))
	Model        = Define(new(contracts.Model))
)
