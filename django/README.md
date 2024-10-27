# Como rodar o projeto

``` bash
# Inicia o container
$ docker compose up

# Instala o Pipenv, que é uma ferramenta para gerenciar ambientes virtuais e dependências em projetos Python
$ pip install pipenv

# Instala todas as dependências listadas no Pipfile.lock ou no Pipfile
$ pipenv install

#Ativa o ambiente virtual
$ pipenv shell

# Inicia o servidor do Django
$ python manage.py runserver

http://localhost:8000/  #Django  - admin@admin    | admin
http://localhost:5050/  #PGAdmin - admin@user.com | secret
```

---

## Requisitos
- Python 3.8+ (OBS:. Python 3.12)
- pip

## Instalação
1. Crie um ambiente virtual:

```bash
  # INSTALAR E CRIAR O AMBINTE VIRTUAL

  # Habilita a criação do ambiente virtual dentro do diretório do projeto
  $ set PIPENV_VENV_IN_PROJECT=1

  # Instala o Pipenv, que é uma ferramenta para gerenciar ambientes virtuais e dependências em projetos Python
  #$ pip install pipenv
  # pipenv --python 3.13 --venv C:\Caminho\Para\O\Diretorio

  # Cria um novo ambiente virtual e especifica a versão do Python a ser utilizada (3.12)
  $ pipenv --python 3.12

  # Atualize o Pipenv
  #pip install --upgrade pipenv

  # Ativa o ambiente virtual criado pelo Pipenv
  $ pipenv shell

  # Instala o Django no ambiente virtual
  #$ pipenv install django

  # Instala todas as dependências listadas no Pipfile.lock ou no Pipfile
  #$ pipenv install

  # Exibe a árvore de dependências instaladas, mostrando quais pacotes estão no ambiente
  #$ pipenv graph

  # Mostra o caminho do executável Python que está sendo usado no ambiente virtual
  #$ which python

  # Sai do ambiente virtual
  #$ exit
```

## Dependências

  - pipenv install Django
  - pipenv install psycopg2-binary
  - pipenv install Pillow

---
# DJANGO

### VER COMANDOS

``` bash
# Mostra a lista de comandos
$ python .\manage.py

# Inicia o container
$ docker compose up

# PGAdmin
$ http://localhost:5050/

# Cria todas as tabelas:
$ python manage.py migrate

# Criar SuperUser para a database
$ python manage.py createsuperuser
  #USUARIO E SENHA
  - USER: admin
  - MAIL: admin@admin.com
  - PASS: admin
```

---

## CORE

``` bash
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


pipenv shell
python manage.py shell 
from core.models import Video
Video.objects.get(id=1)
Video.objects.get(id=1).video_media
Video.objects.get(id=1).video_media.status