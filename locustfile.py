from locust import HttpUser, task, between
import json
import random

class MyUser(HttpUser):
    wait_time = between(1, 5)  # Define o intervalo de espera entre as solicitações

    @task
    def post_request(self):
        # Gerar um IP aleatório (substitua isso pela lógica de geração real)
        random_ip = ".".join(map(str, (random.randint(0, 255) for _ in range(4))))

        # Criar o payload do corpo da solicitação
        payload = {
            "player_id": 123,
            "ip": random_ip,
            "user_agent": "teste"
        }

        # Define os cabeçalhos da solicitação
        headers = {
            "Content-Type": "application/json"
        }

        # Envia a solicitação POST para a rota desejada
        self.client.post("/requests", data=json.dumps(payload), headers=headers)
