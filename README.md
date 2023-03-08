# Recipe API

Basic API to retrieve cooking recipes and the chefs who created them.

## Payloads

Payload exampleas for necessary API endpoints

### Chef

```json
{
 "name": "John Stewart",
 "country": "USA",
 "yearsOfExperience": 5
}
```

### Recipe

```json
{
 "name": "Spaghetti",
 "keywords": ["pasta", "italian", ],
 "instructions": ["Boil spaghetti", "prepare stew/sauce", "eat"],
 "ingredients": ["water", "tomatoes", "pepper", "salt", "golden penny spaghetti"],
 "chefId": "any_chef_id"
}
```

## API Endpoints
