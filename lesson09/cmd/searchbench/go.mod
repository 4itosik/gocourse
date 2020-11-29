module lesson09/cmd/searchbench

go 1.15

replace (
	lesson09/pkg/binary => ../../pkg/binary
	lesson09/pkg/bst => ../../pkg/bst
)

require (
	lesson09/pkg/binary v0.0.0-00010101000000-000000000000
	lesson09/pkg/bst v0.0.0-00010101000000-000000000000
)
