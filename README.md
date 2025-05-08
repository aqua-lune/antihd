# AntiHD Outil chiant™ de suivi des tâches

## Fonctionnalités
- Liste les tâches
- Chiant™ asf
- Ne se ferme pas tant que toutes les tâches ne sont pas complétées
- Ne permet pas de supprimer de tâches (oui oui, on se sait avec les petits "arrangements entre nous")
- Indique le temps ~~de paralysie exécutive~~ depuis le lancement du programme
- S'adapte au DPI et tout le bazar
- Très relativement léger, mais good enough je suppose (3 mo)
- Nom un peu rigolo
- Icône un peu rigolote :catnobrain: 
- C'est tout

![screen](https://github.com/user-attachments/assets/44a12aac-85fa-4c5c-a886-835b7d57e3d1)

Le programme ne fonctionne que sous Windows 64 bits. 
Je *suppose* qu'il n'y a pas grand-chose qui l'empêcherait de tourner ailleurs, mais j'ai sincèrement la flemme.

Livré sans aucune forme de service après vente. 
Techniquement limité à 128 tâches, mais si vous arrivez à ce chiffre, c'est que vous avez un problème.

## Build

Nécessite CGO (donc un toolchain C installé)
```bash
go build -ldflags="-H=windowsgui"
```
