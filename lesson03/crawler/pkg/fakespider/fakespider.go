package fakespider

// Crawler - краулер
type Crawler struct{}

// New - контсруктор краулера
func New() *Crawler {
	c := Crawler{}
	return &c
}

// Scan возращает жёстко закодированные данные
func (с *Crawler) Scan(url string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://habr.com":               "Страница Лучшие публикации за сутки",
		"https://habr.com/ru/news/":      "ИТ Новости на Хабре: главные новости технологий",
		"https://habr.com/ru/companies/": "Компании / Хабр",
	}

	return data, nil
}
