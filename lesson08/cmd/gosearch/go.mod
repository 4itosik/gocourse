module lesson08/cmd/gosearch

go 1.15

replace (
	lesson08/pkg/crawler => ../../pkg/crawler
	lesson08/pkg/crawler/membot => ../../pkg/crawler/membot
	lesson08/pkg/crawler/spider => ../../pkg/crawler/spider
	lesson08/pkg/engine => ../../pkg/engine
	lesson08/pkg/index => ../../pkg/index
	lesson08/pkg/index/memory => ../../pkg/index/memory
	lesson08/pkg/index/hash => ../../pkg/index/hash
	lesson08/pkg/index/hash/bst => ../../pkg/index/hash/bst
	lesson08/pkg/storage => ../../pkg/storage
	lesson08/pkg/storage/filestore => ../../pkg/storage/filestore
)

require (
	lesson08/pkg/crawler v0.0.0-00010101000000-000000000000
	lesson08/pkg/crawler/spider v0.0.0-00010101000000-000000000000
	lesson08/pkg/engine v0.0.0-00010101000000-000000000000
	lesson08/pkg/index v0.0.0-00010101000000-000000000000
	lesson08/pkg/index/hash v0.0.0-00010101000000-000000000000
	lesson08/pkg/storage v0.0.0-00010101000000-000000000000
	lesson08/pkg/storage/filestore v0.0.0-00010101000000-000000000000
)
