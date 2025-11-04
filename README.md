# MediTrack - your way to the ESST.

**MediTrack** est une application web qui permet la gestion de cabinets médicaux. Elle concerne cinq acteurs : 
le SuperAdmin, l’Admin, les médecins, l’assistant ainsi que le patient. L’objectif est de permettre une bonne organisation des activités du cabinet, en définissant les fonctionnalités de chacun des acteurs, et ce en assurant la cohérence et la fiabilité des informations disponibles sur l’application. Chaque entité aura accès à un dashboard avec des options spécifiques aux tâches qui lui sont attribuées.

# Rôles & Fonctionnalités : 

## 1/ Le SuperAdmin
Il est le concepteur/développeur de l’application et est en charge de l’aspect organisationnel et technique.
Ses fonctionnalités sont :
- S’authentifier avec identifiant et mot de passe;
- Créer/modifier/supprimer les cabinets;
- Créer/modifier/supprimer les comptes des médecins Admin; 
- Associer chaque cabinet à son médecin Admin;
- Affecter un cabinet à un pack premium particulier selon le choix de l’Admin;
- Assigner un médecin à la place de l’Admin actuel dans le cas où celui en fonction ne l’est plus;
- Envoyer des notifications globales.

## 2/ Le Médecin Admin
Il est le médecin qui gère l’organisation globale du cabinet et est lui même praticien au sein de ce dernier.
Ses fonctionnalités sont :
- S’authentifier avec identifiant et mot de passe;
- Créer/modifier/archiver les comptes des médecins du cabinet;
- Créer/modifier/archiver les comptes assistants;
- Modifier les informations générales du cabinet (nom, adresse, téléphone, liste des médecins…);
- Gestion des horaires d’ouverture, durée approximative des consultations, des congés, des fermetures…;
- Consulter les informations administratives des patients;
- Définir la politique de réservation;
- Consulter le planning global du cabinet (de tous les médecins);
- Sélectionner un pack premium pour le cabinet.

## 3/ Le Médecin
Il assure le suivi médical des patients.
Ses fonctionnalités sont :
- S’authentifier avec identifiant et mot de passe;
- Créer/modifier (observations/prescriptions complémentaires) une consultation pour un patient;
- Consulter l’historique des consultations effectuées;
- Consulter son agenda, mettre à jour ses disponibilités…;
- Recevoir les notifications des rendez-vous modifiés/annulés du patient par l’assistant;
- Rechercher le patient dans la base du cabinet;
- Créer/consulter la fiche patient (informations administratives et dossier médical);
- Accéder à l’historique complet du patient (consultations, prescriptions, examens…);
- Mettre à jour les informations médicales (antécédents, traitements…);
- Orienter le patient à un confrère au sein du cabinet;
- Modifier son profil (spécialité, contact…);
- Gérer son mot de passe;
- Définir ses honoraires.

## 4/ L’Assistant
Il assure la gestion administrative et logistique du cabinet afin de faciliter le travail des médecins et la gestion des patients.

Ses fonctionnalités sont :
- S’authentifier avec identifiant et mot de passe;
- Planifier un rendez-vous pour un patient avec un médecin du cabinet (choix du médecin, date, heure…);
- Modifier/annuler un rendez-vous existant;
- Enregistrer la présence/absence du patient le jour du rendez-vous;
- Afficher l’agenda global/programme du jour du cabinet ou celui d’un médecin spécifique;
- Ajouter un patient/mettre à jour en insérant ses informations administratives;                                                                                                - Encaisser le patient en attribuant la somme au médecin ayant effectué la consultation.

## 5/ Le Patient
Il est l’utilisateur qui fait appel aux services proposés par le cabinet.
Ses fonctionnalités sont :
- Créer/modifier le compte patient;
- S’authentifier avec identifiant et mot de passe;
- Modifier les informations personnelles;
- Prendre/modifier/annuler un rendez-vous avec un médecin d’un cabinet spécifique;
- Consulter les consultations passées et à venir;
- Consulter les antécédents médicaux, bilans, comptes rendus, prescriptions…;
- Accéder aux factures;
- Évaluer le médecin (note, commentaire…).
