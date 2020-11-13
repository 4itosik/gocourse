module lesson06/cmd/engine

go 1.15

replace lesson06/pkg/spider v0.0.0 => ../../pkg/spider

replace lesson06/pkg/bst v0.0.0 => ../../pkg/bst // А почему мы должны подключить модуль, который уже подключен внутри index? (без этого не работало)

replace lesson06/pkg/index v0.0.0 => ../../pkg/index

require (
	lesson06/pkg/spider v0.0.0
	lesson06/pkg/bst v0.0.0
	lesson06/pkg/index v0.0.0
)
