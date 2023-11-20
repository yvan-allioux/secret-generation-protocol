# secret-generation-protocol
 
Ce projet implémente un protocole de partage de secrets utilisant des WebSockets pour la communication en temps réel entre le client et le serveur.


## Usage Général

Démarrage: Les clients se connectent au serveur via WebSocket.
Échanges de Données: Les clients envoient des données au serveur à des intervalles aléatoires.
Génération de Secret: Une fois toutes les données collectées, le serveur et les clients utilisent ces données pour générer un secret partagé.

## Enjeux de Timing

Synchronisation en Temps Réel:

La gestion précise des temps d'envoi et de réception des données via les WebSockets est cruciale. Tout décalage ou latence peut affecter la fiabilité du protocole.
Délais de Réseau:

Contrairement à un script de démonstration, les délais inhérents à la communication réseau dans une application réelle peuvent introduire des incertitudes dans les timestamps et la séquence des données.
Traitement Asynchrone:

L'envoi et la réception asynchrones des données exigent une gestion rigoureuse des timings pour maintenir la cohérence des données entre le client et le serveur.
Gestion de la Concurrence
La simultanéité des requêtes de multiples clients nécessite une synchronisation précise pour éviter les conflits et garantir que les données sont traitées dans le bon ordre.
Conclusion
Dans une mise en œuvre réelle, la gestion précise des timings et de la synchronisation est bien plus critique que dans un script de démonstration. Ces aspects doivent être méticuleusement gérés pour assurer l'intégrité et la fiabilité du protocole dans un environnement de production.



## Code Client (HTML & JavaScript)

Structure HTML
Interface Utilisateur (UI): Utilise Bootstrap pour le style. Contient des éléments pour afficher le secret généré, le statut des clients, un tableau pour obtenir les résultats, et des boutons pour diverses actions (démarrer, obtenir le tableau, réinitialiser, etc.).
Formulaires et Entrées: Inclut des boutons pour démarrer le processus, obtenir et réinitialiser les données du tableau, ainsi qu'une entrée pour la tolérance en millisecondes.
Fonctionnalités JavaScript
Connexion WebSocket: Établit une connexion WebSocket avec le serveur.
Gestion des Événements WebSocket:
onopen: Active les boutons et envoie une requête pour réinitialiser le tableau côté serveur.
onmessage: Gère les différents types de messages reçus (confirmation, tableau de données, etc.).
onerror et onclose: Gère les erreurs et la fermeture de la connexion WebSocket.
Envoi de Données: Des fonctions pour envoyer des messages au serveur (démarrage du client, demande de tableau, réinitialisation du tableau).
Génération et Comparaison des Données: Envoie des requêtes de manière séquentielle avec des délais aléatoires et compare les résultats reçus du serveur avec les données locales.
Logique du Secret Partagé
Le client participe à un protocole où il envoie des requêtes avec des délais aléatoires et compare les réponses du serveur avec ses propres enregistrements pour générer un "secret".

## Code Serveur (Go)
Structure Principale
Struct SecretSharing: Gère l'état global, y compris les secrets partagés, les clients connectés, et leur état (prêts, terminés).
WebSocket Upgrader: Permet la mise à niveau des requêtes HTTP en connexions WebSocket.
Fonctionnalités Principales
Gestion des Connexions WebSocket: Accepte et gère les connexions WebSocket entrantes.
Traitement des Messages Clients: Réagit aux différents types de messages envoyés par le client (nouvelle valeur, demande de tableau, réinitialisation du tableau, état du client).
Logique de Synchronisation: Assure que tous les clients sont prêts et ont terminé d'envoyer des requêtes avant de procéder à l'étape suivante.
Fonctions Clés
handleConnections: Gère les nouvelles connexions WebSocket.
handleClientMessage: Traite les messages individuels des clients.
sendConfirmationToAll: Envoie une confirmation à tous les clients une fois que tous ont terminé.
handleMessages: (vide dans votre code) Prévu pour gérer les messages globaux ou les événements.

