// В [попередньому](range) прикладі ми бачили як `for` та
// `range` забезпечують ітерацію для базових структур даних.
// Цей синтаксис також може бути використаний для
// отримання значень з каналу.

package main

import "fmt"

func main() {

    // Ми будемо ітерувати 2 значення, які ми зараз надсилаємо
    // до каналу `queue`.
    queue := make(chan string, 2)
    queue <- "раз"
    queue <- "два"
    close(queue)

    // Цей `range` ітерує через кожен елемент що він отримує
    // з `queue`. І тому що ми закрили цей канал вище ітерація
    // припиняється, після отримання усіх 2 елементів з нього.
    for elem := range queue {
        fmt.Println(elem)
    }
}
