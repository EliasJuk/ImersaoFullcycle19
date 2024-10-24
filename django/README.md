## Como rodar o projeto

## Requisitos
- Python 3.8+
- pip

## Instalação
1. Crie um ambiente virtual:

    ```bash
    # INSTALAR E CRIAR O AMBINTE VIRTUAL
    $ pip install pipenv
    $ python -m venv venv
    ```

    ```bash
    # BASH
    $ .\venv\Scripts\activate.bat     # ENTRAR DO AMBIENTE VIRTUAL DO PYTHON
    $ .\venv\Scripts\deactivate.bat   # SAIR DO AMBIENTE VIRTUAL DO PYTHON

    # POWERSHELL
    $ Set-ExecutionPolicy RemoteSigned -Scope CurrentUser #Permitir execução scripts locais
    $ .\venv\Scripts\Activate.ps1
   ```

## Dependências

  ```Python
  # Instale as dependências:
  $ pip install -r requirements.txt
  ```

  ```Python
  # Gerar o arquivo de dependências:
  $ pip freeze > requirements.txt
  ```

---

## INICIANDO NOVO PROJETO

```Python
# Criar novo projeto com django:
$ python -m django startproject videos

# Aplique as migrações:
$ python manage.py migrate

# Iniciar servidor:
$ python manage.py runserver

# ACESSO PELO NAVEGADOR
http://localhost:8000/admin
```
# DJANGO

### VER COMANDOS

``` bash
# Mostra a lista de comandos
$ python .\manage.py

# Cria todas as tabelas:
$ python manage.py migrate

# Criar SuperUser para a database
$ python manage.py createsuperuser
  #USUARIO E SENHA
  - USER: admin
  - MAIL: admin@admin.com
  - PASS: admin
```


``` bash
# Criar container
$ docker compose up

# PGAdmin
$ http://localhost:5050/

# Criar modulo core
$ django-admin startapp core

# Cria as migrações
$ python manage.py makemigrations
  - core\migrations\0001_initial.py
    + Create model Tag
    + Create model Video
    + Create model VideoMedia

# Executar migração
$ python manage.py migrate
```