import boto3
from pathlib import Path


class Ec2():

    def __init__(self, profile, region):
        self.profile = profile
        self.region = region

    def create_key(self):
        conn = boto3.session.Session(profile_name=self.profile, region_name=self.region)
        client = conn.client('ec2')

        response = client.create_key_pair(
            KeyName='hive_key',
        )
        try:
            home = str(Path.home())
            with open(home+'/.ssh/hive_key', 'w+') as key:
                key.write(response['KeyMaterial'])
        except Exception as e:
            print(e)

        return response

    def create_instance(self, ami, type, key):
        conn = boto3.session.Session(profile_name=self.profile, region_name=self.region)
        ec2 = conn.resource('ec2')
        tag_name = {"Key": "Name", "Value": "Hive_App"}

        instance = ec2.create_instances(
            ImageId=ami,
            MinCount=1,
            MaxCount=1,
            InstanceType=type,
            KeyName=key,
            TagSpecifications=[{'ResourceType': 'instance',
                            'Tags': [tag_name]}])

        return instance
