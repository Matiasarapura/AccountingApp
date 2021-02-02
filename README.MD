# Accounting Notebook

### API-REST
##### /transaction POST
request body
```json
{
    "amount":500,
    "type":"debit/credit" 
} 
```
##### /transactions GET

##### /transaction/{id} GET
    
##### /balance GET
### Tech
* Go 1.14
* Gorilla Mux 1.8
* Bootstrap 4.4
* Jquery

```
### Run

```sh
$ go build -a -installsuffix cgo -o accounting .
$ ./accounting -b {starter balance} -p {port}
```

For production environments...

```sh
$ docker build . -t {name}
$ docker run {name}
```

¡


