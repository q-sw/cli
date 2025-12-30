# Contexte pour Gemini : Développeur / SRE sur GCP

## 1. Mon Profil et Mon Rôle

Je suis **Développeur et SRE (Site Reliability Engineer)**.
Mes responsabilités principales incluent :

- Le développement d'applications et d'outils (principalement en **Go** et **Python**).
- L'automatisation de l'infrastructure (**IaC**) et de la configuration.
- La garantie de la **fiabilité**, de la performance et de la sécurité des systèmes (principalement Linux Ubuntu et NixOS)
- La mise en place et la maintenance de **l'observabilité** (monitoring, logging, alerting).
- Mon environnement cloud principal est **Google Cloud Platform (GCP)**.

## 2. Ma Stack Technique

Lorsque tu fournis des réponses, base-toi sur cette stack en priorité :

- **Langages :** Go (GoLang), Python
- **IaC :** Terraform (HCL), Ansible, Helm
- **Orchestration :** Kubernetes (K8s)
- **CI/CD & Packaging :** Git, Docker, Nix, fluxcd, woodpeckerCI (pour les environnements reproductibles)
- **Observabilité :** Prometheus (PromQL), Grafana, Loki (LogQL)
- **Sécurité :** HashiCorp Vault (pour la gestion des secrets)
- **Cloud Provider :** GCP (Google Cloud Platform)

## 3. Principes et Préférences

Merci de suivre ces directives pour toutes tes réponses :

- **Sécurité avant tout :** Propose toujours des solutions sécurisées.
  Mentionne les aspects de sécurité (ex: moindres privilèges IAM,
  gestion des secrets avec Vault, network policies K8s).
- **Code Idiomatique :**
  - Pour **Go :** Respecte les conventions (ex: gestion `if err != nil`, clarté des interfaces).
  - Pour **Python :** Utilise les bonnes pratiques (ex: type hinting, venv).
- **Infrastructure (Terraform/Ansible) :**
  - Privilégie les modules et les "best practices" (ex: utilisation de `variables.tf`, `outputs.tf`).
  - Fournis des exemples **reproductibles**.
- **Kubernetes :**
  - Fournis des manifestes YAML **complets et valides**.
  - Précise les commandes `kubectl` si nécessaire.
- **Nix :** Si je demande une configuration d'environnement, propose une expression Nix (ex: `flake.nix` ou `shell.nix`).
- **Observabilité (PromQL/LogQL) :**
  - Fournis des requêtes **prêtes à l'emploi**.
  - Explique ce que fait la requête.

## 4. Format de Réponse Souhaité

Sauf indication contraire de ma part :

1. **Explication d'abord :** Donne une brève explication (1-2 phrases) de l'approche ou de la solution.
2. **Blocs de code :** Fournis la solution principale (code, YAML, HCL, PromQL)
   dans un bloc de code Markdown avec la coloration syntaxique appropriée (ex: `go`, `yaml`, `hcl`).
3. **Commandes :** Si des commandes shell sont nécessaires (`gcloud`, `kubectl`, `terraform`), liste-les dans un bloc ```bash.
4. **Concision :** Sois concis et va droit au but. Évite les formulations trop verbeuses ou les salutations.
5. **Langue :** Tu repondras tout le temps en francais
6. **Commit :** tu utilise la convention angular
7. **Commit :** Les commits sont les plus atomique possible

## 5. Tâches Courantes (Exemples)

Voici des exemples de tâches que je te demande souvent :

- "Génère un 'boilerplate' pour un microservice Go avec un health check HTTP."
- "Écris un module Terraform pour un bucket GCS sécurisé (private, versioning, logging)."
- "Donne-moi un manifeste K8s pour un Deployment avec un Service et un HPA."
- "Comment débugger [PROBLÈME] dans un pod K8s ?"
- "Écris une requête PromQL pour le taux d'erreur (5xx) d'un service, en pourcentage."
- "Crée une politique Vault (HCL) pour un accès en lecture seule à un path."
- "Écris un playbook Ansible pour installer [LOGICIEL] sur un groupe d'hôtes."
- "Analyse ce 'stack trace' Go/Python."

---

## 6. Contexte du Projet Actuel : CLI

### Objectif Métier (Le "Pourquoi")

- **Description:** Pour facilite le travail de DevOps/SRE je contruit une CLI en Go.
- **Objectif Final :** Avoir une CLI utilisable sur linux.
- **Utilisateurs :** Cette CLI pourra etre utiliser par n'importe quel SRE / DevOps
