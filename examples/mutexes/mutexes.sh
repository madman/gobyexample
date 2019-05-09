# Запуск програми покаже - що ми приблизно запустили
# біля 90 000 операцій всього що працювали з нашою
# `mutex`- синхронізованою мапою `state`.
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]

# В наступному прикладі ми спробуємо керувати станом,
# лише, за допомогою горутин та канналів.
