// Часто нам необхідно сконвертувати часло з рядкового
// значення і ось як це робити в Go.

package main

// Конвертація з рядків в числа є можливою завдяки пакету
// `strconv`.
import "strconv"
import "fmt"

func main() {

    // `64` щш ви бачите у фунції `ParseFloat` вказує на те
    // до скількох бітів точності ми будемо конвернувати
    // це число.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // `0` в `ParseInt` вказує на те, що тип визначатиметься
    // базуючись на фомарті у рядку, а `64` це вимога до
    // результату поміститсть у 64 біти.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt` може розпізнати навіть шістнадцяткові числа.
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // Як бачите `ParseUint` також доступний.
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `Atoi` зручно використовувати для конвертації
    // числа за основою 10 (десятичного) в ціле число.
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // Якщо на вхід подано "погані данні" - ми маємо
    // можливість перехопити помилку.
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}