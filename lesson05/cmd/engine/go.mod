module lesson05/cmd/engine

go 1.15

replace lesson05/pkg/spider v0.0.0 => ../../pkg/spider

replace lesson05/pkg/bst v0.0.0 => ../../pkg/bst // А почему мы должны подключить модуль, который уже подключен внутри index? (без этого не работало)

replace lesson05/pkg/index v0.0.0 => ../../pkg/index

require (
	lesson05/pkg/spider v0.0.0
	lesson05/pkg/bst v0.0.0
	lesson05/pkg/index v0.0.0
)
