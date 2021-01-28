# Ascii-Art-Web
## Description
Ce projet a pour but de réutiliser notre ancien projet Ascii-art et réussir à l'incorporer dans une page web. 

L'utilisateur peut entrer un texte et choisir une police pour ensuite que le serveur communique avec l'hôte de ce dernier pour renvoyer les mêmes caractères en Ascii-art et l'afficher sur la page web. 

Ce site est aussi capable de gérer les erreurs éventuelles liées à : 
- un port différent de celui utilisé pour lancer le serveur,
- une erreur de lecture des fichiers.

## Authors 
- **Mathéo LEGER** :  
                        - la création du serveur  
                        - la liaison entre programme Ascii-Art et l'affichage sur le site web.  
                        - CSS  
                        - HTML  

- **Matteo GRELLIER** :  
                        - CSS  
                        - HTML  
                        - gestion de git pour le transfert de fichier.  

## Usage : How to run
Pour lancer le serveur et utiliser la page web il faut :
1. Lancer le fichier ``server.go`` avec la commande ``go run server.go``
2. Entrer comme url ``localhost:50000``
Résultat : la page web apparaît.

Pour utiliser ce programme sur la page web :

1. Entrer du texte dans le premier bloc de gauche ``Votre texte``
Et ensuite : 
2. Choisir une police d'écriture dans le menu déroulant ``Police d'écriture``
Pour finir : 
3. Cliquer sur le bouton ``Convertir``

## Implementation details : algorithm

