# Задание 1

Как известно, в США президент выбирается не прямым голосованием, а путем двухуровневого голосования. Сначала проводятся выборы в каждом штате и определяется победитель выборов в данном штате. Затем проводятся государственные выборы: на этих выборах каждый штат имеет определенное число голосов — число выборщиков от этого штата. На практике, все выборщики от штата голосуют в соответствии с результами голосования внутри штата, то есть на заключительной стадии выборов в голосовании участвуют штаты, имеющие различное число голосов. Вам известно за кого проголосовал каждый штат и сколько голосов было отдано данным штатом. Подведите итоги выборов: для каждого из участника голосования определите число отданных за него голосов.

## Формат ввода
Каждая строка входного файла содержит фамилию кандидата, за которого отдают голоса выборщики этого штата, затем через пробел идет количество выборщиков, отдавших голоса за этого кандидата.

## Формат вывода
Выведите фамилии всех кандидатов в лексикографическом порядке, затем, через пробел, количество отданных за них голосов.

# Задание 2 

Теорема Лагранжа утверждает, что любое натуральное число можно представить в виде суммы четырех точных квадратов. По данному числу n найдите такое представление: напечатайте от 1 до 4 натуральных чисел, квадраты которых дают в сумме данное число.

# Задание 3

Некоторый банк хочет внедрить систему управления счетами клиентов, поддерживающую следующие операции:

Пополнение счета клиента. Снятие денег со счета. Запрос остатка средств на счете. Перевод денег между счетами клиентов. Начисление процентов всем клиентам.

Вам необходимо реализовать такую систему. Клиенты банка идентифицируются именами (уникальная строка, не содержащая пробелов). Первоначально у банка нет ни одного клиента. Как только для клиента проводится операция пололнения, снятия или перевода денег, ему заводится счет с нулевым балансом. Все дальнейшие операции проводятся только с этим счетом. Сумма на счету может быть как положительной, так и отрицательной, при этом всегда является целым числом.

## Формат ввода
Входной файл содержит последовательность операций. Возможны следующие операции: DEPOSIT name sum - зачислить сумму sum на счет клиента name. Если у клиента нет счета, то счет создается. WITHDRAW name sum - снять сумму sum со счета клиента name. Если у клиента нет счета, то счет создается. BALANCE name - узнать остаток средств на счету клиента name. TRANSFER name1 name2 sum - перевести сумму sum со счета клиента name1 на счет клиента name2. Если у какого-либо клиента нет счета, то ему создается счет. INCOME p - начислить всем клиентам, у которых открыты счета, p% от суммы счета. Проценты начисляются только клиентам с положительным остатком на счету, если у клиента остаток отрицательный, то его счет не меняется. После начисления процентов сумма на счету остается целой, то есть начисляется только целое число денежных единиц. Дробная часть начисленных процентов отбрасывается.

## Формат вывода
Для каждого запроса BALANCE программа должна вывести остаток на счету данного клиента. Если же у клиента с запрашиваемым именем не открыт счет в банке, выведите ERROR.

# Задание 4

В генеалогическом древе у каждого человека, кроме родоначальника, есть ровно один родитель. Каждом элементу дерева сопоставляется целое неотрицательное число, называемое высотой. У родоначальника высота равна 0, у любого другого элемента высота на 1 больше, чем у его родителя. Вам дано генеалогическое древо, определите высоту всех его элементов.

## Формат ввода
Программа получает на вход число элементов в генеалогическом древе N. Далее следует N-1 строка, задающие родителя для каждого элемента древа, кроме родоначальника. Каждая строка имеет вид имя_потомка имя_родителя.

## Формат вывода
Программа должна вывести список всех элементов древа в лексикографическом порядке. После вывода имени каждого элемента необходимо вывести его высоту.