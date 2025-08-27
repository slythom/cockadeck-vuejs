Principe :
Application web qui contient un formulaire permettant de rechercher une carte Magic via l'API Scryfall. Les cartes recherchées sont enregistrées en mémoire et peuvent être "exportées" dans un fichier *cod que l'utilisateur peut télécharger pour être utilisé dans Cockatrice.

# Architecture (résumée)

- Avoir un frontend avec VueJS 
- Sur le frontend, avoir un formulaire qui envoie un POST avec les données saisies (par exemple, le set code et le numéro de la carte) en JSON 
- Ce POST cible le backend en Go 
- Le backend contient les handlers/fonctions. Le POST récupère les informations du formulaire 
- Le backend construit alors la requête GET vers l'API Scryfall, avec pour paramètre de l'URL les données récupérées par le POST (exemple : https://api.scryfall.com/cards/{setCode}/{collectorNumber} ---> https://api.scryfall.com/cards/otj/123) 
- Le backend éxécute la requête et récupère les informations de la carte en JSON 
- Le backend filtre les informations du JSON (par exemple, ne garde que le nom de la carte et son lien d'image PNG) 
- Le frontend récupère via un GET ces informations et les affichent à l'utilisateur
- Le frontend propose de télécharger les informations de ces cartes au format *cod
