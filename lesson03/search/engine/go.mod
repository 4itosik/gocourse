module lesson03/search/engine

go 1.15

replace lesson03/crawler/pkg/spider v0.0.0 => ../../crawler/pkg/spider
replace lesson03/crawler/pkg/fakespider v0.0.0 => ../../crawler/pkg/fakespider

require lesson03/crawler/pkg/spider v0.0.0
require lesson03/crawler/pkg/fakespider v0.0.0
