# LazyPm GUI

Interface graphique en mode terminal pour LazyPm, inspirée de lazynpm.

## Architecture

L'interface est organisée en plusieurs composants modulaires :

### Structure des dossiers

```
packages/gui/
├── app.go              # Point d'entrée principal
├── config/             # Configuration
│   ├── config.go       # Configuration générale
│   └── keybindings.go  # Configuration des raccourcis
├── views/              # Vues spécialisées
│   ├── package_manager_view.go  # Affichage du gestionnaire de packages
│   ├── packages_view.go         # Liste des packages
│   ├── dependencies_view.go     # Gestion des dépendances
│   └── scripts_view.go          # Scripts exécutables
├── components/         # Composants réutilisables
│   ├── panel.go        # Interface Panel commune
│   └── layout.go       # Gestionnaire de mise en page
└── handlers/           # Gestionnaires d'événements
    └── keybindings.go  # Gestion des raccourcis clavier
```

## Interface utilisateur

### Layout

L'interface est divisée en 6 panels :

**Côté gauche (4 panels empilés) :**
1. **Package Manager** - Affiche le gestionnaire de packages détecté
2. **Packages** - Liste des packages avec leurs chemins
3. **Dependencies** - Dépendances avec informations de versions
4. **Scripts** - Scripts disponibles (navigables et exécutables)

**Côté droit (2 panels empilés) :**
1. **Output** - Sortie des commandes exécutées
2. **Details** - Détails du package/script sélectionné

### Raccourcis clavier (style Vim)

#### Navigation
- `j/k` - Naviguer vers le bas/haut dans une liste
- `h/l` - Naviguer vers la gauche/droite (non utilisé actuellement)
- `Tab` - Passer au panel suivant
- `H` - Passer au panel précédent

#### Actions
- `Enter` - Exécuter le script sélectionné ou afficher les détails
- `q` - Quitter l'application
- `Ctrl+C` - Forcer la fermeture
- `F5` - Actualiser le panel courant

## Fonctionnalités

### Panels interactifs

1. **Scripts Panel** - Permet d'exécuter des scripts directement
2. **Dependencies Panel** - Affiche les détails des dépendances
3. **Packages Panel** - Montre les informations des packages

### Intégration

- Utilise les helpers existants (`WithManager`, `ListScripts`, `ListDependencies`)
- Compatible avec tous les gestionnaires de packages supportés (npm, go)
- Exécution de scripts en temps réel avec affichage des résultats

## Configuration

### Personnalisation des raccourcis

Les raccourcis peuvent être personnalisés via `config.KeyBindings`. Par défaut, utilise un style Vim.

### Thèmes

Support des thèmes via `config.Theme` (couleurs, styles).

## Extensibilité

### Ajouter un nouveau panel

1. Créer un fichier dans `views/`
2. Implémenter l'interface `components.Panel`
3. Ajouter le panel dans `app.go:createPanels()`

### Personnaliser les raccourcis

1. Modifier `config/keybindings.go`
2. Ajouter les handlers correspondants dans `handlers/keybindings.go`

## Développement

### Conventions

- Chaque vue implémente l'interface `Panel`
- Séparation claire des responsabilités
- Noms cohérents et descriptifs
- Code documenté pour les fonctions publiques

### Tests

Les tests peuvent être ajoutés pour chaque composant individuellement grâce à l'architecture modulaire.