module lesson08/pkg/engine

go 1.15

replace lesson08/pkg/crawler => ../crawler

replace lesson08/pkg/index => ../index

replace lesson08/pkg/index/memory => ../index/memory

require (
	lesson08/pkg/crawler v0.0.0-00010101000000-000000000000
	lesson08/pkg/index v0.0.0-00010101000000-000000000000
	lesson08/pkg/index/memory v0.0.0-00010101000000-000000000000
)
