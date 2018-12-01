// _Фільтр рядків_ це тип програм, що отримують дані з
// [стандартного потоку введення](https://uk.wikipedia.org/wiki/Стандартні_потоки),
// обробляють їх та видають результати роботи назад (до стандартного
// потоку виведення). Одним з найвідоміших прикладів рядкових
// фільтрів є `grep` та `sed`. Ми розглянемо приклад створення
// рядкового фільтра: що поверне текст переданий до нього
// сконвертованим до верхнього регістру. Скористайтесь цим зразоком,
// для написання власного фільтра в майбутньому.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // Використання небуферизованого `os.Stdin` через буферезований сканнер
    // дає нам доступ до методу `Scan`, який пересуватиме сканер
    // до наступного токену (а ним у нас є символ переносу "\n").
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {

        // `Text` повертає поточний токен (в нашому прикладі -
        // наступна лінія вхідних данних).
        ucl := strings.ToUpper(scanner.Text())

        // Передача результату назад до стандартного потоку виводу,
        // за допомогою `fmt.Println`.
        fmt.Println(ucl)
    }

    // Переврка помилок підчас роботи `Scan`.
    // "Кінець файлу" - це очікуванна помилка і не буде
    // зарапортована `Scan` як "помилка" взагалі.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
