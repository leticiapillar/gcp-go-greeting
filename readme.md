### Desafio Go!

Crie um aplicativo Go que disponibilize um servidor web na porta 8000 que quando acessado seja exibido em HTML (em negrito) Code.education Rocks!

Configuração de uma CI no GCP - Google Cloud Plataform

***Rodando o main***
```
> go run src/greeting/main.go
```

***Build do main***
```
# Gerando o build
> go build src/greeting/main.go
# Executando
> ./main
```

Acesse o link: http://localhost:8000/, aparecerá no browser "Code.education Rocks!"

***Rodando os testes***
```
> go test -v src/greeting/main_test.go src/greeting/main.go
=== RUN   TestMainSuccess
    main_test.go:14: Function sum success
--- PASS: TestMainSuccess (0.00s)
PASS
ok  	command-line-arguments	1.027s
``` 