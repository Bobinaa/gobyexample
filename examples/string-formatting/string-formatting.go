// Go предлагает отличную поддержку форматирования строк
// с помощью `printf`. Вот несколько примеров типичных
// задач форматирования строк.

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go предлагает несколько “глаголов” созданных для
	// форматирования общих Go значений. Например,
	// это выведет инстанс нашей `point` структуры.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// Если значение является структурой, запись `%+v`
	// выведет названия полей структуры.
	fmt.Printf("%+v\n", p)

	// Вариант `%#v` печатает синтаксическое представление
	// Go, то есть фрагмент исходного кода, который будет
	// генерировать это значение.
	fmt.Printf("%#v\n", p)

	// Для вывода типа значения, используйте `%T`.
	fmt.Printf("%T\n", p)

	// Форматирование логических значений не вызывает затруднений.
	fmt.Printf("%t\n", true)

	// Возможно большое количество опций для форматирования
	// целых чисел. Используйте `%d` для стандартного,
	// десятеричного вывода.
	fmt.Printf("%d\n", 123)

	// Бинарный вывод
	fmt.Printf("%b\n", 14)

	// Вывод символа, соответсвующего заданному числу.
	fmt.Printf("%c\n", 33)

	// `%x` - шестнадцатиричное значение.
	fmt.Printf("%x\n", 456)

	// Так же есть несколько вариантов форматирования
	// чисел с плавающей точкой. Стандартный вывод `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` и `%E` приводит числло с плавающей точкой
	// к экспоненциальному представлению.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// Для стандартного вывода строк используйте `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// Для двойных ковычек как в исходниках Go, используйте `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// Так же как и с целочисленными ранее, `%x` отображает
	// строку в виде шестнадцатеричного исчисления, с двумя
	// символами вывода за каждый байт ввода.
	fmt.Printf("%x\n", "hex this")

	// Чтобы вывести ссылку на указатель, используйте `%p`.
	fmt.Printf("%p\n", &p)

	// При форматировании чисел вам часто захочется
	// контролировать ширину и точность получаемого значения.
	// Чтобы указать ширину целого числа, используйте
	// число после `%`. По-умолчанию результат будет
	// выровнен по правому краю и дополнен пробелами.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// Вы также можете указать ширину чисел с плавающей точкой,
	// также вы можете ограничить десятичную точность
	// одновременно с помощью синтаксиса `ширина.точность`.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// Для выравнивания по левому краю используйте флаг `-`.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Вы также можете контролировать ширину при форматировании
	// строк, особенно для обеспечения их выравнивания в табличном
	// выводе. Стандартное выравнивание по правому краю.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// Для выравнивания по левому краю используйте флаг `-`.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// До сих пор мы видели `Printf`, который печатает
	// отформатированную строку в `os.Stdout`. `Sprintf`
	// форматирует и возвращает строку, нигде не печатая.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Вы можете отформатировать+вывести в `io.Writers`, используя
	// `Fprintf`.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
