# Full Cycle Learning Experience

## Aula 01 Docker, Containers and Microservice of ChatGPT

- Entender o projeto prático
- Tecnologias que serão utilizadas
- Docker
- Início do Microsserviço

### Projeto Prático

- Desenvolveremos duas formas/interfaces para utilizarmos o CHatGPT:
  - Interface Web
  - WhatsApp

#### Dinâmica do Projeto

- Utilizaremos API do da OpenAI, para conseguirmos gerar a conversação com ChatGPT
- Criaremos um Microsserviço de chat (MS Chat)
- Construiremos um BackEnd, em NextJS, para se comunicar com o Chat MS. E também criaremos um FrontEnd para se comunicar com o BackEnd.

  - Assim o usuário utilizará o FrontEnd para chamar o BackEnd para chamar o ChatMS, através do gRPC, que irá consumir a API da OpenAI.
