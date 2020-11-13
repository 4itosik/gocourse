module lesson07/cmd/gosearch

go 1.15

replace (
	lesson07/pkg/crawler => ../../pkg/crawler
	lesson07/pkg/crawler/membot => ../../pkg/crawler/membot
	lesson07/pkg/crawler/spider => ../../pkg/crawler/spider
	lesson07/pkg/index => ../../pkg/index
	lesson07/pkg/index/hash => ../../pkg/index/hash
	lesson07/pkg/index/hash/bst => ../../pkg/index/hash/bst // А почему надо и подзависимости указывать?
	lesson07/pkg/storage => ../../pkg/storage
	lesson07/pkg/storage/filestore => ../../pkg/storage/filestore
)

require (
	lesson07/pkg/crawler v0.0.0-00010101000000-000000000000
	lesson07/pkg/crawler/spider v0.0.0-00010101000000-000000000000
	lesson07/pkg/index v0.0.0-00010101000000-000000000000
	lesson07/pkg/index/hash v0.0.0-00010101000000-000000000000
	lesson07/pkg/storage v0.0.0-00010101000000-000000000000
	lesson07/pkg/storage/filestore v0.0.0-00010101000000-000000000000
)
