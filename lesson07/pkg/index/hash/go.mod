module lesson07/pkg/index/hash

go 1.15

replace lesson07/pkg/crawler => ../../crawler

replace lesson07/pkg/index/hash/bst => ./bst

require (
	lesson07/pkg/crawler v0.0.0-00010101000000-000000000000
	lesson07/pkg/index/hash/bst v0.0.0-00010101000000-000000000000
)
