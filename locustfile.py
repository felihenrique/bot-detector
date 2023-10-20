from locust import HttpUser, task, between
import json
import random
from datetime import datetime, timedelta

class MyUser(HttpUser):
    wait_time = between(1, 5)  # Define o intervalo de espera entre as solicitações

    @task
    def post_request(self):
        start_date = datetime(2023, 1, 1)
        end_date = datetime(2023, 10, 1)

        # Generate a random date within the range
        random_date = start_date + timedelta(days=random.randint(0, (end_date - start_date).days))

        # Format the date in the desired format
        formatted_date = random_date.strftime('%Y-%m-%dT00:00:00Z')

        # Gerar um IP aleatório (substitua isso pela lógica de geração real)
        random_ip = ".".join(map(str, (random.randint(0, 255) for _ in range(4))))

        # Criar o payload do corpo da solicitação
        payload = {
            "player_id": 111,
            "ip": random_ip,
            "user_agent": "teste",
            "created_at": formatted_date
        }

        # Define os cabeçalhos da solicitação
        headers = {
            "Content-Type": "application/json"
        }

        # Envia a solicitação POST para a rota desejada
        self.client.post("/logs", data=json.dumps(payload), headers=headers)
