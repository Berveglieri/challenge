import os
import subprocess

class Builder():

    def docker_build(self, tag):
        try:
            root = os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))
            df_path = root+os.path.join('/hive')
            dockerfile = os.path.isfile(os.path.join(df_path,'Dockerfile'))
            if dockerfile:
                os.chdir(df_path)
                call = subprocess.Popen(['docker', 'build', '-t', '017136389210.dkr.ecr.eu-central-1.amazonaws.com/web_app:'+tag, '.'],
                                        stdout=subprocess.PIPE,
                                        stderr=subprocess.PIPE)
                stdout, stderr = call.communicate()
                print(stdout.decode('UTF-8'), stderr.decode('UTF-8'))
        except FileNotFoundError as e:
            print(e)
