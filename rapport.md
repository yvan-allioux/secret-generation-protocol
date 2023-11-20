# Rapport de TP : Canal de Diffusion Anonyme et Génération de Secret

**Date de réalisation :** 30 novembre 2020  
**Durée du TP :** 4 heures  
**Date de soumission :** 18 décembre 2020

## Introduction

Ce TP aborde la conception et l'implémentation d'un canal de diffusion anonyme et d'un mécanisme de génération de secret. L'objectif est de faciliter l'échange anonyme de messages et de générer un secret partagé entre deux parties, même en présence d'un adversaire passif.

## Contexte Théorique

Le canal de diffusion anonyme est crucial en cryptographie et sécurité des réseaux, permettant l'envoi anonyme de messages à un groupe. La génération de secret est essentielle pour les communications sécurisées, surtout lorsque le canal de communication est surveillé.

## Partie 1 : Canal de Diffusion Anonyme

### Objectifs et Méthodologie
Création d'une classe pour simuler un canal de diffusion anonyme, capable de poster et récupérer des messages anonymes, avec des étiquettes temporelles pour le tri.

### Résultats
Le système permet de poster et récupérer des messages de manière anonyme, avec un tri chronologique efficace.

## Partie 2 : Génération de Secret via Canal de Diffusion Anonyme

### Objectifs et Méthodologie
Développement d'un protocole pour générer un secret partagé entre deux entités, Alice et Bob, utilisant le canal anonyme.

### Implémentation et Résultats
Le protocole a permis la génération d'un secret robuste, inconnu d'un adversaire passif, grâce à l'échange anonyme de messages.

## Partie 3 : Analyse du Protocole

### Applications Potentielles
Utilisations dans des systèmes de vote électronique sécurisé et des plateformes pour lanceurs d'alerte. Utile pour les communications sécurisées entre utilisateurs.

### Sécurité du Protocole
L'anonymat et la génération aléatoire du secret protègent contre la déduction du secret par un adversaire passif.

### Importance des Primitives Cryptographiques
Ces primitives sont cruciales pour la confidentialité et l'intégrité des données dans les environnements surveillés, assurant des communications sécurisées.

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

## Conclusion
Le TP a mis en pratique des concepts importants en cryptographie et sécurité des réseaux. Le canal de diffusion anonyme et le protocole de génération de secret démontrent une application pratique efficace dans divers contextes, protégeant l'anonymat et la sécurité des communications.
