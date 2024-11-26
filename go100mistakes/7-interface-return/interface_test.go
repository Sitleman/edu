// Будьте консервативны в том, что вы делаете, и либеральны в том, что принимаете от других.
// Transmission Control Protocol (TCP, протокол управления передачей)
// Если применить эту идиому к Go, то это будет означать:
// * возврат структур вместо интерфейсов;
// * допущение использования интерфейсов, если это возможно.
//
// func LimitReader(r Reader, n int64) Reader {
// return &LimitedReader{r, n}
// }
// В этом примере функция возвращает экспортированную структуру
// io.LimitedReader. Но сигнатура функции — это интерфейс, io.Reader. В чем
// причина нарушения правила, которое мы обсуждали до сих пор? io.Reader —
// это предварительная абстракция. Это не тот уровень, который определяется
// клиентами, а навязываемый разработчиками языка, которые заранее знали, что
// этот уровень абстракции будет полезен (например, с точки зрения возможности
// переиспользования и компоновки).
package __interface_return