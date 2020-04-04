import boto3

class Ecr():

    def __init__(self, profile, region):
        self.profile = profile
        self.region = region

    def create_repo(self, reponame):
        conn = boto3.session.Session(profile_name=self.profile, region_name=self.region)
        client = conn.client('ecr')

        response = client.create_repository(
            repositoryName=reponame
        )

        return response