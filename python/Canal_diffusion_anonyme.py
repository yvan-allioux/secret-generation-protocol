import time

class CanalDiffusionAnonyme:
    def __init__(self):
        self.messages = []

    def posterMessageAnonyme(self, message):
        timestamp = time.strftime("%Hh%Mm%Ss", time.localtime())
        nouveau_message = {"message": message, "timestamp": timestamp}
        self.messages.append(nouveau_message)
        self.messages = sorted(self.messages, key=lambda x: x["timestamp"])

    def recupererMessagesAnonymes(self, debut, fin):
        messages_anonymes = []
        for message in self.messages:
            if debut <= message["timestamp"]  and  message["timestamp"] <= fin:
                messages_anonymes.append(message["message"])
        return messages_anonymes

    def afficherMessages(self):
        for message in self.messages:
            print(f"{message['timestamp']} : {message['message']}")

# Exemple d'utilisation
canal = CanalDiffusionAnonyme()
canal.posterMessageAnonyme("l'enssat est la plus belle ecole au Monde")
canal.afficherMessages()

# Récupérer les messages postés entre deux moments
debut_periode = "14h00m00s"
fin_periode = "16h30m00s"
messages_recuperes = canal.recupererMessagesAnonymes(debut_periode, fin_periode)

print("\nMessages postés entre", debut_periode, "et", fin_periode, ":\n", messages_recuperes)
