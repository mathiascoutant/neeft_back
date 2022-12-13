# Neeft back
# Endpoints exposés par l'API

## Utilisateurs

### GET `/api/login`
Login un utilisateur et créé son token JWT qui lui est retourné

**Arguments**
```json
{
    "username": "John Doe",
    "password": "mot de passe"
}
```

**Réponse**
```json
{
    "message": "login success",
    "token": "token",
    "user": { ... }
}
```

### POST `/api/user` `/api/register`
Créé un utilisateur

**Arguments**
```json
{
    "username": "John Doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@salut.fr",
    "password": "mot de passe"
}
```

**Réponse**
```json
{
    "username":  "John Doe",
    "firstName": "John",
    "lastName":  "Doe",
    "email":     "john.doe@salut.fr"
}
```

### GET `/api/users`
Retourne l'ensemble des utilisateurs

**Arguments**
```json
{ }
```

**Réponse**
```json
[
    {
        "username": "John Doe",
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@salut.fr"
    },
    {
        "username": "John Doe",
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@salut.fr"
    }
]
```

### GET `/api/user/{id}`
Retourne les informations de l'utilisateur avec l'id `id`

**Arguments**

`id` (URL) : id de l'utilisateur auquel on veut récupérer les données

```json
{ }
```

**Réponse**
```json
{
    "username": "John Doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@salut.fr"
}
```

### PUT `/api/user/{id}`
Met à jour les informations de l'utilisateur avec l'id `id`

**Arguments**

`id` (URL) : id de l'utilisateur auquel on veut mettre à jour les données

```json
{
    "username": "John Doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@salut.fr"
}
```

**Réponse**
```json
{
    "username": "John Doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@salut.fr"
}
```

### DELETE `/api/user/{id}`
Supprime l'utilisateur avec l'id `id`

**Arguments**

`id` (URL) : id de l'utilisateur à supprimer

```json
{
    "username": "John Doe",
    "firstName": "John",
    "lastName": "Doe",
    "email": "john.doe@salut.fr"
}
```

**Réponse**
```json
"Successfully deleted User"
```

### GET `/api/user/{id}/profile/public`
Retourne le profil public de l'utilisateur avec l'id `id`

**Arguments**

`id` (URL) : ID de l'utilisateur auquel on veut récupérer le profil public

```json
{ }
```

**Réponse**
```json
{
  "id": 11,
  "username": "John Doe",
  "image": "/images/players/profiles/player_placeholder.png",
  "description": "One guy",
  "teams": {...}
}
```

## Amis

### POST `/api/friend`
Ajoute un ami à un utilisateur

**Arguments**
```json
{
    "userId": 42,
    "friendId": 43,
    "isFriend": true
}
```

**Réponse**
```json
{
    "id": 9,
    "user": {...},
    "friend": {...}
}
```

### GET `/api/show-friend/{id}`
Ajoute un ami à un utilisateur ayant l'ID `id`

**Arguments**

`id` (URL) : ID de l'utilisateur à qui on veut ajouter l'ami

```json
{ }
```

**Réponse**
```json
{
    "id": 9,
    "user": {...},
    "friend": {...}
}
```

## Equipes

### POST `/api/team`
Créé une équipe

**Arguments**

`createBy` : Id du leader de l'équipe (le créateur)

```json
{
    "createBy": 3,
    "name": "test",
    "userCount": 4,
    "gameName": "Hello",
    "tournamentCount": 4
}
```

**Réponse**
```json
{
    "user": {...},
    "name": "test",
    "userCount": 4,
    "gameName": "Hello",
    "tournamentCount": 4
}
```

### GET `/api/teams`
Récupère l'ensemble des équipes

**Arguments**
```json
{ }
```

**Réponse**
```json
[
    {
        "id": 54,
        "user": {...},
        "name": "test",
        "userCount": 4,
        "gameName": "Hello",
        "tournamentCount": 4
    }
]
```

### GET `/api/team/{id}`
Récupère les informations d'une équipe avec l'id `id`

**Arguments**

`id` (URL) : L'ID de l'équipe

```json
{ }
```

**Réponse**
```json
{
    "id": 54,
    "user": {...},
    "name": "test",
    "userCount": 4,
    "gameName": "Hello",
    "tournamentCount": 4
}
```

## Tournois

### POST `/api/tournament`
Créé un tournoi

**Arguments**
```json
{ 
    "name": "Nom",
    "count": 2,
    "price": 10000,
    "game": "Lol",
    "teamsCount": 3,
    "isFinished": false,
    "mode": "duo"
}
```

**Réponse**
```json
{
    "name": "Nom",
    "count": 2,
    "price": 10000,
    "game": "Lol",
    "teamsCount": 3,
    "isFinished": false,
    "mode": "duo"
}
```

### GET `/api/tournaments`
Récupère l'ensemble des tournois

**Arguments**
```json
{ }
```

**Réponse**
```json
[
    {
        "name": "Nom",
        "count": 2,
        "price": 10000,
        "game": "Lol",
        "teamsCount": 3,
        "isFinished": false,
        "mode": "duo"
    },
    {
        "name": "Tournoi 2",
        "count": 3,
        "price": 14000,
        "game": "Lol",
        "teamsCount": 6,
        "isFinished": true,
        "mode": "trio"
    }
]
```

### GET `/api/tournament/{id}`
Récupère les informations du tournoi avec l'id `id`

**Arguments**

`id` (URL) : L'ID du tournoi

```json
{ }
```

**Réponse**
```json
{
    "name": "Nom",
    "count": 2,
    "price": 10000,
    "game": "Lol",
    "teamsCount": 3,
    "isFinished": false,
    "mode": "duo"
}
```

### DELETE `/api/tournament/{id}`
Supprime le tournoi ayant l'id `id`

**Arguments**

`id` (URL) : L'ID du tournoi

```json
{ }
```

**Réponse**
```json
"Successfully deleted Tournament"
```