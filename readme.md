# RESTful-API-Gorilla
___________________________________

## RESTful API реализован на Go. При разработке использовался Gorilla framework, а документация реализована при помощи go-swagger API.
___________________________________

## HTTP методы реализованные в микросервисе:
Микросервис использует JSON как формат обмена данными. RESTful API реализует такие основные HTTP методы, как **GET**, **POST** и **PUT*:
**func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request)** - реализует HTTP метод GET, в качестве ответа возвращает пользователю список всех данных из хранилища данных.
Если во время обработки данных не удалось закодировать их в формат JSON, то в качестве ответа возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)).

**func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request)** - реализует HTTP метод POST, во время выполенния запроса в хранилище данных добавляется новый элемент типа **Product**. Если во время декодирования данных из формата JSON произошла ошибка, то в качестве ответа возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)).

**func (p *Products) UpdateData(rw http.ResponseWriter, r *http.Request)** - реализует HTTP метод PUT, URI имеет вид /{id}, где {id} может принимать лишь целочисленные значений, иначе в качестве ответа на запрос будет вовзращена **ошибка 400** (http.WriteHeader(http.StatusBadRequest)).
Во время обработки запроса происходит декодирование данных из формата JSON и при возникновении ошибки возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)). Если во время перебора объектов в хранилище данных объекта с указанным {id} не было найдено, то возвращается **ошибка 404** (http.WriteHeader(http.StatusNotFound)). 

___________________________________

## Промежуточное программное обеспечение (Middleware)
Во время обработки запросов, запросы **PUT** и **POST** сначала проходят через связывающее программное обеспечение (Middleware), где данные декодируются из формата **JSON**, а далее передаются уже декодированными в следующий запрос обработчик.

**func (p *Products) MiddlewareValidationForDatatransfer(next http.Handler) http.Handler**

Возможные ошибки Middleware:
      http.StatusInternalServerError - внутренняя ошибка сервера (500)
     
