# Ascii-Art-Web

## Description
ASCII-Art-Web a pour but de réutiliser notre précédent projet "Ascii-art" et réussir à l'incorporer dans une page web afin d'utiliser une interface utilisateur graphique. 

L'utilisateur peut entrer un texte et choisir une police pour ensuite que le serveur communique avec l'hôte de ce dernier pour renvoyer les mêmes caractères en Ascii-art et l'afficher sur la page web. 

Ce site est aussi capable de gérer les erreurs éventuelles liées à : 
- une requête différente de celle demandée (GET ou POST).
- une erreur de lecture des fichiers.
- une erreur interne au programme ASCII-ART.
- une page non trouvée.

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
1. Lancer le fichier ``server.go`` avec la commande ``go run server.go`` (en se trouvant dans le dossier ``ascii/src``)
2. Entrer comme URL : ``localhost:50000``
Résultat : la page web apparaît.

Pour utiliser ce programme sur la page web :

1. Entrer du texte dans le premier bloc de gauche ``Votre texte``
Et ensuite : 
2. Choisir une police d'écriture dans le menu déroulant ``Police d'écriture``
Pour finir : 
3. Cliquer sur le bouton ``Convertir``

## Implementation details : algorithm

### Algorithme global

Dans un premier temps, voici le programme dans sa globalité.

À noter que cet algorithme est seulement représentatif  de l'idée générale (c'est-à-dire qu'il y a moins de variable et de condition). Il ne représente en aucun cas le véritable programme. En effet, il n'y a que les parties importantes au niveau algorithmique.

```
Variable
    Variables du formulaire :
    select(string), text(string) 
    
    Variable principale du serveur:
    finalStr(string)

Début
    select  <-- Choisir(standard, shadow, thinkertoy ?)
    text    <-- Saisir(votre texte)

    Si la requête est POST
    alors:
        finalStr <-- Exécuter le programme ASCII-ART avec pour arguments "select" et "text" et qui retourne une chaîne de caractère en ASCII-ART
    Sinon
       on retourne "BAD REQUEST"
    Fin de Si.

    Afficher finalStr
Fin
```

### Algorithme de ASCII-ART

Voici, ci-dessous, l'algorithme représentant le programme ASCII-ART :

À noter, que c'est aussi une représentation simplifiée du programme (c'est-à-dire qu'il y a moins de variable et de condition). Le choix de la police d'écriture n'est pas non plus représenté dans cet algorithme (car il n'est pas forcément pertinent à expliquer).

Ici, la fonction PrintAscii a été représentée comme ceci, elle était dans la fonction principale (or, elle est normalement dans un fichier séparé).

```
Variable

    Variables importantes :
    finalStr(string)
    asciiLines(tableau de string)

    Variables secondaires:
    str(string)
    everyLines(int)
    characterLine(int)

Début
    asciiLines <-- Exécuter la fonction FileToLines qui lit le fichier contenant les caractères en ASCII-Art et retourne chaque ligne de ce fichier.

    Si asciiLines est vide (car mauvaise lecture du fichier)
    alors :
        retourne une erreur (qui engendre une erreur 500)
    Sinon
        
        everyLine <-- 0
        Tant que everyLines < 9 (correspond à chaque ligne d'une lettre en ASCII-Art)
        faire:
             str <-- valeur nul

             Tant qu'il y a des lettres qui ne sont pas passées (de la chaîne de caractère de l'utilisateur)
             faire:

                    characterLine <-- numéro de la première ligne de la lettre correspondante dans le fichier ASCII-ART
                    str <-- str + asciiLine[everyLines + characterLine]
             Fin Tant que.

             finalStr <-- finalStr + str + retour à la ligne

             everyLine <-- everyLine + 1
        Fin Tant que.

        Afficher finalStr.
    Fin Si.
Fin

```

> :warning: Les algorithmes ci-dessus restent relativement loin du programme final mais permettent de le comprendre plus facilement.

Les différents fichiers sont commentés afin d'améliorer la compréhension (donc pour plus de détails sur la partie serveur, regardez les fichiers).

## Links

 - [Gitea](https://git.ytrack.learn.ynov.com/MLEGER/ascii-art-web.git)

## Try the program

Pour Contribuer au projet ou juste l'esseyer vous pouvez retrouver la dernière version en date dans les releases.