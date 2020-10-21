module main

go 1.15

replace lesson03/crawler/pkg/spider v0.0.0 => ../pkg/spider
replace lesson03/crawler/pkg/fakespider v0.0.0 => ../pkg/fakespider

require lesson03/crawler/pkg/spider v0.0.0
require lesson03/crawler/pkg/fakespider v0.0.0
