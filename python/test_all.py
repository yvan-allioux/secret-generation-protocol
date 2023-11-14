import random
import time

class CanalDiffusionAnonyme:
    def __init__(self):
        self.messages = []

    def posterMessageAnonyme(self, message):
        timestamp = time.time()
        nouveau_message = {"message": message, "timestamp": timestamp}
        self.messages.append(nouveau_message)
        self.messages = sorted(self.messages, key=lambda x: x["timestamp"])

    def recupererMessagesAnonymes(self, debut, fin):
        messages_anonymes = []
        debut = float(debut)
        fin = float(fin)
        for message in self.messages:
            if debut <= message["timestamp"] <= fin:
                messages_anonymes.append(message["message"])
        return messages_anonymes

    def afficherMessages(self):
        for message in self.messages:
            print(f"{message['timestamp']} : {message['message']}")

class GenerationSecretThroughAnonymousCanal:
    def __init__(self):
        self.canal_a = CanalDiffusionAnonyme()
        self.canal_b = CanalDiffusionAnonyme()

    def genererSecret(self, nom_a, nom_b, duree):
        secret = []

        debut_periode = time.time()
        fin_periode = debut_periode + duree

        while time.time() < fin_periode:
            b_a = random.randint(0, 1)
            b_b = random.randint(0, 1)

            message_a = f"{nom_a}" if b_a == 0 else f"{nom_b}"
            message_b = f"{nom_b}" if b_b == 0 else f"{nom_a}"

            self.canal_a.posterMessageAnonyme(message_a)
            self.canal_b.posterMessageAnonyme(message_b)

            time.sleep(random.uniform(0.001, 0.01))  # Attente aléatoire entre 1 et 10 ms

        messages_a = self.canal_a.recupererMessagesAnonymes(debut_periode, fin_periode)
        messages_b = self.canal_b.recupererMessagesAnonymes(debut_periode, fin_periode)

        for message in messages_a:
            if nom_a in message:
                secret.append(0)
            else:
                secret.append(1)

        return secret

    def extraireSecret(self, transcript_a, transcript_b):
        secret_commun = []

        for message_a, message_b in zip(transcript_a, transcript_b):
            if message_a == message_b:
                secret_commun.append(0)
            else:
                secret_commun.append(1)

        return secret_commun

# Exemple d'utilisation
generateur_secret = GenerationSecretThroughAnonymousCanal()

nom_a = "Alice"
nom_b = "Bob"
duree_protocole = 10  # en secondes

transcript_a = generateur_secret.genererSecret(nom_a, nom_b, duree_protocole)
transcript_b = generateur_secret.genererSecret(nom_a, nom_b, duree_protocole)

secret_extrait = generateur_secret.extraireSecret(transcript_a, transcript_b)
print(f"Secret extrait : {secret_extrait}")
