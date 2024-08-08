# **Mangi**
Liste de courses et + si affinités...

## **Intro**

Dans la société actuelle, à l'heure du tout numérique, l'organisation quotidienne et la gestion du temps représentent souvent des défis majeurs pour de nombreuses personnes. Faire ses courses et planifier les repas pour la semaine nécessite souvent une organisation bien réfléchie entre la liste des courses et les recettes à préparer. Ainsi, il devient essentiel de disposer d'un outil pratique et efficace pour simplifier cette tâche Comment concevoir un outil numérique efficace pour aider les individus à organiser leurs courses et planifier leurs repas de manière optimale dans un contexte où la gestion du temps et l'organisation quotidienne posent des défis importants ?

Pour répondre à la problématique nous proposons les fonctions clés suivantes:   
Planification des repas pour la semaine   
Création de la liste de course selon repas plannifiés   
Proposition de recettes   
Filtrage des recettes en fonctions de critères de préférences (sans gluten, vegan.)   
Application collaborative (possibilité d'ajouter les membres de la famille, amis, colocataires, autres utlisateurs de l’app...)

## **Contents**

- [Prerequisites]()
- [Installation guide]()
- [User manual]()
- [Crew]()

## **Prerequisites**

- Golang
- mysql
- react / TS
-

## **Installation guide**

From mangi's folder in your terminal    
run cli to build binary files   
`make api`   
run bin/api to run the binary files   
`bin/api`    
The instance of localhost:8080 is your api's binary runing    
Open a new terminal's window to check the endpoints as in the user manual.

## **User manual**

### User collection:
user register   
`POST /user/register HTTP/1.1
Host: localhost:8080
Content-Length: 67

{
  "name": "caroll",
  "password": "admin1",
  "email": "admin1"
}`

user login
`POST /user/login HTTP/1.1
Host: localhost:8080
Content-Length: 47

{
  "email": "admin1",
  "password": "admin1"
}`

### recipe collection:
create
`
POST /recipe/create HTTP/1.1
Host: localhost:8080
Content-Length: 32

{
  "name": "tarte aux pommes"
}`

### meal collection:
create
`POST /meal/create HTTP/1.1
Host: localhost:8080
Content-Length: 77

{
  "planned_at": "2024-08-01T19:00:00.000Z",
  "guests": 2,
  "user_id": 3
}`

fetch all meals
`POST /meals HTTP/1.1
Host: localhost:8080
Content-Length: 90

{
  "user_id": 2,
  "from": "2020-08-01T19:00:00.00Z",
  "to": "2025-08-01T19:00:00.00Z"
}`

## **Crew**

- Valentine B. : architecture / base de données / engineer devops
- Caroll K. : architecture / engineer backend / base de données
- Shaïnez B. : front End / design
- Vincent P. : engineer fullstack
- Hilda B. : engineer fullstack
