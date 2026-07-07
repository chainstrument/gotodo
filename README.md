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

Le projet tourne dans Docker (pas besoin d'installer Go en local) :

```bash
docker compose up --build
```

Le serveur écoute sur `http://localhost:8080`.

## Documentation API

La spec OpenAPI est disponible dans [`openapi.yaml`](openapi.yaml).

## Front

Un front minimal en HTML/CSS/JS vanilla (pas de build, pas de framework), servi par le serveur Go lui-même via `go:embed` sur `/`. Même origine que l'API, donc pas de CORS à gérer. Permet d'ajouter, cocher (done) et supprimer des todos depuis le navigateur : `http://localhost:8080`.

## Structure du projet

```
gotodo/
├── cmd/gotodo/          # point d'entrée (main.go)
├── internal/todo/       # modèle, store in-memory, handlers HTTP
├── internal/web/        # front HTML/JS/CSS embarqué (go:embed) et servi sur /
├── openapi.yaml         # spec OpenAPI
├── Dockerfile           # build multi-stage de l'API
├── compose.yaml         # docker compose pour lancer l'API
└── README.md
```

## Statut

✅ CRUD de todos implémenté et testé (in-memory, via Docker).
✅ Front HTML/JS simple servi par le serveur Go, testé dans le navigateur.
