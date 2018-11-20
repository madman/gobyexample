// _Рядковий фільтр_ це узагальнений тип програм що читають
// дані з [стандартного потоку введення](https://uk.wikipedia.org/wiki/Стандартні_потоки),
// обробляють, і видають якісь похідні результати до стандартного
// потоку виведення. Программи `grep` та  `sed` є загально вживаними
// рядковими фільтрами. Розгялнемо приклад рядкового фільтра:
// що повертає скорвертовану до верхнього регістру версію тексту,
// що який він попередньо отримує. Ми может використати цей зразок,
// щоб написати власний рядковий фільтр за допомогою Go.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // Використання не буфуризованого `os.Stdin` через буферезований сканнер
    // дає нам доступ до методу `Scan` який пересуває сканер до наступного
    // токену;  яким у нас є наступна лінія нашого сканера.
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text` поверне поточний токен, в нашому прикладі наступну
        // лінію з вхідних даних.

        ucl := strings.ToUpper(scanner.Text())

        // Запис результату до стандартного потоку виводу.
        fmt.Println(ucl)
    }

    // Перевірка щодо помилок під час `Scan`. Кінець файлу - очікуванна
    // помилка не буде зарапортований `Scan` як така.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
