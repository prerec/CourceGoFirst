## Задача L
**Условие:**
Часто, когда нам требуется зарегистрироваться на каком-нибудь сервисе, требуется ввести логин и почту. Насколько вам известно, на логин и адрес электронной налагаются определенные требования.

Напишите программу, которая проверяет, что все хорошо.
````
Логин считается валидным (приемлимым), если он содержит не меньше 10 символов и не содержит символа "@".

Почта считается валидной - если она содержит символ "@" , а также символ точки ".".

Ваша программа должна реагировать на введенные пользователем данные следующим образом:

если логин короче 10 символов или содержит "@" - вывести "Некорректный логин" и завершиться.
если логин валиден, а адрес почты не содержит "@" или "." - вывести "Некорректная почта" и завершиться.
если и логин и почта - валидны - вывести "ОК" (кириллица) и завершиться.
````