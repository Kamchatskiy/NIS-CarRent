import requests
from faker import Faker  # type: ignore
from random import randint

url = "http://app:8080"
admin_key_file = open("/run/secrets/admin-key", "r")
admin_key = admin_key_file.readline()
fake = Faker()


def generate_cars(size):
    successful_counts = 0
    for i in range(size):
        resp = requests.post(
            url=f"{url}/admin/cars/new",
            headers={"admin-key": admin_key},
            json={
                "brand": fake.company(),
                "model": fake.word(),
                "year": randint(1886, 2050),
                "daily_price": randint(100, 1000000),
                "insurance_price": randint(100, 500000),
            },
        )
        if resp.status_code in (200, 201):
            successful_counts += 1
    print(f"successfully added {successful_counts}/{size} cars")


def generate_clients(size):
    successful_counts = 0
    for i in range(size):
        resp = requests.post(
            url=f"{url}/admin/clients/new",
            headers={"admin-key": admin_key},
            json={
                "name": fake.first_name(),
                "surname": fake.last_name(),
                "email": fake.email(),
                "phone_number": fake.phone_number(),
            },
        )
        if resp.status_code in (200, 201):
            counts += 1
    print(f"successfully added {successful_counts}/{size} clients")


def generate_rents(rents_size, cars_size, clients_size):
    successful_counts = 0
    for i in range(rents_size):
        resp = requests.post(
            url=f"{url}/admin/rents/new",
            headers={"admin-key": admin_key},
            json={
                "start_date": fake.date_time_this_year().strftime("%Y-%m-%dT%H:%M:%SZ"),
                "end_date": fake.date_time_this_year().strftime("%Y-%m-%dT%H:%M:%SZ"),
                "car_id": randint(1, cars_size),
                "client_id": randint(1, clients_size),
            },
        )
        if resp.status_code in (200, 201):
            successful_counts += 1
    print(f"successfully added {successful_counts}/{rents_size} rents")


def main():
    cars_size = 10
    generate_cars(cars_size)
    clients_size = 25
    generate_clients(clients_size)
    rents_size = 5
    generate_rents(rents_size, cars_size, clients_size)


if __name__ == "__main__":
    main()
