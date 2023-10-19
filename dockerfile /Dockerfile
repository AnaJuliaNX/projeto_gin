#Iniciando a imagem base golang mais recente 
FROM golang:latest
#Add informações do "dono" desse código golang 
LABEL maintainer="Ana <anajulia.treeunfe@gmail.com>"
#Definindo o diretório de trabalho atual dentro do contêiner
WORKDIR /app
#Copiando o modulos 
COPY  go.mod .
COPY go.sum . 
#Fazendo o dowload das dependencias
RUN go mod download 

COPY . . 
#Inidicando qual porta que é a do server
ENV PORT 5000
#Construindo o app
RUN go build
#Removendo todos os arquivos depois de construir o app
RUN find . -name "*.go" -type f -delete
#Fazendo a porta 5000 ser visivel para o mundo fora do container 
EXPOSE $PORT
#Executando o app
CMD ["./projeto-gin"]

#Para ver as imagens que tenho no dockerfile: docker images
#Para adicionar uma nova imagem dockerfile: docker build --tag nomeescolhido/golang-gin
#Para rodar o Dockerfile: docker run -p 5000:5000 nomeescolhido/golang-gin 
#Para fazer login no docker hub: docker login e as credenciais de acesso criados no site 
#Para subir a imagem no docker: docker push nomeescolhido/golang-gin