from infra import Ec2 as ec2
from infra import Ecr as ecr
from infra import Rds as rds
from infra import Cdn as cdn
from builder import Builder as bd
from pusher import Push as ps



if __name__ == '__main__':

    #infra

    # #create ssh key pair and save it in your home
    # ec2("hive", "eu-central-1").create_key()
    # #create instance using the new key pair
    # ec2("hive", "eu-central-1").create_instance("ami-0ec1ba09723e5bfac", "t2.micro", "hive_key")
    # #create ecr repo to store the application
    # print(ecr("hive", "eu-central-1").create_repo("web_app"))
    #create RDS instance
    #rds("hive", "eu-central-1").create_rds()




    #build

    # #build the application
    bd().docker_build("v1.0")
    #
    # #push
    # ps("hive", "eu-central-1").ecr_push()



