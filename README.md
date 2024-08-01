# Projeto de Cadastro Simples em Go

Este projeto foi criado com o intuito de servir como um modelo para futuros projetos de cadastro e gerenciamento de usuários, utilizando a linguagem de programação Go. Nele foram aplicados conhecimentos em diversas áreas, incluindo banco de dados (MySQL), estrutura MVC, API REST, e manipulação de dados em JSON.

## Funcionalidades

O projeto possui as seguintes funcionalidades através de endpoints:

- `GET /lista` - Lista todos os cadastros presentes no banco de dados.
- `POST /cadastro` - Realiza novos cadastros.
- `PUT /atualizar` - Atualiza dados já cadastrados.
- `DELETE /limpar/{id}` - Exclui um cadastro através do ID do usuário.
- `GET /user/{id}` - Apresenta o usuário através do ID.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para o desenvolvimento do projeto.
- **MySQL**: Banco de dados utilizado para armazenar as informações dos cadastros.
- **MVC**: Padrão de arquitetura utilizado para organizar o código em Model, View e Controller.
- **API REST**: Interface de programação de aplicações que permite a comunicação entre cliente e servidor utilizando operações HTTP.
- **JSON**: Formato de intercâmbio de dados utilizado para transmitir as informações entre cliente e servidor.

## Endpoints

1. **Listar todos os cadastros**

   ```http
   GET /lista
   ```

   Retorna uma lista de todos os cadastros presentes no banco de dados.

2. **Realizar novo cadastro**

   ```http
   POST /cadastro
   ```

   Recebe um JSON com as informações do novo cadastro e o insere no banco de dados.

3. **Atualizar cadastro existente**

   ```http
   PUT /atualizar
   ```

   Recebe um JSON com as informações atualizadas de um cadastro existente e atualiza os dados no banco de dados.

4. **Excluir cadastro pelo ID**

   ```http
   DELETE /limpar/{id}
   ```

   Exclui o cadastro correspondente ao ID fornecido.

5. **Obter usuário pelo ID**
   ```http
   GET /user/{id}
   ```
   Retorna as informações do usuário correspondente ao ID fornecido.

## Autor

Este projeto foi desenvolvido por:
**Victor A. Maximiano**

---

Esperamos que este projeto seja útil como base para futuros desenvolvimentos. Sinta-se à vontade para adaptar e expandir conforme necessário!

## Como Executar

Para executar o projeto, siga os passos abaixo:

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   ```
2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```
3. **Configure o banco de dados:**
   Atualize as configurações de conexão com o banco de dados MySQL no arquivo de configuração.

4. **Execute a aplicação:**
   ```bash
   go run main.go
   ```

Agora, a aplicação estará disponível nos endpoints especificados.

---

Este README.md foi elaborado para fornecer uma visão clara e detalhada do projeto, facilitando a compreensão e o uso do código. Se houver dúvidas ou sugestões, por favor, entre em contato.
