# gotodo

API REST de gestion de todos, écrite en Go avec la bibliothèque standard (`net/http`).

## Objectif

Un petit projet pour manipuler une API Go simple : CRUD sur des todos, stockage en mémoire (pas de base de données), avec une documentation OpenAPI.

## Stack

- Go, `net/http` standard (routeur `ServeMux` de Go 1.22+)
- Stockage in-memory (les données sont perdues au redémarrage du serveur)
- Documentation API : OpenAPI 3.0 (fichier `openapi.yaml`)

## Modèle de données

Un `Todo` :

| Champ       | Type     | Description                          |
|-------------|----------|--------------------------------------|
| `id`        | string   | Identifiant unique (généré au create)|
| `title`     | string   | Titre du todo (obligatoire)          |
| `description` | string | Description libre (optionnel)        |
| `done`      | bool     | Statut fait / pas fait               |
| `created_at`| datetime | Date de création                     |
| `updated_at`| datetime | Date de dernière modification        |

## Endpoints

| Méthode | Route            | Description                     |
|---------|------------------|----------------------------------|
| GET     | `/todos`         | Liste tous les todos             |
| GET     | `/todos/{id}`    | Récupère un todo par son id      |
| POST    | `/todos`         | Crée un todo                     |
| PUT     | `/todos/{id}`    | Met à jour un todo (complet)     |
| DELETE  | `/todos/{id}`    | Supprime un todo                 |

Réponses en JSON. Codes HTTP standards (`200`, `201`, `404`, `400`, etc).

## Lancer le projet

```bash
go run ./cmd/gotodo
```

Le serveur écoute par défaut sur `http://localhost:8080`.

## Documentation API

La spec OpenAPI est disponible dans [`openapi.yaml`](openapi.yaml).

## Front

Pas de front pour l'instant. On teste l'API avec curl/Postman. Un front (probablement HTML/JS simple servi par le serveur Go, pour éviter le CORS) pourra être ajouté plus tard.

## Structure du projet (prévue)

```
gotodo/
├── cmd/gotodo/         # point d'entrée (main.go)
├── internal/todo/      # modèle, store in-memory, handlers HTTP
├── openapi.yaml         # spec OpenAPI
└── README.md
```

## Statut

🚧 En cours de construction — ce README définit le scope avant l'implémentation.
