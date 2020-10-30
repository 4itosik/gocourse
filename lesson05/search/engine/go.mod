module lesson05/search/engine

go 1.15

replace lesson05/crawler/pkg/spider v0.0.0 => ../../crawler/pkg/spider

replace lesson05/crawler/pkg/bst v0.0.0 => ../../crawler/pkg/bst // А почему мы должны подключить модуль, который уже подключен внутри index? (без этого не работало)

replace lesson05/crawler/pkg/index v0.0.0 => ../../crawler/pkg/index

require lesson05/crawler/pkg/spider v0.0.0

require (
	lesson05/crawler/pkg/bst v0.0.0
	lesson05/crawler/pkg/index v0.0.0
)
