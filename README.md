<h1 align="center">Нагрузочное тестирование: Путь джедая или Кодисекция*?!</h1>
<p align="right">*жестокие эксперименты над кодом с возможным летальным исходом последнего...</p>
<img align="right" src="https://user-images.githubusercontent.com/96039786/199320129-d36fea32-81ce-4277-8050-e080df0a60bb.gif" height="100"/>
<br>
<br>

<h2 align="left">Цели:</h2>

1. Получить степень живучести своего/командного творения, т.е. при каких нагрузках выдаётся с вероятностью в 100% ожидаемый результат,
2. Получить конкретную величину краткосрочной пиковой нагрузки на имеющемся "железе"  условного заказчика проекта,
3. Получить кривую отношения величины нагрузки к проценту выдаваемых ошибок проекта.

<h2 align="left">Задачи:</h2>

1. Поиск варианта для нагрузочного тестирования (самописные тесты, плагины/утилиты, программные продукты),
2. Поиск варианта отслеживания таймингов выполнения каждой строчки кода (самописные тесты, плагины/утилиты, программные продукты),
3. По результатам тестирования - оптимизация кода (логгирование, проверки данных, запросы к БД и т.д.)  

<h2 align="left">Приступим...</h2>

<table>
  <tr>
    <th>Варианты</th>
    <th>Порог входа</th>
    <th>Нагрузка на железо</th>
    <th>Скорость работы</th>
  </tr>
  <tr>
    <td>Самописные тесты</td>
    <td>
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
</td>
    <td>
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
    <td>
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
  </tr>
   <tr>
    <td>Postman</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
     </td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
     </td>
  </tr>
  <tr>
    <td>Яндекс.Танк, JMeter и другие</td>
    <td align="center" colspan="3">
      No comments
     </td>
  </tr>
   <tr>
    <td>Grafana K6</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
  </tr>
  </table>
