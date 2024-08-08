# **How to**

## Intro

Attention, je suis à deux doigts de nommer ce fichier les 10 commandements.    
Ceci doit être votre référence de travail afin que nous uniformisions nos processus et de permettre d'avoir plus de facilité de travail collectif ainsi que de compréhension.    
Tout est discutble et modifiable. Par contre, tant que cela n'est pas validé, hors de question que  cela soit la foire au "allez je teste on verra bien" sous peine de **Kamé Hamé Ha** loool

## **Table des matières**

- [Git]()
    - [Création de Branche](https://rendu-git.etna-alternance.net/module-9681/activity-52182/group-1038068/-/blob/main/HOWTO.md#cr%C3%A9ation-de-branche)
    - [Commit sur branche de travail](https://rendu-git.etna-alternance.net/module-9681/activity-52182/group-1038068/-/blob/main/HOWTO.md#commit-sur-branche-de-travail)
    - [Merge vers main](https://rendu-git.etna-alternance.net/module-9681/activity-52182/group-1038068/-/blob/main/HOWTO.md#merge-vers-main)

## **Git**

Je propose le process suivant :   
Nous travaillons à partir de la branche main (ou dev : nous verrons avec Valentine selon le type de déploiement continu), nous créons des branches selon les tickets de taches que nous devrons faire. Puis une fois la tache faite, nous mergeons vers la branche de départ.

Je conseille très fortement par conséquent que les tickets découpes les taches dans le plus petit dénominateur possible. Cela permettra non seulement de sentir l'avancement du projet mais aussi de pouvoir avancer avec des petits morceaux de code:   
- plus facile de rollback en cas de besoin.
- plus facile de controler le code.
- moins lourd pour celui qui doit s'occuper de la tache.


### **Création de Branche**

Dans une soucis d'uniformité de travail, voici les process.  
Nommage de branche:   
```Nom-de-Ticket_description-efficace-et-rapide-du changement```    
Depuis la branche main afin d'être à jour :   
```git pull```
```git checkout -b <nom_branche>```

Exemple de nom de commande:
```git checkout -b MG-2501_add-database-connexion```   

### **Commit sur branche de travail**
Depot standart
Afin de garder de la lisibilité, nous n'allons faire qu'un commit par branche.   
Pas de panique, je vous explique.   
vous avez du travail à push ...   
```git add <fichier>```    
```git commit -m "Nom-de-Ticket_description-efficace-et-rapide-du changement"```   
d'où l'importance de bien choisir le nom de ticket/branche.   
```git push```    

#### Depot complémentaire
Vous avez déjà un commit mais vous avez des changements à ajouter.    
```git add <fichier>```     
```git commit --amend```    
la console qui va s'ouvrir va vous permettre de changer le nom de commit si besoin.   
sinon vous sortez de la console avec ```:wq```   
```git push``` ou ```git push --force-with-lease```   

### **Rebase avant de merge vers main**

Selon nous verrons si nous faison des MR ou non.   
Sinon, nous arrivons au moment de merge vers le main. Il va falloir en amont vous mettre à jour avec main.   
```git checkout main```   
```git pull```   
Puis mettre à jour votre branche.    
```git checkout <branche>```    
```git rebase main```   
Résoudre les conflits si besoin.   
```git push```  

Il se peut que vous ayez besoin de force le push à cause de votre MR si tel est le cas merce de saisir.    
ATTENTION : normalement cela de devrait être necessaire uniquement en cas de MR.
```git push --force-with-lease```   

### **Merge vers main**

Nous en arrivons à la partie qui devra être éclaircie.   
#### SOIT Gitbal
nous faisons des MR via gitlab et donc il suffit de créer / valider la MR sur l'interface gitlab. Normalement Gitlab fermera automatiquement la branche.
#### SOIT Via le terminal
Nous ne mettons pas en place les MR et donc à ce moment là :    
Retour sur la branche main   
```git checkout main```   
Merger la branche vers main en local    
```git merge <branche>```   
Pousser sur le distant votre merge local    
```git push```   
Supprimer votre branche locale si vous le souhaitez (perso je m'en fou) en distant abandonner l'idée ... très relou.