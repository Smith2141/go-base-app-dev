package main

import "fmt"

func main() {
	titles := map[string][]string{
		"Что делать?":               {"планы", "думы"},
		"Где отдохнуть в выходные":  {"отдых", "планы"},
		"Кто виноват?":              {"поиск", "думы", "общество"},
		"Лучшие рестораны города":   {"еда", "отдых"},
		"С заботой о народе":        {"думы", "общество"},
		"Как попасть на дегустацию": {"еда", "планы", "отдых"},
	}
	tags := make(map[string][]string)

	// вставьте недостающий код
	for title := range titles {
		for _, tag := range titles[title] {
			tags[tag] = append(tags[tag], title)
		}
	}

	fmt.Println(len(tags["думы"]), len(tags["общество"]), tags["поиск"])
}
