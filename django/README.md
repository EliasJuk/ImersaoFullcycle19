# Como rodar o projeto

``` bash
# Inicia o servidor do Django
$ python manage.py runserver

# Inicia o container
$ docker compose up

http://localhost:8000/  #Django  - admin@admin    | admin
http://localhost:5050/  #PGAdmin - admin@user.com | secret
```

---

## Requisitos
- Python 3.8+
- pip

## Instalação
1. Crie um ambiente virtual:

```bash
  # INSTALAR E CRIAR O AMBINTE VIRTUAL

  # Habilita a criação do ambiente virtual dentro do diretório do projeto
  $ set PIPENV_VENV_IN_PROJECT=1

  # Instala o Pipenv, que é uma ferramenta para gerenciar ambientes virtuais e dependências em projetos Python
  $ pip install pipenv

  # Cria um novo ambiente virtual e especifica a versão do Python a ser utilizada (3.12)
  $ pipenv --python 3.12

  # Ativa o ambiente virtual criado pelo Pipenv
  $ pipenv shell

  # Instala o Django no ambiente virtual
  $ pipenv install django

  # Instala todas as dependências listadas no Pipfile.lock ou no Pipfile
  $ pipenv install

  # Exibe a árvore de dependências instaladas, mostrando quais pacotes estão no ambiente
  $ pipenv graph

  # Mostra o caminho do executável Python que está sendo usado no ambiente virtual
  $ which python

  # Sai do ambiente virtual
  $ exit
```

## Dependências

  - pipenv install Django
  - pipenv install psycopg2-binary
  - pipenv install Pillow

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