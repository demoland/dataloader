import json
import random
from faker import Faker

fake = Faker()

job_types = ['manager', 'product-manager' , 'executive', 'devops-engineer', 'devsecops-engineer', 'developer', "customer-success"]


users = []

for _ in range(50):
    user = {
        "name": fake.name(),
        "userid": fake.uuid4(),
        "address": fake.address().replace("\n", ", "),
        "phone": fake.phone_number(),
        "useragent": fake.user_agent(),
        "company": "OpenGovCo",
        "email": fake.email(),
        "team": random.choice(job_types),
        "location": fake.city(),
        "creditcard": fake.credit_card_number(card_type='mastercard'),
        "socialsecurity": fake.ssn(),
    }
    users.append(user)

with open("users.json", "w") as f:
    json.dump(users, f, indent=2)

