from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        # self.client.get("/read/1906398420/100")
        self.client.get("/read/1906398420/")