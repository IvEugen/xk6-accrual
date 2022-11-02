<h1 align="center">Система расчёта вознаграждений «Гофермарт»</h1>

<h2 align="left">Реабилитация логгера uber-go/zap в глазах общественности:</h2>
<img align="right" src="https://user-images.githubusercontent.com/96039786/199562535-f885679f-661e-4bba-8448-98944d038f00.gif" height="100"/>

<pre><code>
func main() {
	logger, _ := zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	server.Run(logger, undo)
}</code></pre>
<br>
<pre>
Вызов записи лога в stdout
<code>
zap.L().Error("Ошибка...", zap.Error(err))
</code></pre>
<br>
<pre>
Отображение лога
<code>
{"level":"info","ts":1667409039.5448234,"caller":"config/initConfig.go:41","msg":"Config = ",
"cfg":{"Address":"localhost:8081","DSN":"user=dbuser password=xxx dbname=accrual sslmode=disable",
"Accrual":"localhost:8080"}}
</code></pre>

<h2 align="left">Gin-Gonic: волшебные пузырьки!</h2>
<img align="right" src="https://user-images.githubusercontent.com/96039786/199564582-a858c0bf-8843-43bc-b08c-44542356f58f.png" height="100"/>
<br>
Почему именно Gin:<br>
- Логика работы понятна и проста,<br>
- Локаничный код,<br>
- Пригодились в проекте расширения функционала gzip, pprof (github.com/gin-contrib)<br>
- Синтаксический сахар <code>c.Status(), c.JSON()</code> и т.д.<br>
<br>
<br>
<h1 align="center">Нагрузочное тестирование: Путь джедая или Кодисекция*?!</h1>
<p align="right">*жестокие эксперименты над кодом с возможным летальным исходом последнего...</p>
<img align="right" src="https://user-images.githubusercontent.com/96039786/199320129-d36fea32-81ce-4277-8050-e080df0a60bb.gif" height="100"/>
<br>
<br>

<h2 align="left">Цели:</h2>

1. Получить степень живучести кода, т.е. при каких максимальных нагрузках выдаётся (с вероятностью в 100%) ожидаемый результат,
2. Получить конкретную величину краткосрочной пиковой нагрузки на имеющемся "железе"  условного заказчика проекта,
3. Получить кривую отношения величины нагрузки к проценту выдаваемых ошибок проекта.

<h2 align="left">Задачи:</h2>

1. Поиск варианта для нагрузочного тестирования (собственные тесты, плагины/утилиты, программные продукты),
2. Поиск варианта отслеживания таймингов выполнения каждой строчки кода (собственные тесты, плагины/утилиты, программные продукты),
3. По результатам тестирования - оптимизация кода (логгирование, проверки данных, запросы к БД и т.д.)  

<h2 align="left">Приступим...</h2>

<table>
  <tr>
    <th>Варианты нагрузочного тестирования</th>
    <th>Порог входа</th>
    <th>Нагрузка на железо</th>
    <th>Скорость работы</th>
  </tr>
  <tr>
    <td>go-wrk</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
</td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
    <td align="center">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
      <img height="45" src="https://user-images.githubusercontent.com/96039786/199324062-54ef0e65-b1b3-4f25-adec-e0b54f6abe98.png">
    </td>
  </tr>
  <tr>
    <td>Собственные тесты</td>
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
<h2 align="left">Варианты отслеживания таймингов</h2>
<pre>
Утилита PPROF
<code>
ROUTINE ======================== github.com/etitovets/diploma_team/internal/accrual/handlers.(*Handler).AddGoods in
c:\golang\github.com\etitovets\diploma_team\internal\accrual\handlers\addgoods.go
      30ms      6.26s (flat, cum) 32.40% of Total
         .          .      8:   "github.com/etitovets/diploma_team/internal/accrual/model"
         .          .      9:   "github.com/gin-gonic/gin"
         .          .     10:   "go.uber.org/zap"
         .          .     11:)
         .          .     12:
      10ms       10ms     13:func (h *Handler) AddGoods(c *gin.Context) {
         .      280ms     14:   body, err := h.GetBodyWithGZIPCheck(c)
         .          .     15:   if err != nil {
         .          .     16:           zap.L().Error("Ошибка GetBodyWithGZIPCheck", zap.Error(err))
         .          .     17:           c.JSON(http.StatusInternalServerError, err.Error())
         .          .     18:           return
         .          .     19:   }
         .          .     20:
         .          .     21:   newGoods := model.NewGoods{}
         .      730ms     22:   if err := json.Unmarshal(body, &newGoods); err != nil {
         .          .     23:           c.JSON(http.StatusBadRequest, "JSON is not correct")
         .          .     24:           return
         .          .     25:   }
         .          .     26:
      10ms       20ms     27:   err = h.ValidatingGoods(newGoods)
         .          .     28:   if err != nil {
         .          .     29:           zap.L().Error("Ошибка ValidateGoods", zap.Error(err))
         .          .     30:           c.JSON(http.StatusBadRequest, err.Error())
         .          .     31:           return
         .          .     32:   }
         .          .     33:
         .      5.20s     34:   err = h.serv.AddGoods(newGoods)
         .          .     35:   if err != nil {
         .          .     36:           if errors.Is(err, model.ErrUniqueMatch) {
      10ms       20ms     37:                   c.Status(http.StatusConflict)
         .          .     38:                   return
         .          .     39:           }
         .          .     40:           c.JSON(http.StatusInternalServerError, err.Error())
         .          .     41:           return
         .          .     42:   }
ROUTINE ======================== github.com/etitovets/diploma_team/internal/accrual/storage.(*DBStorage).AddGoods in
c:\golang\github.com\etitovets\diploma_team\internal\accrual\storage\addgoods.go
      10ms      5.20s (flat, cum) 26.92% of Total
         .          .     34:   if err := tx.Commit(); err != nil {
         .          .     35:           zap.L().Error("Ошибка tx.Commit", zap.Error(err))
         .          .     36:           return err
         .          .     37:   } */
         .          .     38:
         .      4.78s     39:   if _, err := d.db.Exec("INSERT INTO reward (match, reward, reward_type) VALUES ($1, $2, $3)",
         .          .     40:           goods.Match,
         .       20ms     41:           goods.Reward,
         .       10ms     42:           goods.RewardType); err != nil {
         .          .     43:           if err.(*pq.Error).Code == model.ErrCodeSQLUniqueOrder {
         .      200ms     44:                   zap.L().Error("Ошибка Unique", zap.Error(err))
         .          .     45:                   return model.ErrUniqueMatch
         .          .     46:           }
         .          .     47:           zap.L().Error("Ошибка tx.Stmt", zap.Error(err))
      10ms       10ms     48:           return err
         .          .     49:   }
         .      180ms     50:   zap.L().Info("Вознаграждение добавлено в БД!", zap.String("goods", goods.Match))
         .          .     51:   return nil
         .          .     52:}</code></pre>
 <h2 align="left">Промежуточные итоги:</h2>
 <h3 align="left">Тестовый стенд:</h3>
 Proc: Xeon E5-2666 v3 @3,5GHz (10 ядер/20 потоков)<br>
 RAM: DDR4 32Gb (4х канальный режим)<br>
 SSD: m2 PCI-E 512Gb<br>
 
 
 <img align="right" src="https://user-images.githubusercontent.com/96039786/199574475-0363b2ec-8da8-4b5c-956a-6da14e65a496.gif" height="200"/>
