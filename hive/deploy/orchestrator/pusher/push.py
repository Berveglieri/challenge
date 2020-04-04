import boto3
import docker
import base64

class Push():

    def __init__(self, profile, region):
        self.profile = profile
        self.region = region

    def ecr_auth(self):
        conn = boto3.session.Session(profile_name=self.profile, region_name=self.region)
        auth = conn.client('ecr').get_authorization_token()
        encoded_token = auth['authorizationData'][0]['authorizationToken']
        token = base64.b64decode(encoded_token).decode()
        username, password = token.split(':')
        auth_config = {'username': username, 'password': password}

        return auth_config

    def ecr_push(self):
        client = docker.from_env()
        print(client.images.push(repository="017136389210.dkr.ecr.eu-central-1.amazonaws.com/web_app", auth_config=self.ecr_auth(), stream=False))




