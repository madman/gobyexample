// Фукнція _recover_ (або відновлення) може бути корисна
// лише працюючи в парі з _defer_ щоб перехоплювати _panic_
// та відновлювати контроль виконання.
package main

import "fmt"

// Основна мета цієї функції викликати паніку.
func errorFunc() {
    panic("Паніка без причини")
}

// Паніка буде викликана в цій функції, і наша ціль
// перехопити цей стан та акуратно завершити роботу.
func normalFunc() {

    // Перехоплення _panic_'и та _recover_ відбувається
    // за допомогою відкладеного виклику який, як ви вже
    // знаєте відбувається останньою дією виклику функції.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Опрацювання паніки:", r)
        }
    }()

    // Момент в який програма за_panic_'ує.
    errorFunc()

    // Цей рядок ніколи не буде надруковано.
    fmt.Println("Нормальний вихід")
}

func main() {
    normalFunc()
    fmt.Println("Повернення з функції normalFunc.")
}
