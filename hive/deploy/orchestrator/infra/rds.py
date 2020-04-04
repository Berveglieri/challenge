import boto3
import string
import random

class Rds():

        def __init__(self, profile, region):
            self.profile = profile
            self.region = region

        def create_rds(self):
            conn = boto3.session.Session(profile_name=self.profile, region_name=self.region)
            client = conn.client('rds')

            response = client.create_db_instance(
                DBName='hive',
                DBInstanceIdentifier='app-db',
                AllocatedStorage=20,
                DBInstanceClass='db.t2.micro',
                Engine='postgres',
                MasterUsername='master',
                MasterUserPassword=self.random_password(),
            )

            return response

        def random_password(self):
            password = string.ascii_letters

            master_pass = ''.join(random.choice(password) for i in range(34))
            print("The master password is "+master_pass)
            return master_pass
