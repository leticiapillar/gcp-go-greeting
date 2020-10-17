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

***Gerando uma imagem Docker***
```
# Gerando uma imagem docker
> docker build -t leticiapillar/go-greeting .

# Executando a imagem docker
> docker run -d -p 8000:8000 leticiapillar/go-greeting
```

Acesse o link: http://localhost:8000/, aparecerá no browser "Code.education Rocks!"

Imagem no Docker Hub: https://hub.docker.com/r/leticiapillar/go-greeting

***Arquivos de configurações para o Kubernetes***

- Deployment: `go-deployment.yaml`
- Service: `go-service.yaml`

```
# Aplicar as configurações de deployment
> kubectl apply -f go-deployment.yml

# Aplicar as configurações de service
> kubectl apply -f go-service.yml

# Executar pelo minikube
> minikube service go-web-service
```

- Referências:
  * [JetBrains: Go development with Docker Containers](https://blog.jetbrains.com/go/2020/05/04/go-development-with-docker-containers/?gclid=CjwKCAjwrKr8BRB_EiwA7eFaprQKMWoOCZWDVaPdVwMXqCzdWuGx3hIj_CgLYJIB7q18nZBwvaNavBoCqV4QAvD_BwE)
  * [Complete Guide to Create Docker Container for Your Golang Application](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e)
  * [Digital Ocean: Tutorial How to Deploy a Resilient Go Application to DigitalOcean Kubernetes](https://www.digitalocean.com/community/tutorials/how-to-deploy-resilient-go-app-digitalocean-kubernetes)
