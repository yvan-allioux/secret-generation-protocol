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

def genererSecret(nom_a, nom_b, duree):
    canal = CanalDiffusionAnonyme()
    secret = []

    debut_periode = time.time()
    fin_periode = debut_periode + duree

    while time.time() < fin_periode:
        b = random.randint(0, 1)
        message = f"{nom_a}" if b == 0 else f"{nom_b}"
        canal.posterMessageAnonyme(message)
        time.sleep(random.uniform(0.001, 0.01))  # Attente aléatoire entre 1 et 10 ms

    messages_recuperes = canal.recupererMessagesAnonymes(debut_periode, fin_periode)

    for message in messages_recuperes:
        if nom_a in message:
            secret.append(0)
        else:
            secret.append(1)

    return secret

# Extraction du secret
def extraireSecret(transcript_a, transcript_b):
    secret_commun = []

    for message_a, message_b in zip(transcript_a, transcript_b):
        if message_a == message_b:
            secret_commun.append(0)
        else:
            secret_commun.append(1)

    return secret_commun


# Exemple d'utilisation
nom_a = "Alice"
nom_b = "Bob"
duree_protocole = 60  # en secondes

transcript_a = genererSecret(nom_a, nom_b, duree_protocole)
print(f"Secret généré : {transcript_a}")
transcript_b = genererSecret(nom_a, nom_b, duree_protocole)
print(f"Secret généré : {transcript_b}")

secret_extrait = extraireSecret(transcript_a, transcript_b)
print(f"Secret extrait : {secret_extrait}")
