# RESTful-API-Gorilla
___________________________________

## RESTful API реализован на Go. При разработке использовался Gorilla framework, а документация реализована при помощи go-swagger API.
___________________________________

## HTTP методы реализованные в микросервисе:
Микросервис использует JSON как формат обмена данными. RESTful API реализует такие основные HTTP методы, как **GET**, **POST** и **PUT**:

      func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request)
Реализует HTTP метод GET, в качестве ответа возвращает пользователю список всех данных из хранилища данных.

Если во время обработки данных не удалось закодировать данные в формат JSON, то в качестве ответа возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)).

     func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request)
Реализует HTTP метод POST, во время выполенния запроса в хранилище данных добавляется новый элемент типа **Product**.

Если во время декодирования данных из формата JSON произошла ошибка, то в качестве ответа возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)).

      func (p *Products) UpdateData(rw http.ResponseWriter, r *http.Request)
Реализует HTTP метод PUT, где URI имеет вид /{id} и {id} может принимать лишь целочисленные положительные значения.

Если {id} не удовлетворяет условиям, то в качестве ответа на запрос будет вовзращена **ошибка 400** (http.WriteHeader(http.StatusBadRequest)).

Во время обработки запроса происходит декодирование данных из формата JSON и при возникновении ошибки возвращается **ошибка 500** (http.WriteHeader(http.StatusInternalServerError)). 

Если во время перебора объектов в хранилище данных объекта с указанным {id} не было найдено, то возвращается **ошибка 404** (http.WriteHeader(http.StatusNotFound)). 

___________________________________

## Связующее программное обеспечение (Middleware)
Во время обработки запросов, запросы **PUT** и **POST** сначала проходят через связующее программное обеспечение (Middleware), где данные декодируются из формата **JSON**, а далее передаются уже декодированными в следующий запрос обработчик.

     func (p *Products) MiddlewareValidationForDatatransfer(next http.Handler) http.Handler

Возможные ошибки Middleware:

      http.StatusInternalServerError - внутренняя ошибка сервера (500)
      
## Документация

Документация реализована при помощи **go-swagger API**, расположенная по ссылке: https://github.com/go-swagger/go-swagger.
В проект добавлен Makefile, выполнение которого создаст файл, содержащий данные для отображения документации. Документация отображается на локальном хосте **(localhost:9090/docs)**. Для того, чтобы выполнить Makefile, в корневой директории проекта пропишите:

     make swagger
     
В качестве результата будет создан файл swagger.yaml необходимый для отображения документации на локальном хосте.
Если **swagger-API** не установлена, то Makefile загрузит его самостоятельно. Результат выполнения Makefile:

     which swagger || GO111MODULE=off go get github.com/go-swagger/go-swagger/cmd/swagger
     /home/ellofae/go/bin/swagger
     GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models

___________________________________

### Результат работы метода GET:
![result1](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/get-method.PNG?raw=true)


### Результат работы метода PUT
![result1](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/put-method.PNG?raw=true)


### Результат работы метода POST
![result1](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/post-method.PNG?raw=true)

___________________________________

### Обображение документации на локальном хосте (localhost:9090/docs)

* Результат отображения документации для HTTP метода GET:

![result1](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/get-swagger.PNG?raw=true)


* Результат отображения документации для HTTP метода POST:

![result2](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/post-swagger.PNG?raw=true)


* Результат отображения документации для HTTP метода PUT:

![result3](https://github.com/ellofae/RESTful-API-Gorilla/blob/main/imgs/put-swagger.PNG?raw=true)
