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

  - Também teremos interação do WhatsApp e do Twilio, o Twilio fornecera o número para ser utilizado no WhatsApp, assim o Twilio receberá a mensagem enviada para o WhatsApp, e chamará o Microsserviço Chat MS, através de uma requisição HTTP. Essa chama HTTP será uma chamada de WebHook.

<p align="center">
  <img alt="Dinâmica do Projeto" src="./github/img01.jpg" width="100%">
</p>
