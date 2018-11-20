// `Slices` або _зрізи_ є ключовим типом даних у Go,
// який являє собою набагато зручніший інтерфейс до
// послідовностей, аніж [масиви](./arrays).

package main

import "fmt"

func main() {

    // Навідміну від масивів, зрізи обмежені лише типом
    // елементів які вони зобов’язуються зберігати
    // (масиви ж, напротивагу, обмежені в кількості елементів).
    // Створити пустий зріз не нульової довжини можливо
    // використовуючи вбудовану функцію `make`. Зверніть увагу
    // на приклад створення зрізу рядків довжиною в 3 елементи
    // (усі строки ініціалізовані з нульовим значенням).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Встановнлюємо та звертаємось до значень аналогічно тому,
    // як ми це робили в масивах.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` повертає довжину зрізу.
    fmt.Println("len:", len(s))

    // На додачу до базових операцій, зрізи підтримують ще
    // кілька операцій - які роблять їх значно привабливішими, ніж масиви.
    // Перша з цих операцій - це вбудована функція `append`,
    // що повертає зріз, який зберігає одне або більше _нових значень_.
    // Зауважимо, що потрібно присвоїти значення змінні зріза,
    // що повертається з `append` - оскільки ми можемо отримати
    // новий зразок зрізу (новий `instance`).
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Зрізи також можуть бути з`copy`йовані. В цьому прикладі
    // ми створемо новий зріз `c` аналогічного типу та довжини
    // що і `s`. після чого використовуючи `copy` скопієюмо значення
    // `s` в `c`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Зрізи підтримуються оператор "зрізання" за допомогою
    // наступного синтаксису `зріз[з_ніжного_індексу:до_верхнього_індексу]`.
    // В нашому прикладі, ця операція отримає зріз з `s`, що
    // складається з елдементів `s[2]`, `s[3]` та `s[4]` і представе
    // його як `l`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Приклад зрізання усіх елеметів до (але виключаючи) `s[5]`.
    l = s[:5]
    fmt.Println("sl2:", l)

    // Приклад зрізання усьіх елементів з (але включаючи) `s[2]`.
    l = s[2:]
    fmt.Println("sl3:", l)

    // Нам дозволяється декларувати та ініціалізувати
    // значення змінних в зрізі в один рядок (так само і масиви).
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Зрізи можуть бути скомпоновані в мультимірні структури
    // даних. Довжина внутрішніх зрізів буде варюватись,
    // навідміну від мультимірних масивів.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
