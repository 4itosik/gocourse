module lesson04/search/engine

go 1.15

replace lesson04/crawler/pkg/spider v0.0.0 => ../../crawler/pkg/spider

replace lesson04/crawler/pkg/index v0.0.0 => ../../crawler/pkg/index

require lesson04/crawler/pkg/spider v0.0.0

require (
	lesson04/crawler/pkg/index v0.0.0
)
