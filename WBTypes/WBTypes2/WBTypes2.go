package main

import (
	"fmt"
)

func main() {

	/*В Go мы можем объявить сразу несколько переменных и присвоить им значения*/

	var wbTech6, wbTech7 = 1, 2
	_, _ = wbTech6, wbTech7

	wbTech6, wbTech7 = wbTech7, wbTech6
	fmt.Println("Значения переменных wbTech7, wbTech6:", wbTech7, wbTech6)

	/*Если мы попробуем присвоить в переменную  wbTech7 значения типа string,
	то IDE подстветит:
	"4"' (тип string) не может быть представлен типом int.*/

	wbTech6, wbTech7 = 3, 4 // '"4"' (type string) cannot be represented by the type int
	//wbTech6, wbTech7 = 3, "4"
	fmt.Println(wbTech6, wbTech7)

	/*В Go мы можем объявить сразу несколько переменных и присвоить им значения
	еще и таким образом */
	var (
		wbTech9  = 1
		wbTech10 = "string"
	)

	fmt.Println("Значения переменных wbTech9, wbTech10:", wbTech9, wbTech10)

	/*
		В Go обычно используется стиль написания переменных,
		называемый "camelCase" (или "lowerCamelCase"),
		где первое слово начинается с маленькой буквы,
		а каждое последующее слово начинается с заглавной буквы,
		без пробелов или подчеркиваний между словами.
		Этот стиль позволяет переменным выглядеть читаемо и единообразно.

		Примеры camelCase:

		var userName string
		var itemCount int

		Кроме того, имена переменных должны быть описательными и понятными.
		Используйте имена, которые ясно указывают на предназначение переменной
		в контексте вашей программы.
		Это поможет другим разработчикам (и вам самим в будущем) легче понять,
		что делает ваш код.

		Например, если переменная содержит количество покупок,
		вы можете назвать ее totalPrice.
		Если переменная содержит имя пользователя, назовите ее userName.
		Такие имена переменных обеспечивают ясность и читаемость вашего кода.

		Вот примеры именования переменных с учетом этих рекомендаций:


		var userName string
		var itemCount int
		var totalPrice float64
		var isLoggedIn bool

		Эти имена переменных ясно указывают на их предназначение,
		что делает код более понятным и легким для поддержки. */

}
