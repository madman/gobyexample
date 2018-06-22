// Go дозволяє використовувати [_вказівники_](https://uk.wikipedia.org/wiki/Вказівник),
// що робить можливим в межах программи
// передачу за посиланням значеннь та записів.

package main

import "fmt"

// Ми спробуємо показати вказівників вказівники порівнюючи її
// з роботою зі значенням - на прикладі двох функцій -
// `zeroval` та `zeroptr`. `zeroval` отримує аргумент типу `int`
// (ціле число), отож аргумент буде передано функції "за значенням".
// `zeroval` отримує лише копію `ival` з якою функцію викликали.
func zeroval(ival int) {
    ival = 0
}

// Функція `zeroptr` аргументом потребує `*int`
// (це означає вказівник на ціле число), тобто прийматими
// лише вказівник на цей тип данних. `*iptr`, в тілі функції,
// розіменовує вказівник з його адреси до значення в цій адресі.
// Присвоєння значення розіменованому вказівнику - змінює
// значення у адресі за посиланням.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // Адресу вказівника можна друкувати таким чином.
    fmt.Println("pointer:", &i)

    // Або скориставшись дієсловом форматування `%p`.
    fmt.Printf("pointer: %p\n", &i)
}
