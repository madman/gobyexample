// Стректури в Go - це колекції полів з визначиним типом.
// Нaдзвичайну користність структур можна оцінити не тільки для
// групування даних, а й тому що вони служать основним рушієм мови
// що орієнтується на данні - на структури.

package main

import "fmt"

// Структура `person` має поля для імені (`name`) та віку (`age`).
type person struct {
    name string
    age  int
}

func main() {

    // Наступний синтаксис створить нову структуру.
    // Дозволяється іменувати поля при ініціалізації структури.
    fmt.Println(person{name: "Alice", age: 30})

    // А можна і уникнути, іменованої ініціалізації.
    fmt.Println(person{"Bob", 20})

    // Пропущені поля будуть створені відповідно нульового.
    // значення типу поля, що ми оминаємо
    fmt.Println(person{name: "Fred"})

    // Префікс `&` поверне вказівник на структуру.
    fmt.Println(&person{name: "Ann", age: 40})

    // Доступ до полів надається через синстксис крапки `.`
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // Дозволяється використовувати крапки з вказівниками структур,
    // вказівники, в такому разі, автоматично розіменовуються.
    sp := &s
    fmt.Println(sp.age)

    // Дані у структурі можна змінювати (тобто вони `mutable`,
    // такі що дозволяється змінювати).
    sp.age = 51
    fmt.Println(sp.age)
}
